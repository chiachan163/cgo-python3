package cgo_python3

// #include "go-python3.h"
import "C"

// PyObject* PyDict_New()
// Return a new empty dictionary, or NULL on failure.
// 返回一个新的空字典，失败时返回 NULL。
func PyDict_New() *PyObject {
	return togo(C.PyDict_New())
}
