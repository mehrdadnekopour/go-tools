package mypes

import (
	"errors"
	"net/http"
)

// Merror ...
type Merror struct {
	Problem    bool      `json:"-"`
	Data       error     `json:"data"`
	Code       ErrorCode `json:"code"`
	HTTPStatus int       `json:"http_status"`
}

// ErrorCode ...
type ErrorCode int

const (
	// Custom Errors ----------------------------------------

	// OpenFileFailed ....
	OpenFileFailed ErrorCode = 1
	// InvalidDataType ....
	InvalidDataType ErrorCode = 2
	// InvalidValue ....
	InvalidValue ErrorCode = 3

	// InvalidID ...
	InvalidID ErrorCode = 10
	// InvalidQueryParams ...
	InvalidQueryParams ErrorCode = 11
	// InvalidUserPass ...
	InvalidUserPass ErrorCode = 12

	// Database Errors --------------------------------------

	// DBConnectionFailed ...
	DBConnectionFailed ErrorCode = 50
	// DBRequestFailed ...
	DBRequestFailed ErrorCode = 51

	// NotCreated ...
	NotCreated ErrorCode = 100
	// NotFound ...
	NotFound ErrorCode = 101
	// NotUpdated ...
	NotUpdated ErrorCode = 102
	// NotDeleted ...
	NotDeleted ErrorCode = 103

	// Http Errors ------------------------------------------

	// HTTPBadRequest ...
	HTTPBadRequest ErrorCode = http.StatusBadRequest
	// HTTPUnauthorized ...
	HTTPUnauthorized ErrorCode = http.StatusUnauthorized
	// HTTPForbidden ...
	HTTPForbidden ErrorCode = http.StatusForbidden
	// HTTPNotFound ...
	HTTPNotFound ErrorCode = http.StatusNotFound
	// HTTPMethodNotAllowed ...
	HTTPMethodNotAllowed ErrorCode = http.StatusMethodNotAllowed
	// HTTPConflict ...
	HTTPConflict ErrorCode = http.StatusConflict
	// HTTPUnprocessableEntity ...
	HTTPUnprocessableEntity ErrorCode = http.StatusUnprocessableEntity
	// HTTPInternalServerError ...
	HTTPInternalServerError ErrorCode = http.StatusInternalServerError
	// HTTPGatewayTimeout ...
	HTTPGatewayTimeout ErrorCode = http.StatusGatewayTimeout
)

// Set ...
func (m *Merror) Set(problem bool, err error, code ErrorCode) {
	m.Problem = problem
	m.Data = err
	m.Code = code
	m.HTTPStatus = m.GetHTTPStatus()
}

// SetInvalidIDMerror ...
func (m *Merror) SetInvalidIDMerror() {
	m.Problem = true
	m.Data = errors.New("invalid id")
	m.Code = InvalidID
	m.HTTPStatus = m.GetHTTPStatus()
}

// SetInvalidUserPassMerror ...
func (m *Merror) SetInvalidUserPassMerror() {
	m.Problem = true
	m.Data = errors.New("invalid username/password")
	m.Code = InvalidUserPass
	m.HTTPStatus = m.GetHTTPStatus()
}

// SetInvalidQueryParamMrror ....
func (m *Merror) SetInvalidQueryParamMrror() {
	m.Problem = true
	m.Data = errors.New("invalid query params")
	m.Code = InvalidQueryParams
	m.HTTPStatus = m.GetHTTPStatus()
}

// SetAccessDeniedError ...
func (m *Merror) SetAccessDeniedError() {
	m.Problem = true
	m.Data = errors.New("access denied")
	m.Code = HTTPForbidden
}

// GetMessage ...
func (m *Merror) GetMessage() (msg string) {

	switch m.Code {
	case OpenFileFailed:
		msg = "oppening file failed"
	case InvalidDataType:
		msg = "casting failed, invalid data type"
	case InvalidValue:
		msg = "invalid value"
	case InvalidID:
		msg = "invalid id"
	case InvalidQueryParams:
		msg = "invalid query params"
	case InvalidUserPass:
		msg = "invalid username/password"
	case DBConnectionFailed:
		msg = "Failed to connecct to DB"
	case HTTPInternalServerError:
		msg = "Internal server error"
	case HTTPBadRequest:
		msg = "Bad request"
	case HTTPForbidden:
		msg = "Forbidden"
	case HTTPNotFound:
		msg = "Not found"
	case NotFound:
		msg = "Not found"
	case NotUpdated:
		msg = "Not updated"
	case NotCreated:
		msg = "Not created"
	case NotDeleted:
		msg = "Not Deleted"
	case HTTPUnprocessableEntity:
		msg = "Unprocessable entity"
	case HTTPUnauthorized:
		msg = "Unauthorized"
	case HTTPMethodNotAllowed:
		msg = "MethodNotAllowed"
	case HTTPConflict:
		msg = "Conflict"
	case http.StatusOK:
		msg = "Ok"
	default:
		msg = "Internal server error"
	}

	return
}

// GetHTTPStatus ...
func (m *Merror) GetHTTPStatus() (httpStatus int) {

	switch m.Code {
	case OpenFileFailed:
		httpStatus = http.StatusInternalServerError
	case InvalidDataType:
		httpStatus = http.StatusExpectationFailed
	case InvalidValue:
		httpStatus = http.StatusBadRequest
	case InvalidID:
		httpStatus = http.StatusBadRequest
	case InvalidQueryParams:
		httpStatus = http.StatusBadRequest
	case InvalidUserPass:
		httpStatus = http.StatusBadRequest
	case DBConnectionFailed:
		httpStatus = http.StatusInternalServerError
	case DBRequestFailed:
		httpStatus = http.StatusInternalServerError
	case NotFound:
		httpStatus = http.StatusNotFound
	case NotUpdated:
		httpStatus = http.StatusInternalServerError
	case NotCreated:
		httpStatus = http.StatusInternalServerError
	case NotDeleted:
		httpStatus = http.StatusInternalServerError
	case HTTPConflict:
		httpStatus = http.StatusConflict
	default:
		httpStatus = int(m.Code)
	}

	return httpStatus
}
