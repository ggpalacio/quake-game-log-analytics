package report

import (
	"sort"
)

type RankingReport struct {
	playersScore PlayersScore
}

type PlayerScore struct {
	Player string
	Score  int
}

type PlayerRanking []PlayerScore

type PlayersScore map[string]int

func (ref *RankingReport) AddPlayerScore(player string, score int) {
	if ref.playersScore == nil {
		ref.playersScore = make(map[string]int)
	}
	ref.playersScore[player] += score
}

func (ref *RankingReport) AddPlayersScore(playersScore PlayersScore) {
	for player, score := range playersScore {
		ref.AddPlayerScore(player, score)
	}
}

func (ref *RankingReport) GePlayerRanking() PlayerRanking {
	var index int
	ranking := make(PlayerRanking, len(ref.playersScore))
	for player, score := range ref.playersScore {
		ranking[index] = PlayerScore{player, score}
		index++
	}
	sort.Sort(ranking)
	return ranking
}

func (ref PlayerRanking) Len() int {
	return len(ref)
}

func (ref PlayerRanking) Less(i, j int) bool {
	return ref[i].Score > ref[j].Score
}

func (ref PlayerRanking) Swap(i, j int) {
	tmp := ref[i]
	ref[i] = ref[j]
	ref[j] = tmp
}
