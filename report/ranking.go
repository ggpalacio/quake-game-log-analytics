package report

import (
	"bytes"
	"fmt"
	"sort"
)

type RankingReport struct {
	scoreByPlayer ScoreByPlayer
}

type PlayerScore struct {
	Player string
	Score  int
}

type Ranking []PlayerScore

type ScoreByPlayer map[string]int

func (ref *RankingReport) AddPlayerScore(playerName string, score int) {
	if ref.scoreByPlayer == nil {
		ref.scoreByPlayer = make(map[string]int)
	}
	ref.scoreByPlayer[playerName] += score
}

func (ref *RankingReport) AddPlayersScore(scoreByPlayer ScoreByPlayer) {
	for playerName, score := range scoreByPlayer {
		ref.AddPlayerScore(playerName, score)
	}
}

func (ref *RankingReport) GetRanking() Ranking {
	var index int
	ranking := make(Ranking, len(ref.scoreByPlayer))
	for player, score := range ref.scoreByPlayer {
		ranking[index] = PlayerScore{player, score}
		index++
	}
	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].Score > ranking[j].Score
	})
	return ranking
}

func (ref *RankingReport) MarshalJSON() ([]byte, error) {
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
