// provides logging at different levels
// TRACE 0
// INFO 1
// WARNING 2
// ERROR 3
package tracer

import (
	"io"
	"io/ioutil"
	"log"
)

type LogLevel int

const (
	TRACE LogLevel = iota
	INFO
	WARNING
	ERROR
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func Initialize(handle io.Writer, logLevel LogLevel) {
	trace, info, warning, err := ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard
	if logLevel <= ERROR {
		err = handle
	}
	if logLevel <= WARNING {
		warning = handle
	}
	if logLevel <= INFO {
		info = handle
	}
	if logLevel == TRACE {
		trace = handle
	}

	Error = log.New(
		err,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Warning = log.New(
		warning,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Info = log.New(
		info,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	Trace = log.New(
		trace,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
