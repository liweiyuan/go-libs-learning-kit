package json

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
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
