package gologger

import (
	"testing"
)

func TestColors(t *testing.T) {
	logger := New(Config{
		ColorLabel: true,
		ColorText:  true,
		LogLevel:   LogDebug,
	})

	logger.Debug("what %s This is ", "a log")
	logger.Info("what %s This is ", "a log")
	logger.Warning("what %s This is ", "a log")
	logger.Error("what %s This is ", "a log")
	logger.Fatal("what %s This is ", "a log")
	logger.Trace("what %s This is ", "a log")

}
