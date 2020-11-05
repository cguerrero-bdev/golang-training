package presentation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cguerrero-bdev/golang-training/final-project/pkg/logic"
)

/*
type Question interface {
	GetId() int
	GetText() string
	GetUserName() string
}


type QuestionManager interface {
	GetQuestions()
	GetQuestionById(id int)
	GetQuestionByUserId(id int)
	CreateQuestion(question Question)
	UpdateQuestion(question Question)
	DeleteQuestionById(id int)
}
*/

type JsonQuestion struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	UserName string `json:"username"`
}

type QuestionController struct {
	QuestionManager logic.QuestionManager
}

func (questionController *QuestionController) GetQuestions(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) GetQuestionById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	question, _ := questionController.QuestionManager.GetQuestionById(id)

	result := JsonQuestion{
		Id:       strconv.Itoa(question.Id),
		Text:     question.Text,
		UserName: question.UserName,
	}

	json.NewEncoder(w).Encode(result)

}

func (questionController *QuestionController) GetQuestionByUserId(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var jsonQuestion JsonQuestion
	json.Unmarshal(reqBody, &jsonQuestion)
	id, _ := strconv.Atoi(jsonQuestion.Id)
	question := logic.Question{Id: id, Text: jsonQuestion.Text, UserName: jsonQuestion.UserName}
	questionController.QuestionManager.CreateQuestion(question)
	json.NewEncoder(w).Encode(jsonQuestion)

}

func (questionController *QuestionController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

}
