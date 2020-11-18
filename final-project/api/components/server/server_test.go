package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cguerrero-bdev/golang-training/final-project/api/components/definition/controller"
	mockcontroller "github.com/cguerrero-bdev/golang-training/final-project/api/components/mock/controller"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetQuestions(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	want := []controller.JsonQuestion{
		{
			Id:         "2",
			Statement:  "Question2",
			UserName:   "user1",
			Answer:     "answere 1",
			AnsweredBy: "user2",
		},
		{
			Id:         "3",
			Statement:  "Question3",
			UserName:   "user1",
			Answer:     "answere 3",
			AnsweredBy: "user2",
		},
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	request, err := http.NewRequest("GET", "/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	json.NewEncoder(response).Encode(want)

	mockQuestionController.EXPECT().
		GetQuestions(response, request).
		Return()

	handler := http.HandlerFunc(server.GetQuestions)

	handler.ServeHTTP(response, request)

	resp := response.Result()
	body, err := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := []controller.JsonQuestion{}
	err = json.Unmarshal(body, &result)

	assert.NoError(t, err)
	assert.Len(t, result, len(want))

	for i, _ := range want {

		assert.Equal(t, want[i].Id, result[i].Id)
		assert.Equal(t, want[i].Statement, result[i].Statement)
		assert.Equal(t, want[i].Answer, result[i].Answer)

	}

}

func TestGetQuestionById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	want := controller.JsonQuestion{

		Id:         "3",
		Statement:  "Question3",
		UserName:   "user1",
		Answer:     "answere 3",
		AnsweredBy: "user2",
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	request, err := http.NewRequest("GET", "/questions/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	json.NewEncoder(response).Encode(want)

	mockQuestionController.EXPECT().
		GetQuestionById(response, request).
		Return()

	handler := http.HandlerFunc(server.GetQuestionById)

	handler.ServeHTTP(response, request)

	resp := response.Result()
	body, err := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := controller.JsonQuestion{}
	err = json.Unmarshal(body, &result)

	assert.NoError(t, err)
	assert.Equal(t, want.Id, result.Id)
	assert.Equal(t, want.Statement, result.Statement)
	assert.Equal(t, want.Answer, result.Answer)

}

func TestCreateQuestion(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	question := controller.JsonQuestion{
		Id:        "4",
		Statement: "Question4",
		UserName:  "user1",
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	response := httptest.NewRecorder()
	q, _ := json.Marshal(question)

	request, err := http.NewRequest("POST", "/questions", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController.EXPECT().
		CreateQuestion(response, request).
		Return()

	handler := http.HandlerFunc(server.CreateQuestion)

	handler.ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestUpdateQuestion(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	question := controller.JsonQuestion{
		Id:        "4",
		Statement: "Question4",
		UserName:  "user1",
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	response := httptest.NewRecorder()
	q, _ := json.Marshal(question)

	request, err := http.NewRequest("PUT", "/questions/4", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController.EXPECT().
		UpdateQuestion(response, request).
		Return()

	handler := http.HandlerFunc(server.UpdateQuestion)

	handler.ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestDeleteQuestion(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	InfoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	question := controller.JsonQuestion{
		Id:        "4",
		Statement: "Question4",
		UserName:  "user1",
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	response := httptest.NewRecorder()
	q, _ := json.Marshal(question)

	request, err := http.NewRequest("PUT", "/questions/4", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController.EXPECT().
		DeleteQuestion(response, request).
		Return()

	handler := http.HandlerFunc(server.DeleteQuestion)

	handler.ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
