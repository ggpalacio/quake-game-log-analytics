package report

import (
	"bytes"
	"fmt"
	"sort"
)

type RankingReport struct {
	playersScore PlayersScore
}

type PlayerScore struct {
	Player string
	Score  int
}

type Ranking []PlayerScore

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

func (ref *RankingReport) GetRanking() Ranking {
	var index int
	ranking := make(Ranking, len(ref.playersScore))
	for player, score := range ref.playersScore {
		ranking[index] = PlayerScore{player, score}
		index++
	}
	sort.Sort(ranking)
	return ranking
}

func (ref RankingReport) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{")

	for index, ranking := range ref.GetRanking() {
		if index > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`"%s":%d`, ranking.Player, ranking.Score))
	}

	buf.WriteString("}")
	return buf.Bytes(), nil
}

func (ref Ranking) Len() int {
	return len(ref)
}

func (ref Ranking) Less(i, j int) bool {
	return ref[i].Score > ref[j].Score
}

func (ref Ranking) Swap(i, j int) {
	tmp := ref[i]
	ref[i] = ref[j]
	ref[j] = tmp
}
