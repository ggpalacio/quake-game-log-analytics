package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch_AddPlayer(t *testing.T) {
	match := new(Match)
	match.AddPlayer("foo")
	match.AddPlayer("bar")
	assert.Equal(t, []string{"foo", "bar"}, match.Players)
}

func TestMatch_RegisterKill(t *testing.T) {
	match := &Match{
		Players: []string{"foo", "bar", "baz"},
	}
	match.RegisterKill("foo", "bar", "MOD_SHOTGUN")
	match.RegisterKill("baz", "foo", "MOD_ROCKET")
	match.RegisterKill("baz", "bar", "MOD_ROCKET")
	match.RegisterKill("bar", "baz", "MOD_GRENADE")
	match.RegisterKill("bar", "bar", "MOD_GRENADE_SPLASH")
	match.RegisterKill("foo", "bar", "MOD_MACHINEGUN")
	match.RegisterKill("foo", "baz", "MOD_MACHINEGUN")
	match.RegisterKill("baz", "foo", "MOD_RAILGUN")
	match.RegisterKill("baz", "bar", "MOD_RAILGUN")
	match.RegisterKill(world, "bar", "MOD_FALLING")
	match.RegisterKill("baz", "foo", "MOD_RAILGUN")
	match.RegisterKill(world, "baz", "MOD_WATER")
	match.RegisterKill("foo", "bar", "MOD_PROXIMITY_MINE")
	match.RegisterKill(world, "foo", "MOD_SLIME")

	assert.Equal(t, 3, match.Kills["foo"])
	assert.Zero(t, match.Kills["bar"])
	assert.Equal(t, 4, match.Kills["baz"])
	assert.Equal(t, 14, match.TotalKills)
	assert.Equal(t, 1, match.KillsByMeans["MOD_SHOTGUN"])
	assert.Equal(t, 2, match.KillsByMeans["MOD_ROCKET"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_GRENADE"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_GRENADE_SPLASH"])
	assert.Equal(t, 2, match.KillsByMeans["MOD_MACHINEGUN"])
	assert.Equal(t, 3, match.KillsByMeans["MOD_RAILGUN"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_PROXIMITY_MINE"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_FALLING"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_WATER"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_SLIME"])
}
