package util

import (
	"fmt"
	"runtime"
)

type ApplicationError struct {
	Code, Message string
}

func GenerateApplicationErrorFromError(err error) *ApplicationError {

	fmt.Printf("Error: %s -> %v\n", "UpdateQuestion", err)

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s \n %v\n", file, line, f.Name(), pc)

	return &ApplicationError{}
}

func GenerateApplicationError() *ApplicationError {

	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s \n %v\n", file, line, f.Name(), pc)

	return &ApplicationError{}
}
