package game

import (
	"fmt"
)

type Report map[string]*Match

func NewReport(logFile *LogFile) Report {
	var match *Match
	var matchIndex int
	report := make(Report)

	for _, logLine := range logFile.Logs {
		if logLine.IsInitGame() {
			match = new(Match)
			matchIndex++

			matchKey := fmt.Sprintf("game_%s", fmt.Sprintf("%03d", matchIndex))
			report[matchKey] = match
		}
		if match == nil {
			continue
		}
		if playerID, playerName := logLine.ClientUserinfoChanged(); playerID != 0 {
			match.AddPlayer(playerName)
			continue
		}
		if killer, killed, deathCause := logLine.Kill(); killer != "" {
			match.RegisterKill(killer, killed, deathCause)
			continue
		}
	}
	return report
}
