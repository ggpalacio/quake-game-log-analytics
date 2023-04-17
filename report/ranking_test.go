package report_test

import (
	"encoding/json"
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRankingReport_GetRanking(t *testing.T) {
	rankingReport := new(report.RankingReport)
	rankingReport.AddPlayersScore(report.PlayersScore{
		"foo": 100,
		"bar": 80,
		"baz": 120,
	})
	ranking := rankingReport.GetRanking()
	assert.Len(t, ranking, 3)
	assert.Equal(t, report.PlayerScore{Player: "baz", Score: 120}, ranking[0])
	assert.Equal(t, report.PlayerScore{Player: "foo", Score: 100}, ranking[1])
	assert.Equal(t, report.PlayerScore{Player: "bar", Score: 80}, ranking[2])
}

func TestRankingReport_MarshalJSON(t *testing.T) {
	rankingReport := new(report.RankingReport)
	rankingReport.AddPlayerScore("foo", 100)
	rankingReport.AddPlayerScore("bar", 80)
	rankingReport.AddPlayerScore("baz", 120)

	rankingReportJSON, err := json.Marshal(rankingReport)
	assert.NoError(t, err)
	assert.Equal(t, `{"baz":120,"foo":100,"bar":80}`, string(rankingReportJSON))
}
