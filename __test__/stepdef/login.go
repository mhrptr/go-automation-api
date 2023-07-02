package stepdef

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/basic-golang-api-automation/features/helper"
	"github.com/cucumber/godog"
	"github.com/yalp/jsonpath"
)

var urlEndpoint, contentBody string
var getEndpoint, postEndpoint, putEndpoint, deleteEndpoint *http.Request
var getResponse, postResponse, putResponse, deleteResponse *http.Response
var id interface{}

// GivenEndpoint : define endpoint
func GivenEndpoint(endpoint string) error {
	urlEndpoint = os.Getenv("BASE_URL") + endpoint

	return nil
}

// GetIDCRUD : get ID for PUT and DELETE
func GetIDCRUD() error {
	var jsonResponse interface{}

	jsonPath, _ := jsonpath.Prepare("$._id")
	responseBody, _ := ioutil.ReadAll(postResponse.Body)

	json.Unmarshal(responseBody, &jsonResponse)

	id, _ = jsonPath(jsonResponse)

	return nil
}

// GetEndpoint : hit endpoint get
func GetEndpoint() error {
	var err error

	client := &http.Client{}
	getEndpoint, err = http.NewRequest(http.MethodGet, urlEndpoint, nil)
	helper.LogPanicln(err)
	getResponse, err = client.Do(getEndpoint)
	helper.LogPanicln(err)

	return nil
}

// PostEndpoint : hit endpoint post
func PostEndpoint(contentBody) error {
	var err error

	body := []byte(contentBody)

	client := &http.Client{}
	postEndpoint, err := http.NewRequest(http.MethodPost, urlEndpoint, bytes.NewBuffer(body))
	helper.LogPanicln(err)

	postEndpoint.Header.Set("Content-Type", "application/json")

	postResponse, err = client.Do(postEndpoint)
	helper.LogPanicln(err)

	return nil
}

// PutEndpoint : hit endpoint put
func PutEndpoint(contentBody) error {
	var err error

	body := []byte(contentBody)

	client := &http.Client{}
	putEndpoint, err := http.NewRequest(http.MethodPut, urlEndpoint+"/"+fmt.Sprintf("%s", id), bytes.NewBuffer(body))
	helper.LogPanicln(err)

	putEndpoint.Header.Set("Content-Type", "application/json")

	putResponse, err = client.Do(putEndpoint)
	helper.LogPanicln(err)

	return nil
}

// DeleteEndpoint : hit endpoint delete
func DeleteEndpoint(contentBody) error {
	var err error

	body := []byte(contentBody)

	client := &http.Client{}
	deleteEndpoint, err := http.NewRequest(http.MethodDelete, urlEndpoint+"/"+fmt.Sprintf("%s", id), bytes.NewBuffer(body))
	helper.LogPanicln(err)
	deleteResponse, err = client.Do(deleteEndpoint)
	helper.LogPanicln(err)

	return nil
}

// ValidatePostResponse : validate endpoint post response
func ValidatePostResponse() error {
	if postResponse.StatusCode != 201 {
		log.Panicln("POST - Error code tidak sesuai")
	}

	return nil
}

func LoginSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`base url with endpoint "([^"]*)"$`, GetEndpoint)
	ctx.Step(`hit POST request with the following data: "([^"]*)"$`, PostEndpoint)
	ctx.Step(`validate POST response`, ValidatePostResponse)
}
