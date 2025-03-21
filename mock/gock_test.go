package mock

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"io"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestSimpleGoMock(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		Get("/").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	resp, err := http.Get("http://example.com/")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"foo": "bar"}`, string(body))
}

func TestMatchHeaders(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		MatchHeader("Authorization", "Bearer token").
		Get("/").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	req, err := http.NewRequest("GET", "http://example.com/", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer token")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"foo": "bar"}`, string(body))
}

func TestQueryParams(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		Get("/").
		MatchParam("foo", "bar").
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	resp, err := http.Get("http://example.com/?foo=bar")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"foo": "bar"}`, string(body))
}

func TestPostRequest(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		Post("/").
		MatchType("json").
		JSON(map[string]string{"name": "test"}).
		Reply(201).
		JSON(map[string]string{"id": "123"})

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://example.com/", io.NopCloser(strings.NewReader(`{"name": "test"}`)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"id": "123"}`, string(body))
}

func TestErrorResponse(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		Get("/").
		Reply(404).
		JSON(map[string]string{"error": "not found"})

	resp, err := http.Get("http://example.com/")
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"error": "not found"}`, string(body))
}

func TestTimeout(t *testing.T) {
	defer gock.Off() // 清除所有的mock

	gock.New("http://example.com").
		Get("/").
		Reply(200).
		Delay(2 * time.Second).
		JSON(map[string]string{"foo": "bar"})

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	_, err := client.Get("http://example.com/")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Client.Timeout exceeded")
}
