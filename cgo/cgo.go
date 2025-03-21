package cgo

/*
#include "calc.h"
#include <stdlib.h>
*/
import "C"
import "errors"

// Add 函数接受两个整数 a 和 b，返回它们的和。
// 该函数通过 C 语言的 add 函数实现加法运算。
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

// Subtract 函数返回两个整数的差。
func Subtract(a, b int) int {
	return int(C.subtract(C.int(a), C.int(b)))
}

// Multiply 函数返回两个整数的乘积。
func Multiply(a, b int) int {
	return int(C.multiply(C.int(a), C.int(b)))
}

// Divide 函数返回两个整数的商。
// 如果除数为0，返回错误。
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return int(C.divide(C.int(a), C.int(b))), nil
}

// AddWithOverflowCheck 函数计算两个整数的和，并检查是否发生溢出。
// 如果发生溢出，返回错误。
func AddWithOverflowCheck(a, b int) (int, error) {
	var hasOverflow C.int
	result := C.add_with_overflow_check(C.int(a), C.int(b), &hasOverflow)
	if hasOverflow != 0 {
		return 0, errors.New("integer overflow")
	}
	return int(result), nil
}

// AddLong 函数计算两个64位整数的和。
func AddLong(a, b int64) int64 {
	return int64(C.add_long(C.longlong(a), C.longlong(b)))
}

// AbsValue 函数返回整数的绝对值。
func AbsValue(a int) int {
	return int(C.abs_value(C.int(a)))
}

// MaxValue 函数返回两个整数中的较大值。
func MaxValue(a, b int) int {
	return int(C.max_value(C.int(a), C.int(b)))
}

// MinValue 函数返回两个整数中的较小值。
func MinValue(a, b int) int {
	return int(C.min_value(C.int(a), C.int(b)))
}
