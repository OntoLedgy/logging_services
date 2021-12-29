package loggrus_services

import (
	"github.com/sirupsen/logrus"
	"io"
	logging "log"
	"os"
)

var (
	logger *logrus.Logger
)

func init() {

	f, err := os.OpenFile("run_log.logging_services", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		logging.Fatalf("error opening file: %v", err)
	}

	logger = logrus.New()

	//logging_services.Formatter = &logrus.JSONFormatter{}

	logger.SetReportCaller(true)

	mw := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(mw)
}

// Info ...
func Info(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warn ...
func Warn(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Error ...
func Error(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

var (

	// ConfigError ...
	ConfigError = "%v type=config.error"

	// HTTPError ...
	HTTPError = "%v type=http.error"

	// HTTPWarn ...
	HTTPWarn = "%v type=http.warn"

	// HTTPInfo ...
	HTTPInfo = "%v type=http.info"
)
