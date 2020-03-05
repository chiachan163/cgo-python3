package cgo_python3

// #include "go-python3.h"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

// PyObject layer
type PyObject struct {
	ptr *C.PyObject
}

func (self *PyObject) topy() *C.PyObject {
	return self.ptr
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

// pyfmt returns the python format string for a given go value
func pyfmt(v interface{}) (unsafe.Pointer, string) {
	switch v := v.(type) {
	case bool:
		return unsafe.Pointer(&v), "b"

		// 	case byte:
		// 		return unsafe.Pointer(&v), "b"

	case int8:
		return unsafe.Pointer(&v), "b"

	case int16:
		return unsafe.Pointer(&v), "h"

	case int32:
		return unsafe.Pointer(&v), "i"

	case int64:
		return unsafe.Pointer(&v), "k"

	case int:
		switch unsafe.Sizeof(int(0)) {
		case 4:
			return unsafe.Pointer(&v), "i"
		case 8:
			return unsafe.Pointer(&v), "k"
		}

	case uint8:
		return unsafe.Pointer(&v), "B"

	case uint16:
		return unsafe.Pointer(&v), "H"

	case uint32:
		return unsafe.Pointer(&v), "I"

	case uint64:
		return unsafe.Pointer(&v), "K"

	case uint:
		switch unsafe.Sizeof(uint(0)) {
		case 4:
			return unsafe.Pointer(&v), "I"
		case 8:
			return unsafe.Pointer(&v), "K"
		}

	case float32:
		return unsafe.Pointer(&v), "f"

	case float64:
		return unsafe.Pointer(&v), "d"

	case complex64:
		return unsafe.Pointer(&v), "D"

	case complex128:
		return unsafe.Pointer(&v), "D"

	case string:
		cstr := C.CString(v)
		return unsafe.Pointer(cstr), "s"

	case *PyObject:
		return unsafe.Pointer(v.topy()), "O"

	}

	panic(fmt.Errorf("python: unknown type (%T)", v))
}

// PyObject* PyObject_GetAttrString(PyObject *o, const char *attr_name)¶
// Retrieve an attribute named attr_name from object o. Returns the attribute value on success, or
// NULL on failure. This is the equivalent of the Python expression o.attr_name.
func (self *PyObject) GetAttrString(attr_name string) *PyObject {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return togo(C.PyObject_GetAttrString(self.ptr, c_attr_name))
}

// PyObject* PyObject_CallFunction(PyObject *callable, const char *format, ...)
// Call a callable Python object callable, with a variable number of C arguments. The C arguments
// are described using a Py_BuildValue() style format string. The format can be NULL, indicating
// that no arguments are provided.
//
// Return the result of the call on success, or raise an exception and return NULL on failure.
//
// This is the equivalent of the Python expression: callable(*args).
//
// Note that if you only pass PyObject * args, PyObject_CallFunctionObjArgs() is a faster alternative.
func (self *PyObject) CallFunctionObjArgs(format string, args ...interface{}) *PyObject {
	return self.CallFunction(args...)
}

func (self *PyObject) CallFunction(args ...interface{}) *PyObject {
	if len(args) > int(C._gopy_max_varargs) {
		panic(fmt.Errorf(
			"gopy: maximum number of varargs (%d) exceeded (%d)",
			int(C._gopy_max_varargs),
			len(args),
		))
	}

	types := make([]string, 0, len(args))
	cargs := make([]unsafe.Pointer, 0, len(args))

	for _, arg := range args {
		ptr, typ := pyfmt(arg)
		types = append(types, typ)
		cargs = append(cargs, ptr)
		if typ == "s" {
			defer func(ptr unsafe.Pointer) {
				C.free(ptr)
			}(ptr)
		}
	}

	if len(args) <= 0 {
		o := C._gopy_PyObject_CallFunction(self.ptr, 0, nil, nil)
		return togo(o)
	}

	pyfmt := C.CString(strings.Join(types, ""))
	defer C.free(unsafe.Pointer(pyfmt))
	o := C._gopy_PyObject_CallFunction(
		self.ptr,
		C.int(len(args)),
		pyfmt,
		unsafe.Pointer(&cargs[0]),
	)
	return togo(o)

}

// PyObject* PyObject_Call(PyObject *callable, PyObject *args, PyObject *kwargs)
// Return value: New reference.
// Call a callable Python object callable, with arguments given by the tuple args, and named
// arguments given by the dictionary kwargs.
//
// args must not be NULL, use an empty tuple if no arguments are needed. If no named arguments
// are needed, kwargs can be NULL.
//
// Return the result of the call on success, or raise an exception and return NULL on failure.
//
// This is the equivalent of the Python expression: callable(*args, **kwargs).
func (self *PyObject) Call(args, kw *PyObject) *PyObject {
	return togo(C.PyObject_Call(self.ptr, args.ptr, kw.ptr))
}

// PyObject* PyObject_CallMethod(PyObject *o, char *method, char *format, ...)
// Call the method named name of object obj with a variable number of C arguments. The C arguments
// are described by a Py_BuildValue() format string that should produce a tuple.
//
// The format can be NULL, indicating that no arguments are provided.
//
// Return the result of the call on success, or raise an exception and return NULL on failure.
//
// This is the equivalent of the Python expression: obj.name(arg1, arg2, ...).
//
// Note that if you only pass PyObject * args, PyObject_CallMethodObjArgs() is a faster alternative.
func (self *PyObject) CallMethod(method string, args ...interface{}) *PyObject {
	if len(args) > int(C._gopy_max_varargs) {
		panic(fmt.Errorf(
			"gopy: maximum number of varargs (%d) exceeded (%d)",
			int(C._gopy_max_varargs),
			len(args),
		))
	}

	cmethod := C.CString(method)
	defer C.free(unsafe.Pointer(cmethod))

	types := make([]string, 0, len(args))
	cargs := make([]unsafe.Pointer, 0, len(args))

	for _, arg := range args {
		ptr, typ := pyfmt(arg)
		types = append(types, typ)
		cargs = append(cargs, ptr)
		if typ == "s" {
			defer func(ptr unsafe.Pointer) {
				C.free(ptr)
			}(ptr)
		}
	}

	if len(args) <= 0 {
		o := C._gopy_PyObject_CallMethod(self.ptr, cmethod, 0, nil, nil)
		return togo(o)
	}

	pyfmt := C.CString(strings.Join(types, ""))
	defer C.free(unsafe.Pointer(pyfmt))
	o := C._gopy_PyObject_CallMethod(
		self.ptr,
		cmethod,
		C.int(len(args)),
		pyfmt,
		unsafe.Pointer(&cargs[0]),
	)

	return togo(o)
}
