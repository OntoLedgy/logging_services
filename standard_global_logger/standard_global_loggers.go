package standard_global_logger

import (
	"github.com/OntoLedgy/logging_services/code/services/file_logging_services"
)

//var errorlog *os.File

var LoggingService *file_logging_services.FileLoggingServiceFactory

func Start_logger(
	log_folder_name,
	log_file_name_prefix string) *file_logging_services.FileLoggingServiceFactory {

	LoggingService =
		new(
			file_logging_services.FileLoggingServiceFactory)

	LoggingService.Create(
		log_folder_name,
		log_file_name_prefix)

	return LoggingService
}
