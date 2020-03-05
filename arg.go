package cgo_python3

// #include "go-python3.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// PyObject* Py_BuildValue(const char *format, ...)
// Create a new value based on a format string similar to those accepted by the PyArg_Parse*() family
// of functions and a sequence of values. Returns the value or NULL in the case of an error; an exception
// will be raised if NULL is returned.
//
// Py_BuildValue() does not always build a tuple. It builds a tuple only if its format string contains two
// or more format units. If the format string is empty, it returns None; if it contains exactly one format
// unit, it returns whatever object is described by that format unit. To force it to return a tuple of size
// 0 or one, parenthesize the format string.
//
// When memory buffers are passed as parameters to supply data to build objects, as for the s and s# formats,
// the required data is copied. Buffers provided by the caller are never referenced by the objects created by
// Py_BuildValue(). In other words, if your code invokes malloc() and passes the allocated memory to
// Py_BuildValue(), your code is responsible for calling free() for that memory once Py_BuildValue() returns.
//
// In the following description, the quoted form is the format unit; the entry in (round) parentheses is the
// Python object type that the format unit will return; and the entry in [square] brackets is the type of the
// C value(s) to be passed.
//
// The characters space, tab, colon and comma are ignored in format strings (but not within format units such
// as s#). This can be used to make long format strings a tad more readable.
//
// s (str or None) [const char *]
// Convert a null-terminated C string to a Python str object using 'utf-8' encoding. If the C string pointer
// is NULL, None is used.
//
// s# (str or None) [const char *, int or Py_ssize_t]
// Convert a C string and its length to a Python str object using 'utf-8' encoding. If the C string pointer is
// NULL, the length is ignored and None is returned.
//
// y (bytes) [const char *]
// This converts a C string to a Python bytes object. If the C string pointer is NULL, None is returned.
//
// y# (bytes) [const char *, int or Py_ssize_t]
// This converts a C string and its lengths to a Python object. If the C string pointer is NULL, None is returned.
//
// z (str or None) [const char *]
// Same as s.
//
// z# (str or None) [const char *, int or Py_ssize_t]
// Same as s#.
//
// u (str) [const wchar_t *]
// Convert a null-terminated wchar_t buffer of Unicode (UTF-16 or UCS-4) data to a Python Unicode object. If the Unicode buffer pointer is NULL, None is returned.
//
// u# (str) [const wchar_t *, int or Py_ssize_t]
// Convert a Unicode (UTF-16 or UCS-4) data buffer and its length to a Python Unicode object. If the Unicode buffer pointer is NULL, the length is ignored and None is returned.
//
// U (str or None) [const char *]
// Same as s.
//
// U# (str or None) [const char *, int or Py_ssize_t]
// Same as s#.
//
// i (int) [int]
// Convert a plain C int to a Python integer object.
//
// b (int) [char]
// Convert a plain C char to a Python integer object.
//
// h (int) [short int]
// Convert a plain C short int to a Python integer object.
//
// l (int) [long int]
// Convert a C long int to a Python integer object.
//
// B (int) [unsigned char]
// Convert a C unsigned char to a Python integer object.
//
// H (int) [unsigned short int]
// Convert a C unsigned short int to a Python integer object.
//
// I (int) [unsigned int]
// Convert a C unsigned int to a Python integer object.
//
// k (int) [unsigned long]
// Convert a C unsigned long to a Python integer object.
//
// L (int) [long long]
// Convert a C long long to a Python integer object.
//
// K (int) [unsigned long long]
// Convert a C unsigned long long to a Python integer object.
//
// n (int) [Py_ssize_t]
// Convert a C Py_ssize_t to a Python integer.
//
// c (bytes of length 1) [char]
// Convert a C int representing a byte to a Python bytes object of length 1.
//
// C (str of length 1) [int]
// Convert a C int representing a character to Python str object of length 1.
//
// d (float) [double]
// Convert a C double to a Python floating point number.
//
// f (float) [float]
// Convert a C float to a Python floating point number.
//
// D (complex) [Py_complex *]
// Convert a C Py_complex structure to a Python complex number.
//
// O (object) [PyObject *]
// Pass a Python object untouched (except for its reference count, which is incremented by one). If
// the object passed in is a NULL pointer, it is assumed that this was caused because the call producing
// the argument found an error and set an exception. Therefore, Py_BuildValue() will return NULL but
// won’t raise an exception. If no exception has been raised yet, SystemError is set.
//
// S (object) [PyObject *]
// Same as O.
//
// N (object) [PyObject *]
// Same as O, except it doesn’t increment the reference count on the object. Useful when the object is
// created by a call to an object constructor in the argument list.
//
// O& (object) [converter, anything]
// Convert anything to a Python object through a converter function. The function is called with anything
// (which should be compatible with void *) as its argument and should return a “new” Python object, or NULL
// if an error occurred.
//
// (items) (tuple) [matching-items]
// Convert a sequence of C values to a Python tuple with the same number of items.
//
// [items] (list) [matching-items]
// Convert a sequence of C values to a Python list with the same number of items.
//
// {items} (dict) [matching-items]
// Convert a sequence of C values to a Python dictionary. Each pair of consecutive C values adds one item
// to the dictionary, serving as key and value, respectively.
//
//If there is an error in the format string, the SystemError exception is set and NULL returned.
func Py_BuildValue(format string, args ...interface{}) *PyObject {
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
		o := C._gopy_Py_BuildValue(C.CString(format), 0, nil)
		return togo(o)
	}

	o := C._gopy_Py_BuildValue(C.CString(format), C.int(len(args)), unsafe.Pointer(&cargs[0]))
	return togo(o)
}

// int PyArg_Parse(PyObject *args, const char *format, ...)
// Function used to deconstruct the argument lists of “old-style” functions — these are functions
// which use the METH_OLDARGS parameter parsing method, which has been removed in Python 3. This
// is not recommended for use in parameter parsing in new code, and most code in the standard
// interpreter has been modified to no longer use this for that purpose. It does remain a convenient
// way to decompose other tuples, however, and may continue to be used for that purpose.
func PyArg_Parse(ptr *PyObject, format string, args ...interface{}) int {
	return 0
}
