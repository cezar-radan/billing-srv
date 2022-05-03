package clog

import (
	"fmt"
	ilog "log"
	"os"
	"sync"
)

type Level struct {
	Name      string
	display   string
	Verbosity uint
}

var verbosity = WarnLevel.Verbosity
var mu sync.Mutex

var (
	NoneLevel    = Level{"NONE", "NONE", 0}
	ErrorLevel   = Level{"ERROR", "ERROR", 100}
	InfoLevel    = Level{"INFO", "INFO", 200}
	WarnLevel    = Level{"WARN", "WARN", 300}
	DebugLevel   = Level{"DEBUG", "DEBUG", 400}
	VerboseLevel = Level{"VERBOSE", "VERBOSE", 500}

	outputLevel = Level{"INFO", "OUTPUT", InfoLevel.Verbosity}
)

var levels = map[string]Level{
	NoneLevel.Name:    NoneLevel,
	ErrorLevel.Name:   ErrorLevel,
	InfoLevel.Name:    InfoLevel,
	WarnLevel.Name:    WarnLevel,
	DebugLevel.Name:   DebugLevel,
	VerboseLevel.Name: VerboseLevel,
}

func FindLevel(name string) (Level, error) {
	l, ok := levels[name]
	if !ok {
		return Level{}, fmt.Errorf("not such level %v", name)
	} else {
		return l, nil
	}
}

func SetLevel(level Level) {
	defer mu.Unlock()
	mu.Lock()
	verbosity = level.Verbosity
}

func Error(v ...interface{}) {
	log(ErrorLevel, v...)
}

func Errorf(format string, v ...interface{}) {
	logf(ErrorLevel, format, v...)
}

func Verbose(v ...interface{}) {
	log(VerboseLevel, v...)
}

func Verbosef(format string, v ...interface{}) {
	logf(VerboseLevel, format, v...)
}

func Debug(v ...interface{}) {
	log(DebugLevel, v...)
}

func Debugf(format string, v ...interface{}) {
	logf(DebugLevel, format, v...)
}

func Info(v ...interface{}) {
	log(InfoLevel, v...)
}

func Infof(format string, v ...interface{}) {
	logf(InfoLevel, format, v...)
}

func Warn(v ...interface{}) {
	log(WarnLevel, v...)
}

func Warnf(format string, v ...interface{}) {
	logf(WarnLevel, format, v...)
}

func Output(v ...interface{}) {
	if outputLevel.Verbosity > getVerbosity() {
		return
	}

	log(outputLevel, v...)
	// TODO: any race condition here?
	if ilog.Writer() != os.Stdout && ilog.Writer() != os.Stderr {
		fmt.Println(v...)
	}
}

func Outputf(format string, v ...interface{}) {
	if outputLevel.Verbosity > getVerbosity() {
		return
	}

	logf(outputLevel, format, v...)
	// TODO: any race condition here?
	if ilog.Writer() != os.Stdout && ilog.Writer() != os.Stderr {
		_, _ = fmt.Printf(format+"\n", v...)
	}
}

func log(level Level, v ...interface{}) {
	if level.Verbosity > getVerbosity() {
		return
	}

	fv := fmt.Sprintln(v...)
	ilog.Printf("%v: %v", level.display, fv)
}

func logf(level Level, format string, v ...interface{}) {
	if level.Verbosity > getVerbosity() {
		return
	}

	format = level.display + ": " + format + "\n"
	ilog.Printf(format, v...)
}

func getVerbosity() uint {
	defer mu.Unlock()

	mu.Lock()
	return verbosity
}
