package report_test

import (
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRankingReport_GetPlayerRanking(t *testing.T) {
	rankingReport := new(report.RankingReport)
	rankingReport.AddPlayersScore(report.PlayersScore{
		"foo": 100,
		"bar": 80,
		"baz": 120,
	})
	ranking := rankingReport.GePlayerRanking()
	assert.Len(t, ranking, 3)
	assert.Equal(t, report.PlayerScore{Player: "baz", Score: 120}, ranking[0])
	assert.Equal(t, report.PlayerScore{Player: "foo", Score: 100}, ranking[1])
	assert.Equal(t, report.PlayerScore{Player: "bar", Score: 80}, ranking[2])
}
