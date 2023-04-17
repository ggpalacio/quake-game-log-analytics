package report_test

import (
	"github.com/ggpalacio/quake-game-log-analytics/game"
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.Equal(t, report.PlayerNames{"bar", "foo"}, matchReport.Players)
	assert.Equal(t, 3, matchReport.Kills["foo"])
	assert.Equal(t, 2, matchReport.Kills["bar"])
	assert.Equal(t, 3, matchReport.KillsByMeans[game.DeathCauseMachineGun])
	assert.Equal(t, 2, matchReport.KillsByMeans[game.DeathCauseShotGun])
}
