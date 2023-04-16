package logfile

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	logFile, err := New("../resources/valid_game.log")
	assert.NotNil(t, logFile)
	assert.NoError(t, err)

	logFile, err = New("../resources/not_found.log")
	assert.Nil(t, logFile)
	assert.ErrorIs(t, err, os.ErrNotExist)

	logFile, err = New("../resources/invalid_format.log")
	assert.Nil(t, logFile)
	assert.EqualError(t, err, "invalid log line format at line 6: 2023-04-16 DEBUG invalid log")
}

func TestLogLine_IsInitGame(t *testing.T) {
	logFile, err := New("../resources/valid_game.log")
	assert.NoError(t, err)
	assert.True(t, logFile.Lines[0].IsInitGame())
}

func TestLogLine_ClientUserinfoChanged(t *testing.T) {
	logFile, err := New("../resources/valid_game.log")
	assert.NoError(t, err)

	clientID, clientName := logFile.Lines[1].ClientUserinfoChanged()
	assert.Empty(t, clientID)
	assert.Zero(t, clientName)

	clientID, clientName = logFile.Lines[2].ClientUserinfoChanged()
	assert.Equal(t, 2, clientID)
	assert.Equal(t, "Isgalamido", clientName)
}

func TestLogLine_Kill(t *testing.T) {
	logFile, err := New("../resources/valid_game.log")
	assert.NoError(t, err)

	killer, killed, deathCause := logFile.Lines[6].Kill()
	assert.Empty(t, killer)
	assert.Empty(t, killed)
	assert.Empty(t, deathCause)

	killer, killed, deathCause = logFile.Lines[7].Kill()
	assert.Equal(t, "<world>", killer)
	assert.Equal(t, "Isgalamido", killed)
	assert.Equal(t, "MOD_TRIGGER_HURT", deathCause)
}
