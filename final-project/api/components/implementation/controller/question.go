package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/service"
	"github.com/gorilla/mux"
)

type JsonQuestion struct {
	Id         string `json:"id"`
	Statement  string `json:"statement"`
	UserName   string `json:"username"`
	Answer     string `json:"answer"`
	AnsweredBy string `json:"answeredby"`
}

type Question struct {
	Id         int
	Statement  string
	UserName   string
	Answer     string
	AnsweredBy string
}

type QuestionController struct {
	QuestionManager service.QuestionManager
}

func (questionController *QuestionController) GetQuestions(w http.ResponseWriter, r *http.Request) {

	questions, _ := questionController.QuestionManager.GetQuestions()

	result := make([]JsonQuestion, 0)
	for _, question := range questions {
		result = append(result, createJsonQuestion(&question))
	}

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionController) GetQuestionById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	question, _ := questionController.QuestionManager.GetQuestionById(id)

	result := createJsonQuestion(question)

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionController) GetQuestionsByUserName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userName, _ := vars["userName"]

	questions, _ := questionController.QuestionManager.GetQuestionsByUserName(userName)

	result := make([]JsonQuestion, 0)
	for _, question := range questions {
		result = append(result, createJsonQuestion(&question))
	}

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var jsonQuestion JsonQuestion
	json.Unmarshal(reqBody, &jsonQuestion)
	id, _ := strconv.Atoi(jsonQuestion.Id)
	question := service.Question{Id: id, Statement: jsonQuestion.Statement, UserName: jsonQuestion.UserName}
	questionController.QuestionManager.CreateQuestion(&question)
	json.NewEncoder(w).Encode(jsonQuestion)

}

func (questionController *QuestionController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	reqBody, _ := ioutil.ReadAll(r.Body)
	var jsonQuestion JsonQuestion
	json.Unmarshal(reqBody, &jsonQuestion)

	question := service.Question{
		Id:         id,
		Statement:  jsonQuestion.Statement,
		UserName:   jsonQuestion.UserName,
		Answer:     jsonQuestion.Answer,
		AnsweredBy: jsonQuestion.AnsweredBy,
	}

	questionController.QuestionManager.UpdateQuestion(&question)
	json.NewEncoder(w).Encode(jsonQuestion)

}

func (questionController *QuestionController) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	questionController.QuestionManager.DeleteQuestion(id)
}

func createJsonQuestion(question *service.Question) JsonQuestion {

	result := JsonQuestion{
		Id:         strconv.Itoa(question.Id),
		Statement:  question.Statement,
		UserName:   question.UserName,
		Answer:     question.Answer,
		AnsweredBy: question.AnsweredBy,
	}

	return result
}
