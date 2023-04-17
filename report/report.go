package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
)

type Report struct {
	Matches MatchesReport
	Ranking RankingReport
}

func NewReport(logFile *game.LogFile) Report {
	report := Report{}

	matches := process(logFile)
	for _, match := range matches {
		matchReport := NewMatchReport(match)
		report.Matches = append(report.Matches, matchReport)
		report.Ranking.AddPlayersScore(matchReport.Kills)
	}
	return report
}

func process(logFile *game.LogFile) []*game.Match {
	var match *game.Match
	var matchIndex int
	var matches []*game.Match
	for _, log := range logFile.Logs {
		if log.IsInitGame() {
			matchIndex++
			match = game.NewMatch(fmt.Sprintf("game-%d", matchIndex))
			matches = append(matches, match)
		}
		if match == nil {
			continue
		}
		if playerID, playerName := log.ClientUserinfoChanged(); playerID != 0 {
			match.AddPlayer(playerName)
			continue
		}
		if killerName, killedName, deathCause := log.Kill(); killerName != "" {
			match.RegisterKill(killerName, killedName, game.DeathCause(deathCause))
			continue
		}
	}
	return matches
}

func (ref Report) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	if len(ref.Matches) > 0 {
		buf.Write(ref.marshalMatches())
	}
	if len(ref.Ranking.GePlayerRanking()) > 0 {
		if len(ref.Matches) > 0 {
			buf.WriteString(",")
		}
		buf.Write(ref.marshalRanking())
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (ref Report) marshalMatches() []byte {
	var buf bytes.Buffer
	var index int
	for _, match := range ref.Matches {
		if index > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"%s":`, match.MatchID))
		value, _ := json.Marshal(match)
		buf.Write(value)
		index++
	}
	return buf.Bytes()
}

func (ref Report) marshalRanking() []byte {
	var buf bytes.Buffer
	buf.WriteString(`"ranking":{`)
	for index, playerRanking := range ref.Ranking.GePlayerRanking() {
		if index > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"%s":%d`, playerRanking.Player, playerRanking.Score))
	}
	buf.WriteString("}")
	return buf.Bytes()
}
