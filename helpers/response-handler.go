package helpers

import (
    "fmt"
    "github.com/dalecosta1/sinaloa-api/dto/messages/response"
    "os"
)

// Response represents the structure of a success response entity
type Response struct {
    Response bool        `json:"response"`
    Code     string      `json:"code"`
    Message  string      `json:"message"`
    Data     interface{} `json:"data"`
}

// HandleResponse returns a response entity as a struct
func HandleResponse(success bool, code string, message string, err error, data interface{}) response.Response {
    var errorMessage string

    if success == false {
        if os.Getenv("SINALOA_API_DEBUG") == "" || os.Getenv("SINALOA_API_DEBUG") == "false" {
            errorMessage = fmt.Sprintf("%s", err.Error())
        } else {
            errorMessage = fmt.Sprintf("%s --> %s", message, err.Error())
        }
        message = errorMessage
        data = struct{}{}
    }

    return response.NewResponse(success, code, message, data)
}
