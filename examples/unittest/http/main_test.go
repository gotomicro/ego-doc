package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestURL = "/test"

func TestGetHelloEgo(t *testing.T) {
	s := NewHTTPMock()
	s.GET(TestURL, GetHelloEgo)
	response := s.MockGet(TestURL)
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello EGO", string(body))
}

func TestPostHelloEgoInvalid(t *testing.T) {
	s := NewHTTPMock()
	s.POST(TestURL, PostHelloEgo)
	response, err := s.MockPost(TestURL, []byte(`{"dddd"}`))
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "invalid params", string(body))
}

func TestPostHelloEgoOk(t *testing.T) {
	s := NewHTTPMock()
	s.POST(TestURL, PostHelloEgo)
	response, err := s.MockPost(TestURL, []byte(`{"name":"ego"}`))
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "Hello ego", string(body))
}
