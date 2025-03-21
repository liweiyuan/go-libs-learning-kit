package mock

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

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

// 测试各种断言方法
func TestVariousAssertions(t *testing.T) {
	Convey("Testing various assertion methods", t, func() {
		Convey("Numeric comparisons", func() {
			x := 5
			So(x, ShouldBeGreaterThan, 4)
			So(x, ShouldBeLessThan, 6)
			So(x, ShouldBeBetween, 4, 6)
			So(x, ShouldNotBeBetween, 6, 8)
		})

		Convey("String comparisons", func() {
			str := "hello world"
			So(str, ShouldContainSubstring, "world")
			So(str, ShouldNotContainSubstring, "goodbye")
			So(str, ShouldStartWith, "hello")
			So(str, ShouldEndWith, "world")
		})

		Convey("Slice operations", func() {
			nums := []int{1, 2, 3, 4, 5}
			So(nums, ShouldContain, 3)
			So(nums, ShouldNotContain, 6)
			So(nums, ShouldHaveLength, 5)
		})

		Convey("Map operations", func() {
			m := map[string]int{"a": 1, "b": 2}
			So(m, ShouldContainKey, "a")
			So(m, ShouldNotContainKey, "c")
			So(len(m), ShouldEqual, 2)
		})
	})
}

// 测试字符串操作
func TestStringOperations(t *testing.T) {
	Convey("Given a string manipulation scenario", t, func() {
		str := "Hello, World!"

		Convey("When converting to upper case", func() {
			result := strings.ToUpper(str)
			So(result, ShouldEqual, "HELLO, WORLD!")
		})

		Convey("When converting to lower case", func() {
			result := strings.ToLower(str)
			So(result, ShouldEqual, "hello, world!")
		})

		Convey("When splitting the string", func() {
			parts := strings.Split(str, ", ")
			So(parts, ShouldHaveLength, 2)
			So(parts[0], ShouldEqual, "Hello")
			So(parts[1], ShouldEqual, "World!")
		})
	})
}

// 测试切片操作
func TestSliceOperations(t *testing.T) {
	Convey("Given a slice of integers", t, func() {
		numbers := []int{1, 2, 3, 4, 5}

		Convey("When appending elements", func() {
			numbers = append(numbers, 6)
			So(numbers, ShouldHaveLength, 6)
			So(numbers[5], ShouldEqual, 6)
		})

		Convey("When slicing", func() {
			slice := numbers[1:4]
			So(slice, ShouldHaveLength, 3)
			So(slice, ShouldResemble, []int{2, 3, 4})
		})

		Reset(func() {
			numbers = []int{1, 2, 3, 4, 5}
		})
	})
}

// 测试错误处理
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestErrorHandling(t *testing.T) {
	Convey("Given a division function", t, func() {
		Convey("When dividing by non-zero", func() {
			result, err := divide(10, 2)

			Convey("It should return correct result", func() {
				So(err, ShouldBeNil)
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When dividing by zero", func() {
			result, err := divide(10, 0)

			Convey("It should return error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "division by zero")
				So(result, ShouldEqual, 0)
			})
		})
	})
}

// 测试并发操作
func TestConcurrentOperations(t *testing.T) {
	Convey("Given a concurrent scenario", t, func() {
		counter := 0
		var mu sync.Mutex
		var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				mu.Lock()
				counter++
				mu.Unlock()
			}()
		}

		Convey("When all goroutines complete", func() {
			wg.Wait()

			Convey("The counter should be correct", func() {
				So(counter, ShouldEqual, 10)
			})
		})
	})
}

// 测试嵌套结构体
type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func TestNestedStructs(t *testing.T) {
	Convey("Given a person with nested address", t, func() {
		person := Person{
			Name: "John",
			Age:  30,
			Address: Address{
				Street: "123 Main St",
				City:   "New York",
			},
		}

		Convey("When accessing nested fields", func() {
			So(person.Name, ShouldEqual, "John")
			So(person.Address.City, ShouldEqual, "New York")

			Convey("And modifying nested fields", func() {
				person.Address.City = "Boston"
				So(person.Address.City, ShouldEqual, "Boston")
			})
		})
	})
}

// 测试超时场景
func TestTimeout(t *testing.T) {
	Convey("Given a time-sensitive operation", t, func() {
		done := make(chan bool)

		go func() {
			time.Sleep(2 * time.Second)
			done <- true
		}()

		Convey("When waiting for completion", func() {
			select {
			case <-done:
				So(true, ShouldBeTrue) // 操作完成
			case <-time.After(1 * time.Second):
				So(false, ShouldBeFalse) // 超时
			}
		})
	})
}

// 使用Skip功能
func TestSkipFeature(t *testing.T) {
	Convey("Testing skip feature", t, func() {
		if testing.Short() {
			SkipConvey("Skipping long-running tests", func() {
				time.Sleep(1 * time.Second)
				So(true, ShouldBeTrue)
			})
		}

		Convey("This test should run", func() {
			So(1+1, ShouldEqual, 2)
		})
	})
}

// 使用自定义消息
func TestCustomMessages(t *testing.T) {
	Convey("Testing custom failure messages", t, func() {
		x := 5
		y := 10

		Convey("When comparing values", func() {
			So(x, ShouldBeLessThan, y)
			So(y, ShouldBeGreaterThan, x)
		})
	})
}

// 性能测试
func BenchmarkOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convey("Benchmark test", b, func() {
			result := 0
			for j := 0; j < 1000; j++ {
				result += j
			}
			So(result, ShouldBeGreaterThan, 0)
		})
	}
}
