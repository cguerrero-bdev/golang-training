package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/model"
	"github.com/cguerrero-bdev/golang-training/final-project/api/components/service"
	"github.com/gorilla/mux"
)

type QuestionController interface {
	GetQuestions(w http.ResponseWriter, r *http.Request)
	GetQuestionById(w http.ResponseWriter, r *http.Request)
	GetQuestionsByUserId(w http.ResponseWriter, r *http.Request)
	CreateQuestion(w http.ResponseWriter, r *http.Request)
	UpdateQuestion(w http.ResponseWriter, r *http.Request)
	DeleteQuestion(w http.ResponseWriter, r *http.Request)
}

type QuestionControllerImplementation struct {
	QuestionManager service.QuestionManager
	InfoLogger      *log.Logger
	ErrorLogger     *log.Logger
}

func (questionController *QuestionControllerImplementation) GetQuestions(w http.ResponseWriter, r *http.Request) {

	result, _ := questionController.QuestionManager.GetQuestions()
	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionControllerImplementation) GetQuestionById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	result, _ := questionController.QuestionManager.GetQuestionById(id)

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionControllerImplementation) GetQuestionsByUserId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	result, _ := questionController.QuestionManager.GetQuestionsByUserId(userId)

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionControllerImplementation) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	question := model.Question{}
	json.Unmarshal(reqBody, &question)
	questionController.QuestionManager.CreateQuestion(&question)
	json.NewEncoder(w).Encode(question)

}

func (questionController *QuestionControllerImplementation) UpdateQuestion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	reqBody, _ := ioutil.ReadAll(r.Body)
	question := model.Question{}
	json.Unmarshal(reqBody, &question)
	question.Id = id
	questionController.QuestionManager.UpdateQuestion(&question)
	json.NewEncoder(w).Encode(question)

}

func (questionController *QuestionControllerImplementation) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	questionController.QuestionManager.DeleteQuestion(id)
}
