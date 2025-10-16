package pokeapiinteractions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get[T any](url string) (*T, error) {
	data, err := MakeGetRequest(url)
	if err != nil {
		return nil, err
	}

	result, err := DecodeHttpResponse[T](data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MakeGetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP error: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

func DecodeHttpResponse[T any](data []byte) (*T, error) {
	var response T
	err := json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
