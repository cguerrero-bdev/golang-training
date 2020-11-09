package main

import (
	"net/http"

	"github.com/cguerrero-bdev/golang-training/final-project/businesslogic"
	"github.com/cguerrero-bdev/golang-training/final-project/persistence"
	"github.com/cguerrero-bdev/golang-training/final-project/presentation"
	"github.com/gorilla/mux"
)

func handleRequests() {

	connection := persistence.GetDataBaseConnection()

	userRepository := persistence.UserRepository{
		Connection: connection,
	}

	questionRepository := persistence.QuestionRepository{
		Connection: connection,
	}

	questionManager := businesslogic.QuestionManager{
		QuestionRepository: questionRepository,
		UserRepository:     userRepository,
	}
	questionController := presentation.QuestionController{
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
