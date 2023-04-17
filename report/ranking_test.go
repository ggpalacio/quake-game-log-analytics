package report_test

import (
	_ "embed"
	"encoding/json"
	"github.com/ggpalacio/quake-game-log-analytics/report"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed ranking_report.json
var rankingReportFile []byte

func TestRankingReport_MarshalJSON(t *testing.T) {
	rankingReport := new(report.RankingReport)
	rankingReport.AddPlayerScore("foo", 100)
	rankingReport.AddPlayerScore("bar", 80)
	rankingReport.AddPlayerScore("baz", 120)

	rankingReportJSON, err := json.MarshalIndent(rankingReport, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, string(rankingReportFile), string(rankingReportJSON))
}
