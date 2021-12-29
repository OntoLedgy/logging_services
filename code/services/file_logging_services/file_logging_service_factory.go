package file_logging_services

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//TODO add contract and wrappers here - move to code folder and hide implementation details.

type FileLoggingServiceFactory struct {
	Global_logger *log.Logger
}

func (file_logging_service FileLoggingServiceFactory) Create(
	log_folder_name,
	log_file_name_prefix string) {

	file_logging_service.directory_cleanup(
		log_folder_name)

	file_logging_service.initialise_logger(
		log_folder_name,
		log_file_name_prefix)

}

func (
	file_logging_service FileLoggingServiceFactory) Log_info(
	message interface{}) {

	file_logging_service.Global_logger.Print(message)
}

func (
	file_logging_service FileLoggingServiceFactory) Log_info_formated(
	format string,
	data interface{}) {

	file_logging_service.Global_logger.Printf(format, data)
}

func (
	file_logging_service FileLoggingServiceFactory) End_logger() {

}

func (
	file_logging_service FileLoggingServiceFactory) initialise_logger(
	baseFilePath,
	log_file_name_prefix string) {

	baseFilePath = strings.TrimRight(baseFilePath, "/")
	currentDate := time.Now().UTC()
	dateDirectory := time.Now().UTC().Format("2006-01-02")
	dateFile := currentDate.Format("2006-01-02T15-04-05")

	filePath := fmt.Sprintf("%s/%s/", baseFilePath, dateDirectory)
	fileName := strings.Replace(fmt.Sprintf("%s_%s.log", log_file_name_prefix, dateFile), " ", "-", -1)

	err := os.MkdirAll(filePath, os.ModePerm)

	if err != nil {
		log.Fatalf("main : Start : Failed to Create log directory : %s : %s\n", filePath, err)
	}

	logf, err := os.Create(fmt.Sprintf("%s%s", filePath, fileName))
	if err != nil {
		log.Fatalf("main : Start : Failed to Create log file : %s : %s\n", fileName, err)
	}

	file_logging_service.Global_logger = log.New(logf, "applog: ", log.Lshortfile|log.LstdFlags)

}

func (
	file_logging_service FileLoggingServiceFactory) directory_cleanup(
	baseFilePath string) {

	fileInfos, _ := ioutil.ReadDir(baseFilePath)

	// Create the date to compare for directories to remove.
	currentDate := time.Now().UTC()
	compareDate := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day()-1, 0, 0, 0, 0, time.UTC)

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() == false {
			continue
		}

		// The file name look like: YYYY-MM-DD
		parts := strings.Split(fileInfo.Name(), "-")

		year, _ := strconv.Atoi(parts[0])
		month, _ := strconv.Atoi(parts[1])
		day, _ := strconv.Atoi(parts[2])

		// The directory to check.
		fullFileName := fmt.Sprintf("%s/%s", baseFilePath, fileInfo.Name())

		// Create a time type from the directory name.
		directoryDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		// Compare the dates and convert to days.
		daysOld := int(compareDate.Sub(directoryDate).Hours() / 24)

		if daysOld >= 0 {
			os.RemoveAll(fullFileName)
		}
	}

}
