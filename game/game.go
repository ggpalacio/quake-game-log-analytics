package game

const world = "<world>"

type Game struct {
	Players      []string       `json:"players"`
	TotalKills   int            `json:"total_kills"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}

func (ref *Game) AddPlayer(playerName string) {
	ref.Players = append(ref.Players, playerName)
}

func (ref *Game) RegisterKill(killer, killed, deathCause string) {
	if ref.Kills == nil {
		ref.Kills = make(map[string]int)
	}
	if ref.KillsByMeans == nil {
		ref.KillsByMeans = make(map[string]int)
	}

	if killer == world {
		ref.Kills[killed]--
	} else if killer == killed {
		// TODO what to do when a player kills itself?
	} else {
		ref.Kills[killer]++
	}
	ref.KillsByMeans[deathCause]++
	ref.TotalKills++
}
