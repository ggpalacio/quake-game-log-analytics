package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReport(t *testing.T) {
	logFile := &LogFile{
		Logs: []Log{
			{"20:37", `ShutdownGame:`},
			{"20:37", `InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\bot_minplayers\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0`},
			{"20:38", `ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\uriel/zael\hmodel\uriel/zael\g_redteam\\g_blueteam\\c1\5\c2\5\hc\100\w\0\l\0\tt\0\tl\0`},
			{"20:54", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"21:07", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"21:42", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"21:51", `ClientUserinfoChanged: 3 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0`},
			{"21:53", `ClientUserinfoChanged: 3 n\Mocinha\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0`},
			{"22:06", `Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH`},
			{"22:18", `Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH`},
			{"22:40", `Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH`},
			{"23:06", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"25:05", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"25:18", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"25:05", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"25:18", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
			{"25:41", `Kill: 1022 2 19: <world> killed Isgalamido by MOD_FALLING`},
			{"25:52", `Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`},
		},
	}

	report := NewReport(logFile)
	assert.Len(t, report, 1)

	match := report["game_001"]
	assert.Equal(t, []string{"Isgalamido", "Dono da Bola", "Mocinha"}, match.Players)
	assert.Equal(t, -9, match.Kills["Isgalamido"])
	assert.Zero(t, match.Kills["Dono da Bola"])
	assert.Zero(t, match.Kills["Mocinha"])
	assert.Equal(t, 13, match.TotalKills)
	assert.Equal(t, 9, match.KillsByMeans["MOD_TRIGGER_HURT"])
	assert.Equal(t, 3, match.KillsByMeans["MOD_ROCKET_SPLASH"])
	assert.Equal(t, 1, match.KillsByMeans["MOD_FALLING"])
}
