package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	//ResponseSuccess is response success status
	ResponseSuccess = "success"
	//ResponseFailed is response failed status
	ResponseFailed = "failed"

	// JSONUnmarshalError is used for unmarshal json data error message
	JSONUnmarshalError = "Body should be json object"
	// MissingField is missing field
	MissingField = "%s is missing"
	// InValidField is invalid field
	InValidField = "%s is invalid"
	// PermissionDenied is used for the requester request the data without permission
	PermissionDenied = "Permission Denied"
	// AuthenticationFailed is used for the requester request the data without permission
	// AuthenticationFailed = "Authentication Failed"
	AuthenticationFailed = "authentication.failed"
	// PermissionInvalidSession is used for invalid session error message
	PermissionInvalidSession = "Permission Denied"
)

type emptyResult struct{}

// RequestResult is used for http request response
type RequestResult struct {
	Error   error       `json:"error,omitempty"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ToJSON is used to serialize RequestResult to string
func (r RequestResult) ToJSON() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(b)
}

// NewRequestResult generate a response object.
func NewRequestResult(status, message string, data interface{}, err error) RequestResult {
	if err != nil {
		return RequestResult{
			Status:  status,
			Message: err.Error() + message,
			Data:    data,
			Error:   err,
		}
	}

	return RequestResult{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// NewSuccessResponse return a success response
func NewSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	if data == nil {
		w.Write([]byte(NewRequestResult(ResponseSuccess, message, emptyResult{}, nil).ToJSON()))
		return
	}
	w.Write([]byte(NewRequestResult(ResponseSuccess, message, data, nil).ToJSON()))
}

// NewErrorResponse return a error response with a error message
func NewErrorResponse(w http.ResponseWriter, err error) {
	http.Error(w, NewRequestResult(ResponseFailed, "", emptyResult{}, err).ToJSON(), http.StatusBadRequest)
}

// NewErrorBodyJSON return a json unmarshal failed error
func NewErrorBodyJSON(w http.ResponseWriter) {
	http.Error(w, NewRequestResult(ResponseFailed, JSONUnmarshalError, emptyResult{}, nil).ToJSON(), http.StatusBadRequest)
}

// NewErrorMissField return a miss filed response error
func NewErrorMissField(w http.ResponseWriter, filed string) {
	http.Error(w, NewRequestResult(ResponseFailed, fmt.Sprintf(MissingField, filed), emptyResult{}, nil).ToJSON(), http.StatusBadRequest)
}

// NewErrorInvalidField return a miss filed response error
func NewErrorInvalidField(w http.ResponseWriter, filed string) {
	http.Error(w, NewRequestResult(ResponseFailed, fmt.Sprintf(InValidField, filed), emptyResult{}, nil).ToJSON(), http.StatusBadRequest)
}
