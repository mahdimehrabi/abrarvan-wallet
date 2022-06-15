package infrastracture

import (
	"log"
	"os"
)

type ArvanLogger struct {
	LG *log.Logger
}

func (l *ArvanLogger) Error(err string) {
	l.LG.Print(err)
}

func NewLogger() ArvanLogger {
	lg := log.New(os.Stdout, "arvan ", log.LstdFlags)
	return ArvanLogger{LG: lg}
}
