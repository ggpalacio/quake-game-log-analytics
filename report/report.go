package report

import (
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
)

type Report struct {
	Matches MatchesReport  `json:"matches"`
	Ranking *RankingReport `json:"ranking,omitempty"`
}

func NewReport(logFile *game.LogFile) Report {
	report := Report{
		Ranking: new(RankingReport),
	}

	matches := process(logFile)
	for _, match := range matches {
		matchReport := createMatchReport(match)
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

func createMatchReport(match *game.Match) MatchReport {
	return MatchReport{
		MatchID:      match.ID,
		TotalKills:   len(match.Kills),
		Players:      getPlayerNames(match),
		Kills:        getKillScoreByPlayer(match),
		KillsByMeans: countKillsByDeathCause(match),
	}
}

func getPlayerNames(match *game.Match) []string {
	playerNames := make([]string, len(match.Players))
	index := 0
	for name, _ := range match.Players {
		playerNames[index] = name
		index++
	}
	return playerNames
}

func getKillScoreByPlayer(match *game.Match) map[string]int {
	killScoreByPlayer := make(map[string]int)
	for _, player := range match.Players {
		playerKillScore, _ := match.GetKillScore(player.Name)
		killScoreByPlayer[player.Name] = playerKillScore
	}
	return killScoreByPlayer
}

func countKillsByDeathCause(match *game.Match) map[game.DeathCause]int {
	killsByDeathScore := make(map[game.DeathCause]int)
	for _, kill := range match.Kills {
		killsByDeathScore[kill.Cause]++
	}
	return killsByDeathScore
}
