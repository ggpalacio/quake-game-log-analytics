package report

import (
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
)

type Report struct {
	Matches MatchesReport `json:"matches"`
	Ranking RankingReport `json:"ranking,omitempty"`
}

func NewReport(logFile *game.LogFile) Report {
	var report Report

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
