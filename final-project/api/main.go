package main

import (
	"log"
	"os"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/controller"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/server"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/service"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func main() {

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.NewServer(getQuestionController(), "3000", InfoLogger, ErrorLogger)
	server.HandleRequests()
}

func getQuestionController() *controller.QuestionControllerImplementation {

	connection := dao.GetDataBaseConnection()

	questionDao := dao.QuestionDaoImplementation{
		Connection:  connection,
		InfoLogger:  InfoLogger,
		ErrorLogger: ErrorLogger,
	}

	questionManager := service.QuestionManagerImplementation{
		QuestionDao: &questionDao,
		InfoLogger:  InfoLogger,
		ErrorLogger: ErrorLogger,
	}
	questionController := controller.QuestionControllerImplementation{
		QuestionManager: &questionManager,
		InfoLogger:      InfoLogger,
		ErrorLogger:     ErrorLogger,
	}

	return &questionController
}
