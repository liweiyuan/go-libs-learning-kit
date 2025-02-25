package mock

import (
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
	Convey("Given two integers with starting values", t, func() {
		a := 1
		b := 2
		Convey("When the integers are added", func() {
			sum := a + b
			Convey("The result should be the sum of the two integers", func() {
				So(sum, ShouldEqual, 3)
			})
		})
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
