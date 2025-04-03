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
