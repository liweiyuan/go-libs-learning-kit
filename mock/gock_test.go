package mock

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

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

// 测试不同的HTTP方法
func TestDifferentHTTPMethods(t *testing.T) {
	defer gock.Off()

	// 测试PUT请求
	t.Run("PUT Request", func(t *testing.T) {
		gock.New("http://example.com").
			Put("/resource").
			JSON(map[string]string{"name": "updated"}).
			Reply(200).
			JSON(map[string]string{"status": "updated"})

		client := &http.Client{}
		body := bytes.NewBuffer([]byte(`{"name": "updated"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://example.com/resource", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})

	// 测试DELETE请求
	t.Run("DELETE Request", func(t *testing.T) {
		gock.New("http://example.com").
			Delete("/resource/123").
			Reply(204)

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodDelete, "http://example.com/resource/123", nil)
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, 204, resp.StatusCode)
	})

	// 测试PATCH请求
	t.Run("PATCH Request", func(t *testing.T) {
		gock.New("http://example.com").
			Patch("/resource/123").
			JSON(map[string]string{"status": "active"}).
			Reply(200).
			JSON(map[string]string{"status": "active"})

		client := &http.Client{}
		body := bytes.NewBuffer([]byte(`{"status": "active"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://example.com/resource/123", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
	})
}

// 测试请求重试机制
func TestRetryMechanism(t *testing.T) {
	defer gock.Off()

	attempts := 0
	gock.New("http://example.com").
		Get("/flaky").
		Times(2).
		Reply(500).
		SetError(errors.New("server error"))

	gock.New("http://example.com").
		Get("/flaky").
		Reply(200).
		JSON(map[string]string{"status": "success"})

	client := &http.Client{}
	var resp *http.Response
	var err error

	// 模拟重试逻辑
	for i := 0; i < 3; i++ {
		resp, err = client.Get("http://example.com/flaky")
		attempts++
		if err == nil && resp.StatusCode == 200 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 3, attempts)
}

// 测试自定义响应头
func TestCustomHeaders(t *testing.T) {
	defer gock.Off()

	gock.New("http://example.com").
		Get("/").
		Reply(200).
		SetHeader("X-Custom-Header", "custom-value").
		SetHeader("X-Rate-Limit", "100").
		JSON(map[string]string{"status": "ok"})

	resp, err := http.Get("http://example.com/")
	assert.NoError(t, err)
	assert.Equal(t, "custom-value", resp.Header.Get("X-Custom-Header"))
	assert.Equal(t, "100", resp.Header.Get("X-Rate-Limit"))
}

// 测试二进制响应内容
func TestBinaryResponse(t *testing.T) {
	defer gock.Off()

	binaryData := []byte{0x00, 0x01, 0x02, 0x03}
	gock.New("http://example.com").
		Get("/binary").
		Reply(200).
		SetHeader("Content-Type", "application/octet-stream").
		Body(bytes.NewReader(binaryData))

	resp, err := http.Get("http://example.com/binary")
	assert.NoError(t, err)
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, binaryData, body)
}

// 测试网络错误
func TestNetworkErrors(t *testing.T) {
	defer gock.Off()

	// 模拟连接被重置
	gock.New("http://example.com").
		Get("/connection-reset").
		ReplyError(errors.New("connection reset by peer"))

	_, err := http.Get("http://example.com/connection-reset")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connection reset by peer")

	// 模拟DNS解析错误
	gock.New("http://nonexistent.example.com").
		Get("/").
		ReplyError(errors.New("no such host"))

	_, err = http.Get("http://nonexistent.example.com/")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no such host")
}

// 测试请求体验证
func TestRequestBodyValidation(t *testing.T) {
	defer gock.Off()

	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	expectedUser := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	gock.New("http://example.com").
		Post("/users").
		AddMatcher(func(req *http.Request, ereq *gock.Request) (bool, error) {
			var user User
			if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
				return false, err
			}
			return user == expectedUser, nil
		}).
		Reply(201).
		JSON(map[string]string{"status": "created"})

	client := &http.Client{}
	userData, _ := json.Marshal(expectedUser)
	req, _ := http.NewRequest(http.MethodPost, "http://example.com/users", bytes.NewBuffer(userData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)
}

// 测试复杂的路径匹配
func TestComplexPathMatching(t *testing.T) {
	defer gock.Off()

	// 测试路径参数匹配
	gock.New("http://example.com").
		Get("/users/([0-9]+)/posts/([0-9]+)").
		MatchParam("format", "json").
		Reply(200).
		JSON(map[string]string{"post": "found"})

	resp, err := http.Get("http://example.com/users/123/posts/456?format=json")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// 测试通配符匹配
	gock.New("http://example.com").
		Get("/api/.*").
		Reply(200).
		JSON(map[string]string{"status": "ok"})

	resp, err = http.Get("http://example.com/api/any/path/here")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

// 测试持久化连接
func TestPersistentConnection(t *testing.T) {
	defer gock.Off()

	// 启用持久连接
	gock.New("http://example.com").
		Persist().
		Get("/").
		Reply(200).
		JSON(map[string]string{"status": "ok"})

	client := &http.Client{}

	// 多次请求应该都成功
	for i := 0; i < 3; i++ {
		resp, err := client.Get("http://example.com/")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"status": "ok"}`, string(body))
		resp.Body.Close()
	}
}
