package game

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
	logRegex                      = regexp.MustCompile(logLineBeginRegex + `(.*)`)
	ignoreLogRegex                = regexp.MustCompile(logLineBeginRegex + `-{60}$`)
	clientUserinfoChangedLogRegex = regexp.MustCompile(`ClientUserinfoChanged: (\d*) n\\(.*?)\\t`)
	killLogRegex                  = regexp.MustCompile(`Kill: \d* \d* \d*: (.*) killed (.*) by (.*)`)
)

type LogFile struct {
	Logs []Log
}

type Log struct {
	Time    string
	Message string
}

func (ref Log) IsInitGame() bool {
	return strings.HasPrefix(ref.Message, "InitGame:")
}

func (ref Log) ClientUserinfoChanged() (clientID int, clientName string) {
	subMatches := clientUserinfoChangedLogRegex.FindStringSubmatch(ref.Message)
	if subMatches == nil {
		return
	}
	clientID, _ = strconv.Atoi(subMatches[1])
	clientName = subMatches[2]
	return
}

func (ref Log) Kill() (killerName, killedName, deathCause string) {
	subMatches := killLogRegex.FindStringSubmatch(ref.Message)
	if subMatches == nil {
		return
	}
	killerName = subMatches[1]
	killedName = subMatches[2]
	deathCause = subMatches[3]
	return
}

func NewLogFile(fileName string) (*LogFile, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logs, err := readLines(file)
	if err != nil {
		return nil, err
	}

	return &LogFile{
		Logs: logs,
	}, nil
}

func readLines(file *os.File) ([]Log, error) {
	var logs []Log
	scanner := bufio.NewScanner(file)
	for index := 1; scanner.Scan(); index++ {
		log, err := readLine(index, scanner.Text())
		if err != nil {
			return nil, err
		}
		if log == nil {
			continue
		}
		logs = append(logs, *log)
	}
	return logs, nil
}

func readLine(index int, line string) (*Log, error) {
	if ignoreLogRegex.MatchString(line) {
		return nil, nil
	}

	subMatches := logRegex.FindStringSubmatch(line)
	if subMatches == nil {
		return nil, errors.New(fmt.Sprintf("invalid log format at line %d:%s", index, line))
	}
	return &Log{
		Time:    subMatches[1],
		Message: subMatches[2],
	}, nil
}
