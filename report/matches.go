package report

import (
	"github.com/ggpalacio/quake-game-log-analytics/game"
	"sort"
)

type MatchesReport map[string]MatchReport

type PlayerNames []string

type MatchReport struct {
	Players      PlayerNames             `json:"players"`
	TotalKills   int                     `json:"total_kills"`
	Kills        map[string]int          `json:"kills"`
	KillsByMeans map[game.DeathCause]int `json:"kills_by_means"`
}

func NewMatchReport(match *game.Match) MatchReport {
	return MatchReport{
		TotalKills:   len(match.Kills),
		Players:      getPlayerNames(match),
		Kills:        getKillScoreByPlayer(match),
		KillsByMeans: countKillsByDeathCause(match),
	}
}

func getPlayerNames(match *game.Match) PlayerNames {
	playerNames := make(PlayerNames, len(match.Players))
	index := 0
	for name, _ := range match.Players {
		playerNames[index] = name
		index++
	}
	sort.Sort(playerNames)
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

func (ref PlayerNames) Len() int {
	return len(ref)
}

func (ref PlayerNames) Less(i, j int) bool {
	return ref[i] < ref[j]
}

func (ref PlayerNames) Swap(i, j int) {
	tmp := ref[i]
	ref[i] = ref[j]
	ref[j] = tmp
}
