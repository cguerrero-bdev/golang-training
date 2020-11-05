package presentation

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

type jsonQuestion struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	UserName string `json:"username"`
}

func (jsonQuestion *jsonQuestion) GetId() int {
	result, _ := strconv.Atoi(jsonQuestion.Id)
	return result
}

func (jsonQuestion *jsonQuestion) GetText() string {
	return jsonQuestion.Text
}

func (jsonQuestion *jsonQuestion) GetUserName() string {
	return jsonQuestion.UserName
}

type QuestionController struct {
	QuestionManager QuestionManager
}

func (questionController *QuestionController) GetQuestions(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) GetQuestionById(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) GetQuestionByUserId(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var jsonQuestion jsonQuestion
	json.Unmarshal(reqBody, &jsonQuestion)

	questionController.QuestionManager.CreateQuestion(&jsonQuestion)
	json.NewEncoder(w).Encode(jsonQuestion)

}

func (questionController *QuestionController) UpdateQuestion(w http.ResponseWriter, r *http.Request) {

}

func (questionController *QuestionController) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

}
