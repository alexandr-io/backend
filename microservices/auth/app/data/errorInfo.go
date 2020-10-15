package data

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"runtime"
)

// ErrorInfo is used to return error that contain info about where it happened
type ErrorInfo struct {
	Message       string `json:"message"`
	File          string `json:"file"`
	Line          int    `json:"line"`
	ContentType   string `json:"content-type"`
	CustomMessage string `json:"custom-message"`
}

// NewErrorInfo set the info of the error in the returned ErrorInfo
func NewErrorInfo(message string, depth int) ErrorInfo {
	errorInfo := ErrorInfo{
		Message:     message,
		ContentType: fiber.MIMETextHTMLCharsetUTF8,
	}
	// jump to the 3rd function call to get the context of the error
	_, fn, line, _ := runtime.Caller(depth + 1)
	errorInfo.File = fn
	errorInfo.Line = line
	return errorInfo
}

// MarshalErrorInfo marshal a given ErrorInfo
func (errorInfo *ErrorInfo) MarshalErrorInfo() string {
	errorInfoString, _ := json.Marshal(errorInfo)
	return string(errorInfoString)
}

// NewErrorInfoMarshal create and marshal an ErrorInfo
func NewErrorInfoMarshal(message string, depth int) string {
	errorInfo := NewErrorInfo(message, depth+1)
	return errorInfo.MarshalErrorInfo()
}

// NewHttpErrorInfo return a fiber error containing a ErrorInfo JSON as message
func NewHttpErrorInfo(code int, message string) error {
	contentTypeValue := fiber.MIMETextHTMLCharsetUTF8
	return fiber.NewError(code, NewErrorInfoMarshal(message, 1), contentTypeValue)
}

// ErrorInfoUnmarshal unmarshal the errorInfoString into an ErrorInfo
func ErrorInfoUnmarshal(errorInfoString string) (*ErrorInfo, error) {
	var errorInfo ErrorInfo
	err := json.Unmarshal([]byte(errorInfoString), &errorInfo)
	if err != nil {
		return nil, err
	}
	return &errorInfo, nil
}
