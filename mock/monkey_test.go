package mock

import (
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
)

func add(a, b int) int {
	return a + b
}

// use gomonkey to mock add function(函数)
func TestMockAdd(t *testing.T) {
	// create a new patchs instance
	patchs := gomonkey.NewPatches()
	// mock the add function to return 100
	patchs.ApplyFunc(add, func(a, b int) int {
		return a * b
	})
	// defer the reset function to restore the original function
	defer patchs.Reset()

	// test the add function
	result := add(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

//替换方法

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func TestMockCalculatorAddMethod(t *testing.T) {

	cal := &Calculator{}
	patchs := gomonkey.NewPatches()
	//替换掉Add方法
	patchs.ApplyMethod(reflect.TypeOf(cal), "Add", func(_ *Calculator, a, b int) int {
		return a * b
	})
	defer patchs.Reset()

	result := cal.Add(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}
