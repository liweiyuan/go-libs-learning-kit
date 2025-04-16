package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//测试相关的lambda或者闭包

// 闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestAdder(t *testing.T) {
	pos, neg := adder(), adder()
	assert.Equal(t, 10, pos(10))
	assert.Equal(t, -10, neg(-10))
}

// 闭包
func TestClosure(t *testing.T) {
	base := 100

	// 匿名函数（闭包），访问外部变量 base
	double := func(x int) int {
		return base + x*2
	}

	assert.Equal(t, 100, double(0))
	assert.Equal(t, 102, double(1))
	assert.Equal(t, 104, double(2))

	base = 200
	assert.Equal(t, 202, double(1))
}
