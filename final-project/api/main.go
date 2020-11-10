package main

import (
	"net/http"

	"github.com/cguerrero-bdev/golang-training/final-project/api/controller"
	"github.com/cguerrero-bdev/golang-training/final-project/api/dao"
	"github.com/cguerrero-bdev/golang-training/final-project/api/service"
	"github.com/gorilla/mux"
)

func handleRequests() {

	connection := dao.GetDataBaseConnection()

	userRepository := dao.UserDao{
		Connection: connection,
	}

	questionRepository := dao.QuestionDao{
		Connection: connection,
	}

	questionManager := service.QuestionManager{
		QuestionRepository: questionRepository,
		UserRepository:     userRepository,
	}
	questionController := controller.QuestionController{
		QuestionManager: questionManager,
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
