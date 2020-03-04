package cgo_python3

// #include "go-python3.h"
import "C"

// PyObject layer
type PyObject struct {
	ptr *C.PyObject
}

// The Python False object. This object has no methods.
// It needs to be treated just like any other object with respect to
// reference counts.
var Py_False = &PyObject{ptr: C._gopy_pyfalse()}

// PyObject* Py_True
// The Python True object. This object has no methods.
// It needs to be treated just like any other object with respect to
// reference counts.
var Py_True = &PyObject{ptr: C._gopy_pytrue()}

func togo(obj *C.PyObject) *PyObject {
	switch obj {
	case nil:
		return nil
	case Py_None.ptr:
		return Py_None
	case Py_True.ptr:
		return Py_True
	case Py_False.ptr:
		return Py_False
	default:
		return &PyObject{ptr: obj}
	}
}
