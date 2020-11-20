package gologger

import (
	"fmt"
	"log"
)

// Colors
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	RedBold = "\033[31;1m"
	//Green       = "\033[32m"
	//White       = "\033[37m"
	//BlueBold    = "\033[34;1m"
	//MagentaBold = "\033[35;1m"
	//YellowBold  = "\033[33;1m"
)

// Levels
const (
	LogDebug = iota + 1
	LogInfo
	LogWarning
	LogError
	LogFatal
	LogTrace
)

type Config struct {
	ColorLabel bool
	ColorText  bool
	LogLevel   LogLevel
}

type LogLevel int

type LogInterface interface {
	//LogMode(level LogLevel) LogInterface
	WriteLog(string, string, string)
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Trace(string, ...interface{})
}

func New(config Config) LogInterface {
	return &epsLogger{
		Config: config,
	}
}

type epsLogger struct {
	Config
}

func (l *epsLogger) WriteLog(color string, label string, logString string) {
	if l.ColorLabel {
		label = color + label + Reset
	}
	if l.ColorText {
		logString = color + logString + Reset
	}

	log.Printf(label + logString)
}

func (l *epsLogger) Debug(logString string, args ...interface{}) {
	if l.LogLevel > LogDebug {
		return
	}
	label := "[ debug ] "
	color := Magenta

	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}

func (l *epsLogger) Info(logString string, args ...interface{}) {
	if l.LogLevel > LogInfo {
		return
	}
	label := "[ info ] "
	color := Blue
	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}

func (l *epsLogger) Warning(logString string, args ...interface{}) {
	if l.LogLevel > LogWarning {
		return
	}
	label := "[ warning ] "
	color := Yellow
	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}

func (l *epsLogger) Error(logString string, args ...interface{}) {
	if l.LogLevel > LogError {
		return
	}
	label := "[ error ] "
	color := Red
	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}

func (l *epsLogger) Fatal(logString string, args ...interface{}) {
	if l.LogLevel > LogFatal {
		return
	}
	label := "[ fatal ] "
	color := RedBold
	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}

func (l *epsLogger) Trace(logString string, args ...interface{}) {
	if l.LogLevel > LogTrace {
		return
	}
	label := "[ trace ] "
	color := Cyan
	logString = fmt.Sprintf(logString, args...)

	l.WriteLog(color, label, logString)

}
