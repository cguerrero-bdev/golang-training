package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getQuestions(w http.ResponseWriter, r *http.Request) {

}

func getQuestionById(w http.ResponseWriter, r *http.Request) {

}

func getQuestionByUserId(w http.ResponseWriter, r *http.Request) {

}

func createQuestion(w http.ResponseWriter, r *http.Request) {

}

func updateQuestion(w http.ResponseWriter, r *http.Request) {

}

func deleteQuestion(w http.ResponseWriter, r *http.Request) {

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/questions", getQuestions).Methods("GET")
	router.HandleFunc("/questions/{id}", getQuestionById).Methods("GET")
	router.HandleFunc("/questions/user/{id}", getQuestionByUserId).Methods("GET")

	router.HandleFunc("/questions", createQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", updateQuestion).Methods("PUT")
	router.HandleFunc("/questions/{id}", deleteQuestion).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}

func main() {

	handleRequests()
}
