package cgo_python3

// #include "go-python3.h"
import "C"
import (
	"unsafe"
)

func PyImport_ImportModule(name string) *PyObject {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return togo(C.PyImport_ImportModule(c_name))
}
