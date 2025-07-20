package controllers

import "fmt"

// CustomError represents a custom error with a code and message.
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Code %d: %s", e.Code, e.Message)
}

func db_connection_error() error {
	return &CustomError{Code: 23, Message: "Could not connect to the Database"}
}
