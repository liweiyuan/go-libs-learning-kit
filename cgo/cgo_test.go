package cgo

import (
	"math"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 基本算术运算测试
func TestBasicOperations(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		result := Add(3, 4)
		assert.Equal(t, 7, result, "3 + 4 should equal 7")
	})

	t.Run("Subtraction", func(t *testing.T) {
		result := Subtract(7, 3)
		assert.Equal(t, 4, result, "7 - 3 should equal 4")
	})

	t.Run("Multiplication", func(t *testing.T) {
		result := Multiply(3, 4)
		assert.Equal(t, 12, result, "3 * 4 should equal 12")
	})

	t.Run("Division", func(t *testing.T) {
		result, err := Divide(12, 3)
		assert.NoError(t, err)
		assert.Equal(t, 4, result, "12 / 3 should equal 4")
	})
}

// 边界值测试
func TestEdgeCases(t *testing.T) {
	t.Run("Division by Zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "division by zero")
	})

	t.Run("Integer Overflow", func(t *testing.T) {
		_, err := AddWithOverflowCheck(math.MaxInt32, 1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "integer overflow")
	})

	t.Run("Integer Underflow", func(t *testing.T) {
		_, err := AddWithOverflowCheck(math.MinInt32, -1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "integer overflow")
	})

	t.Run("Large Numbers Addition", func(t *testing.T) {
		result := AddLong(int64(math.MaxInt32)+1, int64(math.MaxInt32)+1)
		assert.Greater(t, result, int64(math.MaxInt32))
	})
}

// 特殊函数测试
func TestSpecialFunctions(t *testing.T) {
	t.Run("Absolute Value", func(t *testing.T) {
		assert.Equal(t, 5, AbsValue(5), "abs(5) should be 5")
		assert.Equal(t, 5, AbsValue(-5), "abs(-5) should be 5")
		assert.Equal(t, 0, AbsValue(0), "abs(0) should be 0")
	})

	t.Run("Maximum Value", func(t *testing.T) {
		assert.Equal(t, 5, MaxValue(5, 3), "max(5,3) should be 5")
		assert.Equal(t, 5, MaxValue(3, 5), "max(3,5) should be 5")
		assert.Equal(t, 5, MaxValue(5, 5), "max(5,5) should be 5")
	})

	t.Run("Minimum Value", func(t *testing.T) {
		assert.Equal(t, 3, MinValue(5, 3), "min(5,3) should be 3")
		assert.Equal(t, 3, MinValue(3, 5), "min(3,5) should be 3")
		assert.Equal(t, 5, MinValue(5, 5), "min(5,5) should be 5")
	})
}

// 并发测试
func TestConcurrency(t *testing.T) {
	t.Run("Concurrent Addition", func(t *testing.T) {
		var wg sync.WaitGroup
		results := make([]int, 100)
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = Add(index, index)
			}(i)
		}
		wg.Wait()

		for i := 0; i < 100; i++ {
			assert.Equal(t, i*2, results[i], "Concurrent addition failed")
		}
	})

	t.Run("Concurrent Mixed Operations", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				Add(val, val)
				Subtract(val*2, val)
				Multiply(val, 2)
				if val != 0 {
					Divide(val*2, 2)
				}
			}(i)
		}
		wg.Wait()
	})
}

// 性能测试
func BenchmarkOperations(b *testing.B) {
	b.Run("CGo Addition", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(i, i)
		}
	})

	b.Run("Go Addition", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i + i
		}
	})

	b.Run("CGo Multiplication", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Multiply(i, i)
		}
	})

	b.Run("Go Multiplication", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = i * i
		}
	})
}

// 表格驱动测试
func TestTableDriven(t *testing.T) {
	addTests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive Numbers", 3, 4, 7},
		{"Zero and Positive", 0, 4, 4},
		{"Zero and Zero", 0, 0, 0},
		{"Negative Numbers", -3, -4, -7},
		{"Mixed Numbers", -3, 4, 1},
	}

	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}

	divideTests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"Normal Division", 10, 2, 5, false},
		{"Division by Zero", 10, 0, 0, true},
		{"Negative Division", -10, 2, -5, false},
		{"Zero Division", 0, 5, 0, false},
	}

	for _, tt := range divideTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// 复合操作测试
func TestCompositeOperations(t *testing.T) {
	t.Run("Complex Calculation", func(t *testing.T) {
		// (10 + 5) * 2 - 8
		step1 := Add(10, 5)         // 15
		step2 := Multiply(step1, 2) // 30
		step3 := Subtract(step2, 8) // 22
		assert.Equal(t, 22, step3)
	})

	t.Run("Absolute Operations", func(t *testing.T) {
		// |a + b| where a and b are negative
		a, b := -5, -3
		sum := Add(a, b)
		result := AbsValue(sum)
		assert.Equal(t, 8, result)
	})
}
