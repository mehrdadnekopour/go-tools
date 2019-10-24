package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mehrdadnekopour/go-tools/mypes"
	"github.com/mehrdadnekopour/go-tools/templates"
	// cartType "./cart"
)

// // Adapters ...
// type Adapters struct {
// 	Cart *cartType.CartAdapter
// }

// EndPoint ...
type EndPoint struct {
	Address string
	Port    int
	Token   string
}

// Instance ...
var Instance *EndPoint

// // RequestMethods ...
// type RequestMethods string

// const (
// 	// GET ...
// 	GET RequestMethods = "GET"
// 	// POST ...
// 	POST RequestMethods = "POST"
// 	// PUT ...
// 	PUT RequestMethods = "PUT"
// 	// DELETE ...
// 	DELETE RequestMethods = "DELETE"
// )

// // RequestMap ...
// var RequestMap = map[RequestMethods]func(ctx echo.Context, ep *EndPoint) mypes.Merror{
// 	GET:    GET,
// 	POST:   ep.POST,
// 	PUT:    ep.PUT,
// 	DELETE: ep.DELETE,
// }

// Init ...
func Init(ep *EndPoint) {
	Instance = ep
}

// Path ...
func (ep *EndPoint) Path() string {
	path := fmt.Sprintf("http://%s:%d", ep.Address, ep.Port)
	return path
}

// GET ...
func (ep *EndPoint) GET(url string, header http.Header, body interface{}, response interface{}) (myErr mypes.Merror) {

	urlParts := strings.Split(url, " ")
	urlToSend := strings.Join(urlParts, "%20")

	url = ep.Path() + urlToSend

	req, err := http.NewRequest("GET", url, nil)
	req.Header = header

	// authHeader := log.Sprintf("Bearer %s", ep.Token)
	// req.Header.Add("Authorization", authHeader)

	client := &http.Client{}

	result, err := client.Do(req)
	log.Println("________________________ * GET Request Sent To: * _______________________________")
	log.Println(url)
	log.Println("_________________________________________________________________________________")

	if err != nil {
		myErr.Set(true, err, http.StatusGatewayTimeout)
		return myErr
	}
	defer result.Body.Close()

	//------------------------------------------------------
	// result, err := http.Get(url)

	if err != nil {
		myErr.Set(true, err, http.StatusGatewayTimeout)
		return myErr
	}
	defer result.Body.Close()

	contents, err := ioutil.ReadAll(result.Body)
	log.Println("Response Body:", string(contents))

	if result.StatusCode == http.StatusOK {
		err = json.Unmarshal(contents, &response)
		if err != nil {
			myErr.Set(true, err, http.StatusUnprocessableEntity)
		}
		return myErr
	}

	var receivedError templates.ErrorResponse
	err = json.Unmarshal(contents, &receivedError)
	if err != nil {
		myErr.Set(true, err, http.StatusUnprocessableEntity)
		return myErr
	}

	msg := receivedError.Data.Message
	myErr.Set(true, errors.New(msg), mypes.ErrorCode(result.StatusCode))
	return
}

// POST ...
func (ep *EndPoint) POST(url string, header http.Header, body interface{}, response interface{}) (err mypes.Merror) {

	urlParts := strings.Split(url, " ")
	urlToSend := strings.Join(urlParts, "%20")

	url = ep.Path() + urlToSend

	jsonStr, e := json.Marshal(body)
	req, e := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// req, e := http.NewRequest("POST", url, body)

	if e != nil {
		err.Set(true, e, http.StatusBadRequest)
		return err
	}

	req.Header = header
	authHeader := fmt.Sprintf("Bearer %s", ep.Token)
	req.Header.Add("Authorization", authHeader)
	// for hTitle, hVal := range headers {
	// 	req.Header.Set(hTitle, hVal)
	// }

	client := &http.Client{}

	resp, e := client.Do(req)

	log.Println("________________________ * POST Request Sent To: * _______________________________")
	log.Println(url)
	log.Println("Request Body: ")
	log.Println(body)
	log.Println("_________________________________________________________________________________")

	if e != nil {
		err.Set(true, e, http.StatusGatewayTimeout)
		return err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response Body:", string(responseBody))
	if resp.StatusCode >= http.StatusOK && resp.StatusCode <= http.StatusIMUsed {
		resp.StatusCode = http.StatusOK
		e = json.Unmarshal(responseBody, &response)
		if e != nil {
			err.Set(true, e, http.StatusUnprocessableEntity)
		}
		return err
	}

	var receivedError templates.ErrorResponse
	e = json.Unmarshal(responseBody, &receivedError)
	if e != nil {
		err.Set(true, e, http.StatusUnprocessableEntity)
		return err
	}

	msg := receivedError.Data.Message
	err.Set(true, errors.New(msg), mypes.ErrorCode(resp.StatusCode))
	return err
}

// PUT ...
func (ep *EndPoint) PUT(url string, header http.Header, body interface{}, response interface{}) (err mypes.Merror) {

	urlParts := strings.Split(url, " ")
	urlToSend := strings.Join(urlParts, "%20")

	url = ep.Path() + urlToSend

	jsonStr, e := json.Marshal(body)

	req, e := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if e != nil {
		err.Set(true, e, http.StatusBadRequest)
		return err
	}

	req.Header = header
	authHeader := fmt.Sprintf("Bearer %s", ep.Token)
	req.Header.Add("Authorization", authHeader)

	// for hTitle, hVal := range headers {
	// 	req.Header.Set(hTitle, hVal)
	// }

	client := &http.Client{}

	resp, e := client.Do(req)

	log.Println("________________________ * PUT Request Sent To: * _______________________________")
	log.Println(url)
	log.Println("Request Body: ")
	log.Println(string(jsonStr))
	log.Println("_________________________________________________________________________________")

	if e != nil {
		err.Set(true, e, http.StatusGatewayTimeout)
		return err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("Response Body:", string(responseBody))
	if resp.StatusCode == http.StatusOK {
		e = json.Unmarshal(responseBody, &response)
		if e != nil {
			err.Set(true, e, http.StatusUnprocessableEntity)
		}
		return err
	}

	var receivedError templates.ErrorResponse
	e = json.Unmarshal(responseBody, &receivedError)
	if e != nil {
		err.Set(true, e, http.StatusUnprocessableEntity)
		return err
	}

	msg := receivedError.Data.Message
	err.Set(true, errors.New(msg), mypes.ErrorCode(resp.StatusCode))
	return err
}

// DELETE ...
func (ep *EndPoint) DELETE(url string, header http.Header, body interface{}, response interface{}) (err mypes.Merror) {

	urlParts := strings.Split(url, " ")
	urlToSend := strings.Join(urlParts, "%20")

	url = ep.Path() + urlToSend

	req, e := http.NewRequest("DELETE", url, nil)
	if e != nil {
		err.Set(true, e, http.StatusBadRequest)
		return err
	}

	req.Header = header
	authHeader := fmt.Sprintf("Bearer %s", ep.Token)
	req.Header.Add("Authorization", authHeader)

	client := &http.Client{}
	resp, e := client.Do(req)
	log.Println("DELETE Request Sent to: ", url)
	if e != nil {
		err.Set(true, e, http.StatusGatewayTimeout)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(respBody))
	if resp.StatusCode == http.StatusOK {
		e = json.Unmarshal(respBody, &response)
		if e != nil {
			err.Set(true, e, http.StatusUnprocessableEntity)
		}
		return
	}

	var receivedError templates.ErrorResponse
	e = json.Unmarshal(respBody, &receivedError)
	if e != nil {
		err.Set(true, e, http.StatusUnprocessableEntity)
		return err
	}

	msg := receivedError.Data.Message
	err.Set(true, errors.New(msg), mypes.ErrorCode(resp.StatusCode))
	return err
}
