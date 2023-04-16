package report

import (
	"fmt"
	"github.com/ggpalacio/quake-game-log-analytics/game"
	"github.com/ggpalacio/quake-game-log-analytics/logfile"
)

type Games map[string]*game.Game

func New(logFile *logfile.LogFile) Games {
	var currentGame *game.Game
	var currentGameIndex int
	games := make(Games)

	for _, logLine := range logFile.Lines {
		if logLine.IsInitGame() {
			currentGame = new(game.Game)
			currentGameIndex++

			gameKey := fmt.Sprintf("game_%s", fmt.Sprintf("%03d", currentGameIndex))
			games[gameKey] = currentGame
		}
		if currentGame == nil {
			continue
		}
		if playerID, playerName := logLine.ClientUserinfoChanged(); playerID != 0 {
			currentGame.AddPlayer(playerName)
			continue
		}
		if killer, killed, deathCause := logLine.Kill(); killer != "" {
			currentGame.RegisterKill(killer, killed, deathCause)
			continue
		}
	}
	return games
}
