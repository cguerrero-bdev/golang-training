package server

import (
	"log"
	"net/http"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/implementation/controller"
	"github.com/gorilla/mux"
)

type Server struct {
	questionController *controller.QuestionController
	infoLogger         *log.Logger
	errorLogger        *log.Logger
}

func New(questionController *controller.QuestionController, infoLogger *log.Logger, errorLogger *log.Logger) *Server {

	server := Server{
		questionController,
		infoLogger,
		errorLogger,
	}

	return &server
}

func (server *Server) HandleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/questions", server.questionController.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", server.questionController.GetQuestionById).Methods("GET")
	router.HandleFunc("/questions/user/{userName}", server.questionController.GetQuestionsByUserName).Methods("GET")

	router.HandleFunc("/questions", server.questionController.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", server.questionController.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", server.questionController.DeleteQuestion).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
