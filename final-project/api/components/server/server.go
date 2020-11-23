package server

import (
	"log"
	"net/http"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/controller"
	"github.com/gorilla/mux"
)

type Server struct {
	questionController controller.QuestionController
	port               string
	infoLogger         *log.Logger
	errorLogger        *log.Logger
}

func NewServer(questionController controller.QuestionController, port string, infoLogger *log.Logger, errorLogger *log.Logger) *Server {

	server := Server{
		questionController,
		port,
		infoLogger,
		errorLogger,
	}

	return &server
}

func (server *Server) GetQuestions(w http.ResponseWriter, r *http.Request) {
	server.questionController.GetQuestions(w, r)
}
func (server *Server) GetQuestionById(w http.ResponseWriter, r *http.Request) {
	server.questionController.GetQuestionById(w, r)
}
func (server *Server) GetQuestionsByUserId(w http.ResponseWriter, r *http.Request) {
	server.questionController.GetQuestionsByUserId(w, r)
}
func (server *Server) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	server.questionController.CreateQuestion(w, r)
}
func (server *Server) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	server.questionController.UpdateQuestion(w, r)
}
func (server *Server) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	server.questionController.DeleteQuestion(w, r)
}

func (server *Server) HandleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/questions", server.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", server.GetQuestionById).Methods("GET")
	router.HandleFunc("/questions/user/{id}", server.GetQuestionsByUserId).Methods("GET")

	router.HandleFunc("/questions", server.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", server.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", server.DeleteQuestion).Methods("DELETE")

	http.ListenAndServe(":"+server.port, router)
}
