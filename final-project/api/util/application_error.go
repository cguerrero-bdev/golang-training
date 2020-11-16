package util

import (
	"log"
)

type ApplicationError interface {
	Code() string
	Message() string
}

type ApplicationErrorImp struct {
	code    string
	message string
}

func (applicationErrorImp *ApplicationErrorImp) Code() string {
	return applicationErrorImp.code
}

func (applicationErrorImp *ApplicationErrorImp) Message() string {
	return applicationErrorImp.message
}

const UNKNOWN_ERROR = "UNKNOWN_ERROR"

var UnknownError = ApplicationErrorImp{
	code:    UNKNOWN_ERROR,
	message: "Unknown Error",
}

func GenerateApplicationUnknownError(err error, errorLogger *log.Logger) ApplicationError {

	errorLogger.Println(err.Error())

	return &UnknownError
}
