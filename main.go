package main

import (
	"github.com/OntoLedgy/logging_services/standard_global_logger"
)

func main() {

	logger := standard_global_logger.Global_logger

	logger.Print("hello")

}
