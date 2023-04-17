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

func TestNewMatchReport(t *testing.T) {
	world := &game.Player{Name: game.WorldName}
	foo := &game.Player{Name: "foo"}
	bar := &game.Player{Name: "bar"}
	foo.Kills = []*game.Kill{
		{Killer: foo, Killed: bar, Cause: game.DeathCauseMachineGun},
		{Killer: foo, Killed: bar, Cause: game.DeathCauseMachineGun},
		{Killer: foo, Killed: bar, Cause: game.DeathCauseMachineGun},
	}
	bar.Kills = []*game.Kill{
		{Killer: bar, Killed: foo, Cause: game.DeathCauseShotGun},
		{Killer: bar, Killed: foo, Cause: game.DeathCauseShotGun},
	}
	match := &game.Match{
		ID: "game-1",
		Players: map[string]*game.Player{
			foo.Name: foo,
			bar.Name: bar,
		},
		World: world,
	}
	match.Kills = append(match.Kills, foo.Kills...)
	match.Kills = append(match.Kills, bar.Kills...)

	matchReport := report.NewMatchReport(match)
	assert.Equal(t, "game-1", matchReport.MatchID)
	assert.Equal(t, []string{"foo", "bar"}, matchReport.Players)
	assert.Equal(t, 3, matchReport.Kills["foo"])
	assert.Equal(t, 2, matchReport.Kills["bar"])
	assert.Equal(t, 3, matchReport.KillsByMeans[game.DeathCauseMachineGun])
	assert.Equal(t, 2, matchReport.KillsByMeans[game.DeathCauseShotGun])
}

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
