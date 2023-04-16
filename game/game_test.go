package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_AddPlayer(t *testing.T) {
	game := new(Game)
	game.AddPlayer("foo")
	game.AddPlayer("bar")
	assert.Equal(t, []string{"foo", "bar"}, game.Players)
}

func TestGame_RegisterKill(t *testing.T) {
	game := &Game{
		Players: []string{"foo", "bar", "baz"},
	}
	game.RegisterKill("foo", "bar", "MOD_SHOTGUN")
	game.RegisterKill("baz", "foo", "MOD_ROCKET")
	game.RegisterKill("baz", "bar", "MOD_ROCKET")
	game.RegisterKill("bar", "baz", "MOD_GRENADE")
	game.RegisterKill("bar", "bar", "MOD_GRENADE_SPLASH")
	game.RegisterKill("foo", "bar", "MOD_MACHINEGUN")
	game.RegisterKill("foo", "baz", "MOD_MACHINEGUN")
	game.RegisterKill("baz", "foo", "MOD_RAILGUN")
	game.RegisterKill("baz", "bar", "MOD_RAILGUN")
	game.RegisterKill(world, "bar", "MOD_FALLING")
	game.RegisterKill("baz", "foo", "MOD_RAILGUN")
	game.RegisterKill(world, "baz", "MOD_WATER")
	game.RegisterKill("foo", "bar", "MOD_PROXIMITY_MINE")
	game.RegisterKill(world, "foo", "MOD_SLIME")

	assert.Equal(t, 3, game.Kills["foo"])
	assert.Zero(t, game.Kills["bar"])
	assert.Equal(t, 4, game.Kills["baz"])
	assert.Equal(t, 14, game.TotalKills)
	assert.Equal(t, 1, game.KillsByMeans["MOD_SHOTGUN"])
	assert.Equal(t, 2, game.KillsByMeans["MOD_ROCKET"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_GRENADE"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_GRENADE_SPLASH"])
	assert.Equal(t, 2, game.KillsByMeans["MOD_MACHINEGUN"])
	assert.Equal(t, 3, game.KillsByMeans["MOD_RAILGUN"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_PROXIMITY_MINE"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_FALLING"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_WATER"])
	assert.Equal(t, 1, game.KillsByMeans["MOD_SLIME"])
}
