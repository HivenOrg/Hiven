package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/HivenOrg/Hiven/utils"
	"github.com/gofiber/fiber/v2"
)

func sendRequest(app *fiber.App, method string, path string, reqHeaders map[string]string, reqBody map[string]any) (*http.Response, error) {

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req := httptest.NewRequest(method, path, bytes.NewReader(jsonBytes))

	// It is safe to pass reqHeaders as nil
	for key, value := range reqHeaders {
		req.Header.Set(key, value)
	}

	// Use -1 to disable latency simulation
	res, err := app.Test(req, -1)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func parseResponse(res *http.Response) (map[string]any, error) {

	if res == nil {
		return nil, errors.New("response is nil")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func checkTokenFromResponse(res *http.Response, jwtSecret string) (userId uint, err error) {

	resBody, err := parseResponse(res)

	if err != nil {
		return 0, errors.New("unable to parse response body")
	}

	val, ok := resBody["bearer_token"]
	if !ok {
		return 0, errors.New("bearer_token not found")
	}

	token, ok := val.(string)
	if !ok {
		return 0, errors.New("bearer_token is not a string")
	}

	userId, err = utils.ValidateJWT(token, jwtSecret)

	if err != nil {
		return 0, errors.New("unable to validate JWT")
	}

	return userId, nil
}
