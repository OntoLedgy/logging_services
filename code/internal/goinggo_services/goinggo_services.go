package goinggo_services

import (
	"fmt"
	"github.com/goinggo/tracelog"
	"os"
	"runtime"
)

type Logger struct {
	Base_file_path string
	Logging_level  int32
	Log_file       *os.File
	Log_file_name  string
}

const (
	Trace_level = tracelog.LevelTrace
	Info_Level  = tracelog.LevelInfo
)

func Create_logger(
	base_file_path string,
	logging_level int32) *Logger {

	logger := new(Logger)

	logger.Base_file_path = base_file_path

	logger.Logging_level = logging_level

	logger.StartFile(logger.Logging_level)

	return logger

}

func (logger *Logger) StartFile(logging_level int32) {

	tracelog.StartFile(logging_level, logger.Base_file_path, 1)

}

func (logger *Logger) Start(trace_level int32) {

	tracelog.Start(trace_level)

}

func (logger *Logger) Started() {

	tracelog.Started("main", "main")

}

func (logger *Logger) Completed() {

	tracelog.Completed("main", "main")

}

func (logger *Logger) Trace(title string, function_name string, format string) {
	tracelog.Trace("main", "main", "Hello Trace")

}

func (logger *Logger) Info(log_string_format string) { // rename these

	_, filename, _, _ := runtime.Caller(1)

	logger.Set_file_name(filename)

	pc := make([]uintptr, 10) // at least 1 entry needed

	runtime.Callers(2, pc)

	f := runtime.FuncForPC(pc[0])

	function_name := f.Name()

	tracelog.Info(logger.Log_file_name, function_name, log_string_format)
}

func (logger *Logger) Warning() {
	tracelog.Warning("main", "main", "Hello Warn")
}

func (logger *Logger) Errorf() {
	tracelog.Errorf(fmt.Errorf("Exception At..."), "main", "main", "Hello Error")
}

func (logger *Logger) Stop() {

	tracelog.Stop()
}

func (logger *Logger) Set_file_name(file_name string) {

	logger.Log_file_name = file_name
}

func Start_logger() *Logger {

	logger :=

		Create_logger(
			"./outputs/logs",
			Info_Level)

	logger.
		Started()

	return logger
}

func End_logger(
	logger *Logger) {

	logger.Completed()

	logger.Stop()

}
