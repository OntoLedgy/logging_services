package loggrus_services

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"io"
)

type Logger struct {
	log_entry        *log.Entry
	log_file_handler *logfmt.Handler
}

func (logger *Logger) Prepare_log_entry(
	handler_file io.Writer,
	current_file string,
	current_function string,
) {

	logger.log_file_handler =
		logfmt.
			New(
				handler_file)

	log.SetHandler(
		logger.log_file_handler)

	logger.log_entry = log.
		WithFields(log.Fields{
			"file name": current_file,
			"function":  current_function,
		})

}

func (logger *Logger) Log_info(message string) {

	logger.log_entry.Info(message)

}

func (logger *Logger) Log_warning(message string) {

	logger.log_entry.Warn(message)

}

func (logger *Logger) Log_withError(message string) {

	//logging_services.log_entry.WithError(errors.New("unauthorized")).Error("upload failed")

}

func (logger *Logger) Log_Errorf(message string) {

	logger.log_entry.Errorf("Error: %s", message)

}
