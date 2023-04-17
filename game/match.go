package game

import (
	"errors"
	"fmt"
)

const (
	WorldName = "<world>"

	DeathCauseUnknown       DeathCause = "MOD_UNKNOWN"
	DeathCauseShotGun       DeathCause = "MOD_SHOTGUN"
	DeathCauseGauntlet      DeathCause = "MOD_GAUNTLET"
	DeathCauseMachineGun    DeathCause = "MOD_MACHINEGUN"
	DeathCauseGrenade       DeathCause = "MOD_GRENADE"
	DeathCauseGrenadeSplash DeathCause = "MOD_GRENADE_SPLASH"
	DeathCauseRocket        DeathCause = "MOD_ROCKET"
	DeathCauseRocketSplash  DeathCause = "MOD_ROCKET_SPLASH"
	DeathCausePlasma        DeathCause = "MOD_PLASMA"
	DeathCausePlasmaSplash  DeathCause = "MOD_PLASMA_SPLASH"
	DeathCauseRailGun       DeathCause = "MOD_RAILGUN"
	DeathCauseLighting      DeathCause = "MOD_LIGHTNING"
	DeathCauseBFG           DeathCause = "MOD_BFG"
	DeathCauseBFGSplash     DeathCause = "MOD_BFG_SPLASH"
	DeathCauseWater         DeathCause = "MOD_WATER"
	DeathCauseSlime         DeathCause = "MOD_SLIME"
	DeathCauseLava          DeathCause = "MOD_LAVA"
	DeathCauseCrush         DeathCause = "MOD_CRUSH"
	DeathCauseTelefrag      DeathCause = "MOD_TELEFRAG"
	DeathCauseFalling       DeathCause = "MOD_FALLING"
	DeathCauseSuicide       DeathCause = "MOD_SUICIDE"
	DeathCauseTargetLaser   DeathCause = "MOD_TARGET_LASER"
	DeathCauseTriggerHurt   DeathCause = "MOD_TRIGGER_HURT"
	DeathCauseNail          DeathCause = "MOD_NAIL"
	DeathCauseChainGun      DeathCause = "MOD_CHAINGUN"
	DeathCauseProximityMine DeathCause = "MOD_PROXIMITY_MINE"
	DeathCauseKamikaze      DeathCause = "MOD_KAMIKAZE"
	DeathCauseJuiced        DeathCause = "MOD_JUICE"
	DeathCauseGrapple       DeathCause = "MOD_GRAPLE"
)

var deathCauses = map[string]DeathCause{
	"MOD_UNKNOWN":        DeathCauseUnknown,
	"MOD_SHOTGUN":        DeathCauseShotGun,
	"MOD_GAUNTLET":       DeathCauseGauntlet,
	"MOD_MACHINEGUN":     DeathCauseMachineGun,
	"MOD_GRENADE":        DeathCauseGrenade,
	"MOD_GRENADE_SPLASH": DeathCauseGrenadeSplash,
	"MOD_ROCKET":         DeathCauseRocket,
	"MOD_ROCKET_SPLASH":  DeathCauseRocketSplash,
	"MOD_PLASMA":         DeathCausePlasma,
	"MOD_PLASMA_SPLASH":  DeathCausePlasmaSplash,
	"MOD_RAILGUN":        DeathCauseRailGun,
	"MOD_LIGHTNING":      DeathCauseLighting,
	"MOD_BFG":            DeathCauseBFG,
	"MOD_BFG_SPLASH":     DeathCauseBFGSplash,
	"MOD_WATER":          DeathCauseWater,
	"MOD_SLIME":          DeathCauseSlime,
	"MOD_LAVA":           DeathCauseLava,
	"MOD_CRUSH":          DeathCauseCrush,
	"MOD_TELEFRAG":       DeathCauseTelefrag,
	"MOD_FALLING":        DeathCauseFalling,
	"MOD_SUICIDE":        DeathCauseSuicide,
	"MOD_TARGET_LASER":   DeathCauseTargetLaser,
	"MOD_TRIGGER_HURT":   DeathCauseTriggerHurt,
	"MOD_NAIL":           DeathCauseNail,
	"MOD_CHAINGUN":       DeathCauseChainGun,
	"MOD_PROXIMITY_MINE": DeathCauseProximityMine,
	"MOD_KAMIKAZE":       DeathCauseKamikaze,
	"MOD_JUICED":         DeathCauseJuiced,
	"MOD_GRAPPLE":        DeathCauseGrapple,
}

type DeathCause string

type Match struct {
	ID      string
	Players map[string]*Player
	World   *Player
	Kills   []*Kill
}

type Player struct {
	Name  string
	Kills []*Kill
}

type Kill struct {
	Killer *Player
	Killed *Player
	Cause  DeathCause
}

func NewMatch(ID string) *Match {
	return &Match{
		ID:      ID,
		Players: make(map[string]*Player),
		World:   &Player{Name: WorldName},
	}
}

func (ref *Match) AddPlayer(name string) error {
	if name == WorldName {
		return errors.New(fmt.Sprintf("cannot add a player '%s'", WorldName))
	}
	if ref.Players[name] == nil {
		ref.Players[name] = &Player{
			Name: name,
		}
	}
	return nil
}

func (ref *Match) RegisterKill(killerName, killedName string, deathCause DeathCause) error {
	killer := ref.World
	if killerName != WorldName {
		var err error
		killer, err = ref.GetPlayer(killerName)
		if err != nil {
			return err
		}
	}

	killed, err := ref.GetPlayer(killedName)
	if err != nil {
		return err
	}

	if _, found := deathCauses[string(deathCause)]; !found {
		return errors.New(fmt.Sprintf("death cause '%s' not exists", deathCause))
	}

	kill := &Kill{killer, killed, deathCause}
	killer.Kills = append(killer.Kills, kill)
	ref.Kills = append(ref.Kills, kill)
	return nil
}

func (ref *Match) GetKillScore(playerName string) (int, error) {
	killScore := 0
	player, err := ref.GetPlayer(playerName)
	if err == nil {
		for _, kill := range player.Kills {
			if kill.Killed.Name != playerName {
				killScore++
			}
		}

		for _, kill := range ref.World.Kills {
			if kill.Killed.Name == playerName {
				killScore--
			}
		}
	}
	return killScore, err
}

func (ref *Match) GetPlayer(name string) (*Player, error) {
	player, found := ref.Players[name]
	if !found {
		return nil, errors.New(fmt.Sprintf("player '%s' not found in match", name))
	}
	return player, nil
}
