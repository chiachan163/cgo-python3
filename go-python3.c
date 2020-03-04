#include "go-python3.h"

/* ---- none ----*/
PyObject* _gopy_pynone(void) {
	return Py_None;
}

/* ---- object ----*/
PyObject* _gopy_pyfalse(void) {
	return Py_False;
}

PyObject* _gopy_pytrue(void) {
	return Py_True;
}