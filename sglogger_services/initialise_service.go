package sglogger_services

import (
	"github.com/jaloren/sglogger"
	"os"
)

var globalLog = sglogger.GlobalLogger

func Start(log_directory string) {

	current_directory, _ := os.Getwd()

	logdir := current_directory + log_directory
	filePrefix := "test"

	err := globalLog.SetHandlers(logdir, filePrefix)

	if err != nil {
		panic(err)
	}

	globalLog.SetLogLevel("INFO")

	globalLog.Freeze(true)
}

func Info(message string) {

	var globalLog = sglogger.GlobalLogger
	globalLog.Error("Error message")
	globalLog.Warning("Info message")

}

func Test_service() {

}
