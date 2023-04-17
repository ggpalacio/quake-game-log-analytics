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
