// Package funk provides a set of functions for working with slices.
package funk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/thoas/go-funk"
)

func TestFunkContains(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	container := funk.Contains(numbers, 3)
	if !container {
		t.Errorf("Expected container to be true")
	}

	Convey("Test funk.Contains", t, func() {
		So(container, ShouldBeTrue)
	})
	//复杂的测试用例
	Convey("Test funk.Contains", t, func() {
		n := 9
		Convey("Test funk.Contains is false", func() {
			So(funk.Contains(numbers, n), ShouldBeFalse)
		})
	})
}

// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkMap(t *testing.T) {
	Convey("Tets funk.Map", t, func() {
		numbers := []int{1, 2, 3, 4}
		Convey("Calculate the square of each element", func() {
			So(funk.Map(numbers, func(x int) int { return x * x }).([]int), ShouldResemble, []int{1, 4, 9, 16})
		})
	})
}

// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkFilter(t *testing.T) {
	Convey("Test funk.Filter", t, func() {
		numbers := []int{1, 2, 3, 4}
		Convey("Filter even numbers", func() {
			So(funk.Filter(numbers, func(x int) bool { return x%2 == 0 }).([]int), ShouldResemble, []int{2, 4})
		})
	})
}

// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkReduce(t *testing.T) {
	Convey("Test funk.Reduce", t, func() {
		numbers := []int{1, 2, 3, 4}
		Convey("Sum all elements", func() {
			So(funk.Reduce(numbers, func(acc, x int) int { return acc + x }, 0), ShouldEqual, 10)
		})
	})
}

// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkForEach(t *testing.T) {
	Convey("Test funk.ForEach", t, func() {
		numbers := []int{1, 2, 3, 4}
		Convey("Print each element", func() {
			var result []int
			funk.ForEach(numbers, func(x int) { result = append(result, x) })
			So(result, ShouldResemble, []int{1, 2, 3, 4})
		})
	})
}

// funk.Uniq
// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkUniq(t *testing.T) {
	Convey("Test funk.Uniq", t, func() {
		numbers := []int{1, 2, 2, 3, 4, 4}
		Convey("Remove duplicates", func() {
			So(funk.Uniq(numbers).([]int), ShouldResemble, []int{1, 2, 3, 4})
		})
	})
}

// funk.Keys
// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkKeys(t *testing.T) {
	Convey("Test funk.Keys", t, func() {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		Convey("Get keys", func() {
			So(funk.Keys(m).([]string), ShouldContain, "a")
			So(funk.Keys(m).([]string), ShouldContain, "b")
			So(funk.Keys(m).([]string), ShouldContain, "c")
		})
	})
}

// funk.Values
// 对切片或者数组中的每一个元素应用一个新的函数，返回新的切片或者数据
func TestFunkValues(t *testing.T) {
	Convey("Test funk.Values", t, func() {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		Convey("Get values", func() {
			So(funk.Values(m).([]int), ShouldContain, 1)
			So(funk.Values(m).([]int), ShouldContain, 2)
			So(funk.Values(m).([]int), ShouldContain, 3)
		})
	})
}
