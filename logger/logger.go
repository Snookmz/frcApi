package logger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type Logger struct {
	DebugLevel int
	LogFile string
	TraceL *log.Logger
	InfoL *log.Logger
	WarningL *log.Logger
	ErrorL *log.Logger
}


func NewLogger (debugLevel int, logFile string) (l *Logger) {

	lh := &lumberjack.Logger{
				Filename:   logFile,
				MaxSize:    50, // megabytes
				MaxBackups: 5,
				MaxAge:     28, //days
				Compress:   true, // disabled by default
		}

	l = &Logger{debugLevel,
		logFile,
		log.New(lh, "TRACE: ", log.Ldate|log.Ltime),
		log.New(lh, "INFO: ", log.Ldate|log.Ltime),
		log.New(lh, "WARNING: ", log.Ldate|log.Ltime),
		log.New(lh, "ERROR: ", log.Ldate|log.Ltime),

	}
	return

}

func (l *Logger) LPrintln(level int, v ...interface{}) {
	var err error

	if level > l.DebugLevel {
		return
	}

	switch level {
	case 1 :
		err = l.ErrorL.Output(2, fmt.Sprintln(v...))
	case 2 :
		err = l.WarningL.Output(2, fmt.Sprintln(v...))
	case 3 :
		err = l.InfoL.Output(2, fmt.Sprintln(v...))
	case 4 :
		err = l.TraceL.Output(2, fmt.Sprintln(v...))
	}

	if err != nil {
		log.Printf("error with LPrintln: %s", err.Error())
	}

}

func (l *Logger) LPrintf(level int, format string, v ...interface{}) {
	var err error

	if level > l.DebugLevel {
		return
	}

	switch level {
	case 1 :
		err = l.ErrorL.Output(2, fmt.Sprintf(format, v...))
	case 2 :
		err = l.WarningL.Output(2, fmt.Sprintf(format, v...))
	case 3 :
		err = l.InfoL.Output(2, fmt.Sprintf(format, v...))
	case 4 :
		err = l.TraceL.Output(2, fmt.Sprintf(format, v...))
	}

	if err != nil {
		log.Printf("error with LPrintf: %s", err.Error())
	}

}
