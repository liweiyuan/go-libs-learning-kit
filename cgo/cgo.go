package cgo

/*
#include "calc.h"
*/
import "C"

// Add 函数接受两个整数 a 和 b，返回它们的和。
// 该函数通过 C 语言的 add 函数实现加法运算。
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
