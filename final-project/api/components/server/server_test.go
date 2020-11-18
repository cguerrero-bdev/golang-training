package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
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

	request, err := http.NewRequest("GET", "/questions", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	err = json.NewEncoder(response).Encode(want)
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)
	mockQuestionController.EXPECT().
		GetQuestions(response, request).
		Return()

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	handler := http.HandlerFunc(server.GetQuestions)

	handler.ServeHTTP(response, request)

	resp := response.Result()
	body, err := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := []controller.JsonQuestion{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(result, want))
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

	request, err := http.NewRequest("GET", "/questions/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	json.NewEncoder(response).Encode(want)

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)
	mockQuestionController.EXPECT().
		GetQuestionById(response, request).
		Return()

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	handler := http.HandlerFunc(server.GetQuestionById)
	handler.ServeHTTP(response, request)

	resp := response.Result()
	body, err := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	result := controller.JsonQuestion{}
	err = json.Unmarshal(body, &result)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(result, want))
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

	response := httptest.NewRecorder()
	q, err := json.Marshal(question)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/questions", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)
	mockQuestionController.EXPECT().
		CreateQuestion(response, request).
		Return()

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

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

	response := httptest.NewRecorder()
	q, err := json.Marshal(question)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("PUT", "/questions/4", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)
	mockQuestionController.EXPECT().
		UpdateQuestion(response, request).
		Return()

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

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

	response := httptest.NewRecorder()
	q, err := json.Marshal(question)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("PUT", "/questions/4", bytes.NewReader(q))
	if err != nil {
		t.Fatal(err)
	}

	mockQuestionController := mockcontroller.NewMockQuestionController(ctrl)
	mockQuestionController.EXPECT().
		DeleteQuestion(response, request).
		Return()

	server := NewServer(mockQuestionController, "3000", InfoLogger, ErrorLogger)

	handler := http.HandlerFunc(server.DeleteQuestion)
	handler.ServeHTTP(response, request)

	resp := response.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
