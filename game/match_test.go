package game_test

import (
	"github.com/ggpalacio/quake-game-log-analytics/game"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatch_AddPlayer(t *testing.T) {
	match := game.NewMatch("test")
	match.AddPlayer("foo")
	match.AddPlayer("bar")
	assert.Equal(t, "test", match.ID)
	assert.NotNil(t, match.World)
	assert.Equal(t, game.WorldName, match.World.Name)
	assert.Len(t, match.Players, 2)
	assert.NotNil(t, match.Players["foo"])
	assert.NotNil(t, match.Players["bar"])
	assert.Nil(t, match.Players["baz"])
}

func TestMatch_RegisterKill(t *testing.T) {
	match := game.NewMatch("test")
	match.AddPlayer("foo")
	match.AddPlayer("bar")
	match.AddPlayer("baz")

	assert.NoError(t, match.RegisterKill("foo", "bar", game.DeathCauseShotGun))
	assert.NoError(t, match.RegisterKill("baz", "foo", game.DeathCauseRocket))
	assert.NoError(t, match.RegisterKill("baz", "bar", game.DeathCauseRocket))
	assert.NoError(t, match.RegisterKill("bar", "baz", game.DeathCauseGrenade))
	assert.NoError(t, match.RegisterKill("bar", "bar", game.DeathCauseGrenade))
	assert.NoError(t, match.RegisterKill("foo", "bar", game.DeathCauseMachineGun))
	assert.NoError(t, match.RegisterKill("foo", "baz", game.DeathCauseMachineGun))
	assert.NoError(t, match.RegisterKill("baz", "foo", game.DeathCauseRailGun))
	assert.NoError(t, match.RegisterKill("baz", "bar", game.DeathCauseRailGun))
	assert.NoError(t, match.RegisterKill(game.WorldName, "bar", game.DeathCauseFalling))
	assert.NoError(t, match.RegisterKill("baz", "foo", game.DeathCauseRailGun))
	assert.NoError(t, match.RegisterKill(game.WorldName, "baz", game.DeathCauseWater))
	assert.NoError(t, match.RegisterKill("foo", "bar", game.DeathCauseProximityMine))
	assert.NoError(t, match.RegisterKill(game.WorldName, "foo", game.DeathCauseSlime))

	killScore, err := match.GetKillScore("foo")
	assert.Equal(t, 3, killScore)
	assert.NoError(t, err)

	killScore, err = match.GetKillScore("bar")
	assert.Zero(t, killScore)
	assert.NoError(t, err)

	killScore, err = match.GetKillScore("baz")
	assert.Equal(t, 4, killScore)
	assert.NoError(t, err)

	assert.Equal(t, 14, len(match.Kills))
}

func TestMatch_AddPlayer_Error(t *testing.T) {
	match := game.NewMatch("test")
	match.AddPlayer(game.WorldName)
	err := match.AddPlayer(game.WorldName)
	assert.EqualError(t, err, "cannot add a player '<world>'")
}

func TestMatch_RegisterKill_Error(t *testing.T) {
	match := game.NewMatch("test")
	err := match.RegisterKill("foo", "bar", game.DeathCauseUnknown)
	assert.EqualError(t, err, "player 'foo' not found in match")

	match.AddPlayer("foo")
	err = match.RegisterKill("foo", "bar", game.DeathCauseUnknown)
	assert.EqualError(t, err, "player 'bar' not found in match")

	match.AddPlayer("foo")
	err = match.RegisterKill("foo", "foo", "MOD_PUNCHING")
	assert.EqualError(t, err, "death cause 'MOD_PUNCHING' not exists")
}

func TestMatch_GetKillScore_Error(t *testing.T) {
	match := game.NewMatch("test")
	killScore, err := match.GetKillScore("foo")
	assert.Zero(t, killScore)
	assert.EqualError(t, err, "player 'foo' not found in match")
}
