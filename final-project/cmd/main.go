package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cguerrero-bdev/golang-training/final-project/pkg/persistence"

	"github.com/cguerrero-bdev/golang-training/final-project/pkg/logic"
	"github.com/cguerrero-bdev/golang-training/final-project/pkg/presentation"
)

func handleRequests() {

	connection := persistence.GetDataBaseConnection()

	questionRepository := persistence.QuestionRepository{connection}
	questionManager := logic.QuestionManager{questionRepository}
	questionController := presentation.QuestionController{questionManager}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/questions", questionController.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", questionController.GetQuestionById).Methods("GET")
	router.HandleFunc("/questions/user/{id}", questionController.GetQuestionByUserId).Methods("GET")

	router.HandleFunc("/questions", questionController.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", questionController.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", questionController.DeleteQuestion).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}

func main() {

	handleRequests()
}
