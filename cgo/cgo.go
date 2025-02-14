package cgo

/*
#include "calc.h"
*/
import "C"

func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
