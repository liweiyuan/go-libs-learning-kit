package json

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 基础Person结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

// 带有嵌套结构的Person
type PersonWithDetails struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Address Address   `json:"address"`
	Contact *Contact  `json:"contact,omitempty"`
	Tags    []string  `json:"tags,omitempty"`
	Created time.Time `json:"created"`
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
	ZIP     string `json:"zip,omitempty"`
}

type Contact struct {
	Email  string `json:"email"`
	Phone  string `json:"phone,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

func TestJsonSerializationAndDeserialization(t *testing.T) {
	p := Person{
		Name: "小明",
		Age:  18,
		City: "北京",
	}
	//序列化测试
	t.Run("JSON Serialization", func(t *testing.T) {
		jsonBytes, err := json.Marshal(p)
		assert.NoError(t, err, "Serialization failed")
		assert.Equal(t, `{"name":"小明","age":18,"city":"北京"}`, string(jsonBytes), "Serialization failed")
	})

	//反序列化测试
	t.Run("JSON Deserialization", func(t *testing.T) {
		jsonStr := `{"name":"小明","age":18,"city":"北京"}`
		var p2 Person
		err := json.Unmarshal([]byte(jsonStr), &p2)
		assert.NoError(t, err, "Deserialization failed")
		assert.Equal(t, p, p2, "Deserialization failed")
	})
}

// 测试嵌套结构体
func TestNestedStructs(t *testing.T) {
	now := time.Now()
	person := PersonWithDetails{
		Name: "张三",
		Age:  25,
		Address: Address{
			City:    "上海",
			Country: "中国",
			ZIP:     "200000",
		},
		Contact: &Contact{
			Email: "zhangsan@example.com",
			Phone: "021-12345678",
		},
		Tags:    []string{"学生", "程序员"},
		Created: now,
	}

	t.Run("Nested Struct Serialization", func(t *testing.T) {
		jsonBytes, err := json.Marshal(person)
		assert.NoError(t, err)

		var decoded PersonWithDetails
		err = json.Unmarshal(jsonBytes, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, person.Name, decoded.Name)
		assert.Equal(t, person.Address.City, decoded.Address.City)
		assert.Equal(t, person.Contact.Email, decoded.Contact.Email)
		assert.Equal(t, person.Tags, decoded.Tags)
	})
}

// 测试omitempty标签
func TestOmitEmpty(t *testing.T) {
	contact := Contact{
		Email: "test@example.com",
		// Phone和Mobile字段为空，应该被省略
	}

	jsonBytes, err := json.Marshal(contact)
	assert.NoError(t, err)
	jsonStr := string(jsonBytes)
	assert.Contains(t, jsonStr, "email")
	assert.NotContains(t, jsonStr, "phone")
	assert.NotContains(t, jsonStr, "mobile")
}

// 测试空值和null
func TestNullValues(t *testing.T) {
	jsonStr := `{"name":null,"age":0,"city":""}`
	var p Person
	err := json.Unmarshal([]byte(jsonStr), &p)
	assert.NoError(t, err)
	assert.Empty(t, p.Name)
	assert.Zero(t, p.Age)
	assert.Empty(t, p.City)
}

// 测试JSON数组
func TestJSONArray(t *testing.T) {
	people := []Person{
		{Name: "张三", Age: 20, City: "北京"},
		{Name: "李四", Age: 25, City: "上海"},
	}

	t.Run("Array Serialization", func(t *testing.T) {
		jsonBytes, err := json.Marshal(people)
		assert.NoError(t, err)

		var decoded []Person
		err = json.Unmarshal(jsonBytes, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, len(people), len(decoded))
		assert.Equal(t, people[0].Name, decoded[0].Name)
	})
}

// 测试错误处理
func TestErrorHandling(t *testing.T) {
	t.Run("Invalid JSON", func(t *testing.T) {
		invalidJSON := `{"name":"张三",age:20}` // 缺少引号的无效JSON
		var p Person
		err := json.Unmarshal([]byte(invalidJSON), &p)
		assert.Error(t, err)
	})

	t.Run("Type Mismatch", func(t *testing.T) {
		invalidType := `{"name":"张三","age":"不是数字","city":"北京"}`
		var p Person
		err := json.Unmarshal([]byte(invalidType), &p)
		assert.Error(t, err)
	})
}

// 测试接口类型
type Animal interface {
	GetName() string
}

type Dog struct {
	Name    string `json:"name"`
	Breed   string `json:"breed"`
	Type    string `json:"type"`
	BarkStr string `json:"bark"`
}

func (d Dog) GetName() string {
	return d.Name
}

type Cat struct {
	Name    string `json:"name"`
	Color   string `json:"color"`
	Type    string `json:"type"`
	MeowStr string `json:"meow"`
}

func (c Cat) GetName() string {
	return c.Name
}

func TestInterfaceJSON(t *testing.T) {
	dog := Dog{Name: "旺财", Breed: "金毛", Type: "dog", BarkStr: "汪汪"}
	cat := Cat{Name: "咪咪", Color: "白色", Type: "cat", MeowStr: "喵喵"}

	t.Run("Interface Serialization", func(t *testing.T) {
		dogJSON, err := json.Marshal(dog)
		assert.NoError(t, err)
		assert.Contains(t, string(dogJSON), "旺财")

		catJSON, err := json.Marshal(cat)
		assert.NoError(t, err)
		assert.Contains(t, string(catJSON), "咪咪")
	})
}

// 性能测试
func BenchmarkJSONMarshal(b *testing.B) {
	person := PersonWithDetails{
		Name: "张三",
		Age:  25,
		Address: Address{
			City:    "上海",
			Country: "中国",
			ZIP:     "200000",
		},
		Contact: &Contact{
			Email: "zhangsan@example.com",
			Phone: "021-12345678",
		},
		Tags:    []string{"学生", "程序员"},
		Created: time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(person)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	jsonStr := `{"name":"张三","age":25,"address":{"city":"上海","country":"中国","zip":"200000"},"contact":{"email":"zhangsan@example.com","phone":"021-12345678"},"tags":["学生","程序员"]}`
	var person PersonWithDetails

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal([]byte(jsonStr), &person)
	}
}
