package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/controller"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/service"
	"github.com/gorilla/mux"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func handleRequests() {

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/questions", questionController.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", questionController.GetQuestionById).Methods("GET")
	router.HandleFunc("/questions/user/{userName}", questionController.GetQuestionsByUserName).Methods("GET")

	router.HandleFunc("/questions", questionController.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", questionController.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", questionController.DeleteQuestion).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}

func main() {

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	handleRequests()
}
