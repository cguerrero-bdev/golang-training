package main

import (
	"net/http"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/controller"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/service"
	"github.com/gorilla/mux"
)

func handleRequests() {

	connection := dao.GetDataBaseConnection()

	userDao := dao.UserDao{
		Connection: connection,
	}

	questionDao := dao.QuestionDao{
		Connection: connection,
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

	handleRequests()
}
