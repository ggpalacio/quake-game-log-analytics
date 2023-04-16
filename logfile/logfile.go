package logfile

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const logLineBeginRegex = `(\d*:\d{2})\s`

var (
	logLineRegex               = regexp.MustCompile(logLineBeginRegex + `(.*)`)
	ignoreLogLineRegex         = regexp.MustCompile(logLineBeginRegex + `-{60}$`)
	clientUserinfoChangedRegex = regexp.MustCompile(`ClientUserinfoChanged: (\d*) n\\(.*?)\\t`)
	killRegex                  = regexp.MustCompile(`Kill: \d* \d* \d*: (.*) killed (.*) by (.*)`)
)

type LogFile struct {
	Lines []LogLine
}

type LogLine struct {
	Time    string
	Message string
}

func (ref LogLine) IsInitGame() bool {
	return strings.HasPrefix(ref.Message, "InitGame:")
}

func (ref LogLine) ClientUserinfoChanged() (clientID int, clientName string) {
	subMatches := clientUserinfoChangedRegex.FindStringSubmatch(ref.Message)
	if subMatches == nil {
		return
	}
	clientID, _ = strconv.Atoi(subMatches[1])
	clientName = subMatches[2]
	return
}

func (ref LogLine) Kill() (killer, killed, deathCause string) {
	subMatches := killRegex.FindStringSubmatch(ref.Message)
	if subMatches == nil {
		return
	}
	killer = subMatches[1]
	killed = subMatches[2]
	deathCause = subMatches[3]
	return
}

func New(fileName string) (*LogFile, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		return nil, err
	}

	return &LogFile{
		Lines: lines,
	}, nil
}

func readLines(file *os.File) ([]LogLine, error) {
	var lines []LogLine
	scanner := bufio.NewScanner(file)
	for index := 1; scanner.Scan(); index++ {
		line, err := readLine(index, scanner.Text())
		if err != nil {
			return nil, err
		}
		if line == nil {
			continue
		}
		lines = append(lines, *line)
	}
	return lines, nil
}

func readLine(index int, line string) (*LogLine, error) {
	if ignoreLogLineRegex.MatchString(line) {
		return nil, nil
	}

	subMatches := logLineRegex.FindStringSubmatch(line)
	if subMatches == nil {
		return nil, errors.New(fmt.Sprintf("invalid log line format at line %d:%s", index, line))
	}
	return &LogLine{
		Time:    subMatches[1],
		Message: subMatches[2],
	}, nil
}
