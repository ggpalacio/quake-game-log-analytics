package game

import "fmt"

type MatchReport struct {
	Players      []string           `json:"players"`
	TotalKills   int                `json:"total_kills"`
	Kills        map[string]int     `json:"kills"`
	KillsByMeans map[DeathCause]int `json:"kills_by_means"`
}

type PlayerRankingReport struct {
	Position int    `json:"position"`
	Player   string `json:"player"`
	Score    int    `json:"score"`
}

func NewReport(logFile *LogFile) map[string]any {
	report := make(map[string]any)
	matches := process(logFile)
	for _, match := range matches {
		report[match.ID] = MatchReport{
			TotalKills:   len(match.Kills),
			Players:      getPlayerNames(match),
			Kills:        getKillScoreByPlayer(match),
			KillsByMeans: countKillsByDeathCause(match),
		}
	}
	return report
}

func process(logFile *LogFile) []*Match {
	var match *Match
	var matchIndex int
	var matches []*Match
	for _, log := range logFile.Logs {
		if log.IsInitGame() {
			matchIndex++
			match = NewMatch(fmt.Sprintf("game_%03d", matchIndex))
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
			match.RegisterKill(killerName, killedName, DeathCause(deathCause))
			continue
		}
	}
	return matches
}

func getPlayerNames(match *Match) []string {
	playerNames := make([]string, len(match.Players))
	index := 0
	for name, _ := range match.Players {
		playerNames[index] = name
		index++
	}
	return playerNames
}

func getKillScoreByPlayer(match *Match) map[string]int {
	killScoreByPlayer := make(map[string]int)
	for _, player := range match.Players {
		playerKillScore, _ := match.GetKillScore(player.Name)
		killScoreByPlayer[player.Name] = playerKillScore
	}
	return killScoreByPlayer
}

func countKillsByDeathCause(match *Match) map[DeathCause]int {
	killsByDeathScore := make(map[DeathCause]int)
	for _, kill := range match.Kills {
		killsByDeathScore[kill.Cause]++
	}
	return killsByDeathScore
}
