package cgo_python3

// #include "go-python3.h"
import "C"

// PyObject* PyTuple_New(Py_ssize_t len)
// Return a new tuple object of size len, or NULL on failure.
func PyTuple_New(sz int) *PyObject {
	return togo(C.PyTuple_New(C.Py_ssize_t(sz)))
}
