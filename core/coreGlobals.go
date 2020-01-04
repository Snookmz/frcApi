package core

import "log"

var (
	DEBUG int
	VERSION string = "2.0"
	CONFIG string
)

var (
	Any     *log.Logger
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

const (
	AnyL     = iota // 0
	TraceL          // 1
	InfoL           // 2
	WarningL        // 3
	ErrorL          // 4
)

var Lognames = [5]string{
	"DELETEME",
	"TRACE",
	"INFO",
	"WARNING",
	"ERROR",
}

type NSLogger struct {
}

var Nslogger *NSLogger

