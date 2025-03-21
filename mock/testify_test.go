package mock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestAssertAction(t *testing.T) {
	// assert equal
	assert.Equal(t, 1, 1, "they should be equal")
	// assert Null
	assert.Nil(t, nil)
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert true
	assert.True(t, true, "it should be true")

	// assert false
	assert.False(t, false, "it should be false")

	// assert contains
	assert.Contains(t, "hello world", "world", "\"hello world\" should contain \"world\"")

	// assert not contains
	assert.NotContains(t, "hello world", "planet", "\"hello world\" should not contain \"planet\"")

	// assert len
	assert.Len(t, []int{1, 2, 3}, 3, "slice should have length 3")

	// assert empty
	assert.Empty(t, []int{}, "slice should be empty")

	// assert not empty
	assert.NotEmpty(t, []int{1}, "slice should not be empty")

	// assert type of
	assert.IsType(t, 123, int(0), "value should be of type int")

	// assert no error
	assert.NoError(t, nil, "there should be no error")

	// assert error
	assert.Error(t, fmt.Errorf("an error occurred"), "there should be an error")

	// assert panic
	assert.Panics(t, func() { panic("a problem occurred") }, "the code should panic")

	// assert no panic
	assert.NotPanics(t, func() {}, "the code should not panic")
}

// 定义数据库接口
type Database interface {
	Connect() error
	Query(query string) ([]string, error)
	Close() error
}

// 创建模拟数据库
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockDatabase) Query(query string) ([]string, error) {
	args := m.Called(query)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockDatabase) Close() error {
	args := m.Called()
	return args.Error(0)
}

// 测试模拟对象
func TestMockDatabase(t *testing.T) {
	// 创建模拟对象
	db := new(MockDatabase)

	// 设置期望
	db.On("Connect").Return(nil)
	db.On("Query", "SELECT * FROM users").Return([]string{"user1", "user2"}, nil)
	db.On("Close").Return(nil)

	// 测试连接
	err := db.Connect()
	assert.NoError(t, err)

	// 测试查询
	results, err := db.Query("SELECT * FROM users")
	assert.NoError(t, err)
	assert.Equal(t, []string{"user1", "user2"}, results)

	// 测试关闭
	err = db.Close()
	assert.NoError(t, err)

	// 验证所有期望都被调用
	db.AssertExpectations(t)
}

// 定义测试套件
type ExampleTestSuite struct {
	suite.Suite
	db *MockDatabase
}

// 套件设置
func (suite *ExampleTestSuite) SetupTest() {
	suite.db = new(MockDatabase)
}

// 套件清理
func (suite *ExampleTestSuite) TearDownTest() {
	suite.db = nil
}

// 套件测试方法
func (suite *ExampleTestSuite) TestDatabaseOperations() {
	suite.db.On("Connect").Return(nil)
	err := suite.db.Connect()
	suite.NoError(err)
}

// 运行测试套件
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func TestStructComparison(t *testing.T) {
	person1 := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
		},
	}

	person2 := Person{
		Name: "John",
		Age:  30,
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
		},
	}

	assert.Equal(t, person1, person2)
	assert.EqualValues(t, person1, person2)
}

// 测试切片和映射的深度比较
func TestDeepComparison(t *testing.T) {
	// 切片测试
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	assert.Equal(t, slice1, slice2)

	// 映射测试
	map1 := map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]string{
			"city": "New York",
		},
	}

	map2 := map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]string{
			"city": "New York",
		},
	}

	assert.Equal(t, map1, map2)
}

// 测试并发操作
func TestConcurrentOp(t *testing.T) {
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	assert.Equal(t, 100, counter)
}

// 测试文件和IO操作
func TestFileOperations(t *testing.T) {
	// 创建临时文件
	content := []byte("test content")
	tmpfile, err := os.CreateTemp("", "example")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	// 写入文件
	n, err := tmpfile.Write(content)
	require.NoError(t, err)
	assert.Equal(t, len(content), n)

	// 读取文件
	readContent, err := os.ReadFile(tmpfile.Name())
	require.NoError(t, err)
	assert.Equal(t, content, readContent)
}

// 测试HTTP请求和响应
func TestHTTPOperations(t *testing.T) {
	// 创建测试服务器
	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/test", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// 发送请求
	resp, err := http.Get(server.URL + "/test")
	require.NoError(t, err)
	defer resp.Body.Close()

	// 检查响应
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.JSONEq(t, `{"status": "ok"}`, string(body))
}

// 测试JSON序列化和反序列化
func TestJSONOperations(t *testing.T) {
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// 测试序列化
	user := User{
		Name:  "John",
		Age:   30,
		Email: "john@example.com",
	}

	jsonData, err := json.Marshal(user)
	require.NoError(t, err)
	assert.JSONEq(t, `{"name":"John","age":30,"email":"john@example.com"}`, string(jsonData))

	// 测试反序列化
	var decoded User
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)
	assert.Equal(t, user, decoded)
}

// 测试时间操作
func TestTimeOperations(t *testing.T) {
	now := time.Now()
	future := now.Add(24 * time.Hour)

	assert.True(t, future.After(now))
	assert.True(t, now.Before(future))
	assert.Equal(t, 24*time.Hour, future.Sub(now))
}

// 测试自定义错误类型
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error code %d: %s", e.Code, e.Message)
}

func TestCustomError(t *testing.T) {
	err := &CustomError{
		Code:    404,
		Message: "not found",
	}

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "404")
	assert.Contains(t, err.Error(), "not found")
}

// 测试缓冲区操作
func TestBufferOperations(t *testing.T) {
	var buf bytes.Buffer

	// 写入缓冲区
	n, err := buf.WriteString("hello")
	require.NoError(t, err)
	assert.Equal(t, 5, n)

	// 读取缓冲区
	result := buf.String()
	assert.Equal(t, "hello", result)
}

// 使用require进行强制断言
func TestRequireAssertions(t *testing.T) {
	// 如果这个断言失败，测试将立即停止
	require.True(t, true, "this must be true")

	slice := []int{1, 2, 3}
	require.Len(t, slice, 3)
	require.NotEmpty(t, slice)

	// 在require断言后的代码只有在前面的断言都通过时才会执行
	assert.Equal(t, 1, slice[0])
}
