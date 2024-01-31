package errs

import "fmt"

type ErrorResponse struct {
	Errors struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
type ErrorObj struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type AppErrors struct {
	Code        string     `json:"Code,omitempty"`
	Message     string     `json:"Message,omitempty"`
	ErrorId     int        `json:"ErrorId,omitempty"`
	ErrorString error      `json:"Error,omitempty"`
	StatusCode  int        `json:",omitempty"`
	Errorsarray []ErrorObj `json:"Errors,omitempty"`
}

func ValidateError(Errorarray []ErrorObj, statusCode int) *AppErrors {

	fmt.Println(Errorarray, "entered error")
	return &AppErrors{
		Errorsarray: Errorarray,
		StatusCode:  statusCode,
	}
}

func ValidationErrors(errmsg string, code string) ErrorObj {
	return ErrorObj{
		Code:    code,
		Message: errmsg,
	}
}
