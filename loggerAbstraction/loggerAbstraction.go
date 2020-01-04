package loggerAbstraction

import "bitbucket.org/turbosoftnetworks/netscope-api-v2/core"

type Logger interface {
	LPrintln(level int, v ...interface{})
	LPrintf(level int, format string, v ...interface{})
}

type RegisteredLogger struct {
	LogServer Logger
}

func NewRegisteredLogger(logger Logger) (l *RegisteredLogger) {
	l = &RegisteredLogger{logger}
	return
}

func (l RegisteredLogger) LPrintln(level int, v ...interface{}) {
	caller := core.MyCaller()
	v = append([]interface{}{caller}, v...)
	l.LogServer.LPrintln(level, v...)
}

func (l RegisteredLogger) LPrintf(level int, format string, v ...interface{}) {
	l.LogServer.LPrintf(level, core.MyCaller() + ": " + format, v...)
}