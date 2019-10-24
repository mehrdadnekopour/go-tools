package templates

import (
	"fmt"
	"math"
	"net/http"
	// "./templates"
	// "sika.com/baran-service/config/env"
	// "../../config/env"
)

// ResponseTemplate standard template for http responses
type ResponseTemplate struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
	Links   interface{} `json:"links,omitempty"`
}

// Models for RESTFUL Data transmitions

// ErrorMessageTemplate ...
type ErrorMessageTemplate struct {
	Message string `json:"message"`
}

// ErrorResponse ... -> Customer Response Template witch Receives From ecommerce server
type ErrorResponse struct {
	*ResponseTemplate
	Data ErrorMessageTemplate `json:"data"`
}

// PaginateTemplate ...
type PaginateTemplate struct {
	Pages    int     `json:"pages"`
	Total    int     `json:"total"`
	Limit    int     `json:"limit"`
	Offset   uint    `json:"offset"`
	Next     *string `json:"next"`
	Previous *string `json:"prev"`
}

// NewPaginateTemplate ...
func NewPaginateTemplate(limit, offset, total int, repository string) PaginateTemplate {
	pages := math.Ceil(float64(total) / float64(limit))

	var next *string
	var prev *string

	var tempNext string
	var tempPrev string

	if offset <= total-limit {
		tempNext = fmt.Sprintf("%slimit=%d&offset=%d", repository, limit, offset+limit)
		next = &tempNext
	}

	var newPreviusOffset int

	if offset > 0 {
		if offset > limit {
			newPreviusOffset = offset - limit
		} else {
			newPreviusOffset = offset - 1
		}
		tempPrev = fmt.Sprintf("%s?limit=%d&offset=%d", repository, limit, newPreviusOffset)
		prev = &tempPrev
	}

	meta := PaginateTemplate{}
	meta.Next = next
	meta.Limit = limit
	meta.Offset = uint(offset)
	meta.Pages = int(pages)
	meta.Previous = prev
	meta.Total = total

	return meta
}

// GetWithCode return template with considering error code
func GetWithCode(code int, err error) (template *ResponseTemplate) {

	msg := GetMessage(code, err)

	switch code {
	case http.StatusInternalServerError:
		template = InternalServerError(msg)
	case http.StatusBadRequest:
		template = BadRequest(msg)
	case http.StatusForbidden:
		template = Forbidden(msg)
	case http.StatusNotFound:
		template = NotFound(msg)
	case http.StatusUnprocessableEntity:
		template = UnprocessableEntity(msg)
	case http.StatusUnauthorized:
		template = Unauthorized(msg)
	case http.StatusMethodNotAllowed:
		template = MethodNotAllowed(msg)
	case http.StatusGatewayTimeout:
		template = GatewayTimeOut(msg)
	default:
		template = InternalServerError(err.Error())
	}

	return
}

// GetMessage return error message
func GetMessage(code int, err error) (msg string) {
	// if env.GetBool("debug") {
	// 	msg = err.Error()

	// 	return
	// }

	switch code {
	case http.StatusInternalServerError:
		msg = err.Error()
	case http.StatusBadRequest:
		msg = "Bad request"
	case http.StatusForbidden:
		msg = "Forbidden"
	case http.StatusNotFound:
		msg = "Not found"
	case http.StatusUnprocessableEntity:
		msg = "Unprocessable entity"
	case http.StatusUnauthorized:
		msg = "Unauthorized"
	case http.StatusMethodNotAllowed:
		msg = "MethodNotAllowed"
	case http.StatusOK:
		msg = "Ok"
	default:
		msg = "Internal server error"
	}

	return
}

// BadRequest ...
func BadRequest(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusBadRequest,
		Status:  "BAD_REQUEST",
		Message: "Bad request",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// InternalServerError ...
func InternalServerError(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: "Internal server error",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// InternalServerErrorWithData ...
func InternalServerErrorWithData(data, msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusInternalServerError,
		Status:  "INTERNAL_SERVER_ERROR",
		Message: "Internal server error",
		Data: map[string]interface{}{
			"data":    data,
			"message": msg,
		},
	}
}

// NotFound ...
func NotFound(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusNotFound,
		Status:  "NOT_FOUND",
		Message: "Not found",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// UnprocessableEntity ...
func UnprocessableEntity(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusUnprocessableEntity,
		Status:  "UNPROCESSABLE_ENTITY",
		Message: "Unprocessable entity",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// Unauthorized ...
func Unauthorized(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusUnauthorized,
		Status:  "UNAUHTORIZED",
		Message: "Unauthorized",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// GatewayTimeOut ...
func GatewayTimeOut(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusGatewayTimeout,
		Status:  "GATEWAY_TIMEOUT",
		Message: "Gateway timeout",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// Forbidden ...
func Forbidden(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusForbidden,
		Status:  "FORBIDDEN",
		Message: "Forbidden",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// MethodNotAllowed ...
func MethodNotAllowed(msg interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusMethodNotAllowed,
		Status:  "METHOD_NOT_ALLOWED",
		Message: "Mehod not allowed",
		Data: map[string]interface{}{
			"message": msg,
		},
	}
}

// Ok ...
func Ok(data interface{}, meta interface{}) *ResponseTemplate {
	return &ResponseTemplate{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Ok",
		Data:    data,
		Meta:    meta,
	}
}
