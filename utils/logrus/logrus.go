package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	debugMode = false
)

func init() {
	debugMode = os.Getenv("DEBUG") == "1"
	if debugMode {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetOutput(os.Stdout)

}

func DebugMode() bool {
	return debugMode
}

type Fields logrus.Fields

func WithFields(fields Fields) *logrus.Entry {
	return logrus.WithFields(logrus.Fields(fields))
}

var (
	// Level
	SetLevel = logrus.SetLevel
	GetLevel = logrus.GetLevel

	// With info
	WithError = logrus.WithError
	WithField = logrus.WithField

	// Log
	Debug   = logrus.Debug
	Print   = logrus.Print
	Info    = logrus.Info
	Warn    = logrus.Warn
	Warning = logrus.Warning
	Error   = logrus.Error
	Panic   = logrus.Panic
	Fatal   = logrus.Fatal

	// Logf
	Debugf   = logrus.Debugf
	Printf   = logrus.Printf
	Println   = logrus.Println
	Infof    = logrus.Infof
	Warnf    = logrus.Warnf
	Warningf = logrus.Warningf
	Errorf   = logrus.Errorf
	Panicf   = logrus.Panicf
	Fatalf   = logrus.Fatalf
)
