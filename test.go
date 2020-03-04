package cgo_python3

// #include "go-python3.h"
import "C"

func ShowVersion() string {
	return C.GoString(C.Py_GetVersion())
}
