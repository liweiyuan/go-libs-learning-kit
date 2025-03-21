package mock

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

// use gomonkey to mock add function(函数)
func TestMockAdd(t *testing.T) {
	// create a new patchs instance
	patches := gomonkey.NewPatches()
	// mock the add function to return 100
	patches.ApplyFunc(add, func(a, b int) int {
		return a * b
	})
	// defer the reset function to restore the original function
	defer patches.Reset()

	// test the add function
	result := add(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

// 替换方法
type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

// TestMockCalculatorAddMethod 测试 Calculator 的 Add 方法是否被正确替换。
func TestMockCalculatorAddMethod(t *testing.T) {
	cal := &Calculator{}
	patches := gomonkey.NewPatches()
	//替换掉Add方法
	patches.ApplyMethod(reflect.TypeOf(cal), "Add", func(_ *Calculator, a, b int) int {
		return a * b
	})
	defer patches.Reset()

	result := cal.Add(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

// 模拟接口
type DataStore interface {
	Save(data string) error
	Load() (string, error)
}

type MockDataStore struct{}

func (m *MockDataStore) Save(data string) error {
	return nil
}

func (m *MockDataStore) Load() (string, error) {
	return "", nil
}

func TestMockInterface(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	mock := &MockDataStore{}
	patches.ApplyMethod(reflect.TypeOf(mock), "Save", func(_ *MockDataStore, data string) error {
		return errors.New("mock save error")
	})

	patches.ApplyMethod(reflect.TypeOf(mock), "Load", func(_ *MockDataStore) (string, error) {
		return "mocked data", nil
	})

	// 测试Save方法
	err := mock.Save("test")
	assert.Error(t, err)
	assert.Equal(t, "mock save error", err.Error())

	// 测试Load方法
	data, err := mock.Load()
	assert.NoError(t, err)
	assert.Equal(t, "mocked data", data)
}

// 模拟全局变量
var GlobalConfig = struct {
	Debug   bool
	Version string
}{
	Debug:   false,
	Version: "1.0.0",
}

func TestMockGlobalVariable(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	// 模拟全局变量
	patches.ApplyGlobalVar(&GlobalConfig, struct {
		Debug   bool
		Version string
	}{
		Debug:   true,
		Version: "2.0.0",
	})

	assert.True(t, GlobalConfig.Debug)
	assert.Equal(t, "2.0.0", GlobalConfig.Version)
}

// 模拟时间
func TestMockTime(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	mockTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	patches.ApplyFunc(time.Now, func() time.Time {
		return mockTime
	})

	now := time.Now()
	assert.Equal(t, mockTime, now)
}

// 模拟HTTP请求
func TestMockHTTP(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(http.Get, func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("mocked response")),
		}, nil
	})

	resp, err := http.Get("http://example.com")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "mocked response", string(body))
}

// 模拟文件操作
func TestMockFileOperations(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(os.ReadFile, func(filename string) ([]byte, error) {
		return []byte("mocked file content"), nil
	})

	patches.ApplyFunc(os.WriteFile, func(filename string, data []byte, perm os.FileMode) error {
		return nil
	})

	// 测试读取文件
	content, err := os.ReadFile("test.txt")
	assert.NoError(t, err)
	assert.Equal(t, "mocked file content", string(content))

	// 测试写入文件
	err = os.WriteFile("test.txt", []byte("test"), 0644)
	assert.NoError(t, err)
}

// 模拟panic情况
func dangerousOperation() {
	panic("something went wrong")
}

func TestMockPanic(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	patches.ApplyFunc(dangerousOperation, func() {
		// 不执行panic
	})

	// 现在可以安全调用
	dangerousOperation()
}

// 模拟私有方法
type service struct{}

func (s *service) privateMethod() string {
	return "original"
}

func (s *service) PublicMethod() string {
	return s.privateMethod()
}

func TestMockPrivateMethod(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, "retrieve method by name failed", r)
		}
	}()
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	s := &service{}
	patches.ApplyMethod(reflect.TypeOf(s), "privateMethod", func(_ *service) string {
		return "mocked"
	})

	s.PublicMethod()
	//assert.Equal(t, "mocked", result)
}

// 链式模拟
type Client struct{}

func (c *Client) Connect() error {
	return nil
}

func (c *Client) Query() (string, error) {
	return "", nil
}

func TestChainedMocks(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	client := &Client{}

	// 首先模拟Connect方法
	patches.ApplyMethod(reflect.TypeOf(client), "Connect", func(_ *Client) error {
		return nil
	})

	// 然后模拟Query方法
	patches.ApplyMethod(reflect.TypeOf(client), "Query", func(_ *Client) (string, error) {
		return "mocked result", nil
	})

	// 测试链式调用
	err := client.Connect()
	assert.NoError(t, err)

	result, err := client.Query()
	assert.NoError(t, err)
	assert.Equal(t, "mocked result", result)
}

// 序列化模拟
func TestSequenceMocks(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	counter := 0
	patches.ApplyFunc(time.Now, func() time.Time {
		counter++
		switch counter {
		case 1:
			return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		case 2:
			return time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC)
		default:
			return time.Date(2023, 1, 1, 2, 0, 0, 0, time.UTC)
		}
	})

	// 第一次调用
	t1 := time.Now()
	assert.Equal(t, 2023, t1.Year())
	assert.Equal(t, 0, t1.Hour())

	// 第二次调用
	t2 := time.Now()
	assert.Equal(t, 2023, t2.Year())
	assert.Equal(t, 1, t2.Hour())

	// 第三次调用
	t3 := time.Now()
	assert.Equal(t, 2023, t3.Year())
	assert.Equal(t, 2, t3.Hour())
}

// 模拟错误处理链
type ErrorProcessor struct{}

func (ep *ErrorProcessor) Process(err error) error {
	if err != nil {
		return fmt.Errorf("processed error: %w", err)
	}
	return nil
}

func TestErrorChain(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	processor := &ErrorProcessor{}
	originalErr := errors.New("original error")

	patches.ApplyMethod(reflect.TypeOf(processor), "Process", func(_ *ErrorProcessor, err error) error {
		return fmt.Errorf("mocked error: %w", err)
	})

	err := processor.Process(originalErr)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mocked error")
	assert.Contains(t, err.Error(), "original error")
}

// 模拟复杂对象
type ComplexObject struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      map[string]interface{}
}

func (co *ComplexObject) Process() (string, error) {
	return "", nil
}

func TestMockComplexObject(t *testing.T) {
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	obj := &ComplexObject{
		ID:        1,
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Data: map[string]interface{}{
			"key": "value",
		},
	}

	patches.ApplyMethod(reflect.TypeOf(obj), "Process", func(_ *ComplexObject) (string, error) {
		return "processed", nil
	})

	result, err := obj.Process()
	assert.NoError(t, err)
	assert.Equal(t, "processed", result)
}
