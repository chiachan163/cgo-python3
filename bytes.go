package cgo_python3

// #include "go-python3.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// PyObject* PyBytes_FromString(const char *v)
// Return value: New reference.
// Return a new bytes object with a copy of the string v as value on success,
// and NULL on failure. The parameter v must not be NULL; it will not be checked.
func PyBytes_FromString(v string) *PyObject {
	cstr := C.CString(v)
	defer C.free(unsafe.Pointer(cstr))
	return togo(C.PyBytes_FromString(cstr))
}

// char* PyBytes_AsString(PyObject *o)
// Return a pointer to the contents of o. The pointer refers to the internal buffer
// of o, which consists of len(o) + 1 bytes. The last byte in the buffer is always
// null, regardless of whether there are any other null bytes. The data must not be
// modified in any way, unless the object was just created using
// PyBytes_FromStringAndSize(NULL, size). It must not be deallocated. If o is not a
// bytes object at all, PyBytes_AsString() returns NULL and raises TypeError.
func PyBytes_AsString(self *PyObject) string {

	c_str := C.PyBytes_AsString(self.ptr)
	//c_str := C.PyUnicode_FromObject(self.ptr)
	fmt.Println(C.PyUnicode_FromObject(self.ptr))
	// we dont own c_str...
	//defer C.free(unsafe.Pointer(c_str))
	return C.GoString(c_str)
}
