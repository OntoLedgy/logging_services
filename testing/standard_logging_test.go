package testing

import (
	"github.com/OntoLedgy/logging_services/standard_global_logger"
	"testing"
)

func TestStandardLogger(t *testing.T) {

	var log_file_path = "D:\\S\\go\\src\\github.com\\OntoLedgy\\logging_services\\testing"
	var log_file_prefix = "test"

	var logging_service = standard_global_logger.Start_logger(
		log_file_path,
		log_file_prefix)

	logger := logging_service.Global_logger

	logger.Print("hello")
}
