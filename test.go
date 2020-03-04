package cgo_python3

// #include "go-python3.h"
import "C"

func ShowVersion() string {
	return C.GoString(C.Py_GetVersion())
}

//export PyInit_gopkg
func PyInit_gopkg() *C.PyObject {
	return C.cgo_PyInit_gopkg()
}

//export Py_gopkg_sum
func Py_gopkg_sum(self, args *C.PyObject) *C.PyObject {
	var a, b C.int
	if C.cgo_PyArg_ParseTuple_ii(args, &a, &b) == 0 {
		return nil
	}
	return C.PyLong_FromLong(C.long(a + b))
}
