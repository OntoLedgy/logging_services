package standard_global_logger

import (
	"github.com/OntoLedgy/logging_services/code/services/file_logging_services"
)

//var errorlog *os.File

var LoggingService *file_logging_services.FileLoggingServiceFactory

func Start_logger(
	log_folder_name,
	log_file_name_prefix string) *file_logging_services.FileLoggingServiceFactory {

	loggingService :=
		new(
			file_logging_services.FileLoggingServiceFactory)

	loggingService.Create(
		log_folder_name,
		log_file_name_prefix)

	return loggingService
}
