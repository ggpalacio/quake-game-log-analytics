package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
)

type MatchesReport []MatchReport

type MatchReport struct {
	MatchID      string                  `json:"-"`
	Players      []string                `json:"players"`
	TotalKills   int                     `json:"total_kills"`
	Kills        map[string]int          `json:"kills"`
	KillsByMeans map[game.DeathCause]int `json:"kills_by_means"`
}

func NewMatchReport(match *game.Match) MatchReport {
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

func (ref MatchesReport) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	for index, report := range ref {
		if index > 0 {
			buf.WriteString(",")
		}

		buf.WriteString(fmt.Sprintf(`"%s":`, report.MatchID))

		val, _ := json.Marshal(report)
		buf.Write(val)
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}
