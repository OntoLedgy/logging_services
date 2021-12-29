package standard_global_logger

import (
	"github.com/OntoLedgy/logging_services/code/services/file_logging_services"
)

//var errorlog *os.File

func Start_logger(
	log_folder_name,
	log_file_name_prefix string) *file_logging_services.FileLoggingServiceFactory {

	logging_service :=
		new(
			file_logging_services.FileLoggingServiceFactory)

	logging_service.Create(
		log_folder_name,
		log_file_name_prefix)
	return logging_service
}
