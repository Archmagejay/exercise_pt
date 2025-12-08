package logger

import (
	"log"
	"os"
)

type level int

const (
	LogInfo level = iota
	LogWarning
	LogError
	LogFatal
)

func InitLogger() (*log.Logger, error) {
	logpath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	logpath += "/log.txt"
	f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	logger := log.New(f, "", log.LstdFlags|log.Lmsgprefix)
	if logger == nil {
		log.Fatal("Creating new logger failed")
	}
	return logger, nil
}



func Log (l *log.Logger, severity level, msg any) {
	var str string

	if m, ok := msg.(string); ok {
		str = m
	} else if m, ok := msg.(error); ok {
		str = m.Error()
	} else {
		log.Fatal("Incompatible type")
	}


	switch severity {
	case LogInfo:
		l.Printf("INFO: %s", str)
	case LogWarning:
		l.Printf("WARNING: %s", str)
	case LogError:
		l.Printf("ERROR: %s", str)
	case LogFatal:
		l.Printf("FATAL ERROR: %s", str)
	default:
		l.Printf("UNKNOWN: %s", str)
	}
}