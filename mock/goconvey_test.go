package mock

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		Convey("When the integer is incremented", func() {
			x++
			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}

func TestPositiveNumberAdd(t *testing.T) {
	testCases := []struct {
		num1 int
		num2 int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{4, 5, 9},
		{5, 6, 11},
		{0, 0, 0},
		{1, -1, 0},
	}

	Convey("Given two integers with starting values", t, func() {
		for _, tc := range testCases {
			Convey(fmt.Sprintf("When %d and %d are added", tc.num1, tc.num2), func() {
				sum := tc.num1 + tc.num2
				Convey(fmt.Sprintf("The result should be %d", tc.want), func() {
					So(sum, ShouldEqual, tc.want)
				})
			})
		}
	})

}

func TestNegativeNumberAdd(t *testing.T) {
	Convey("Given a negative integer and a positive integer", t, func() {
		a := -1
		b := 2
		Convey("When the integers are added", func() {
			sum := a + b
			Convey("The result should be the sum of the two integers", func() {
				So(sum, ShouldEqual, 1)
			})
		})
	})
}
