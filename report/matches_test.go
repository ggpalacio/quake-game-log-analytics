package report_test

import (
	_ "embed"
	"encoding/json"
	"github.com/ggpalacio/quake-game-log-analytics/game"
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed matches_report.json
var matchesReportFile []byte

func TestMatchesReport_MarshalJSON(t *testing.T) {
	matchesReport := report.MatchesReport{
		{
			MatchID:    "game-1",
			Players:    []string{"foo", "bar", "baz"},
			TotalKills: 10,
			Kills: map[string]int{
				"foo": 5,
				"bar": 3,
				"baz": 2,
			},
			KillsByMeans: map[game.DeathCause]int{
				game.DeathCauseMachineGun: 3,
				game.DeathCauseShotGun:    7,
			},
		},
		{
			MatchID:    "game-2",
			Players:    []string{"foo", "bar"},
			TotalKills: 12,
			Kills: map[string]int{
				"foo": 5,
				"bar": 7,
			},
			KillsByMeans: map[game.DeathCause]int{
				game.DeathCauseRailGun:  8,
				game.DeathCauseChainGun: 4,
			},
		},
	}

	matchesReportJSON, err := json.MarshalIndent(matchesReport, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, string(matchesReportFile), string(matchesReportJSON))
}
