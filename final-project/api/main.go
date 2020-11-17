package main

import (
	"log"
	"os"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/controller"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/service"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/server"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func main() {

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.New(getQuestionController(), InfoLogger, ErrorLogger)
	server.HandleRequests()
}

func getQuestionController() *controller.QuestionController {

	connection := dao.GetDataBaseConnection()

	userDao := dao.UserDao{
		Connection:  connection,
		InfoLogger:  InfoLogger,
		ErrorLogger: ErrorLogger,
	}

	questionDao := dao.QuestionDao{
		Connection:  connection,
		InfoLogger:  InfoLogger,
		ErrorLogger: ErrorLogger,
	}

	questionManager := service.QuestionManager{
		QuestionDao: &questionDao,
		UserDao:     &userDao,
	}
	questionController := controller.QuestionController{
		QuestionManager: &questionManager,
	}

	return &questionController
}
