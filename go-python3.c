#include "go-python3.h"


/* ---- test ----*/
static int cgo_PyArg_ParseTuple_ii(PyObject *arg, int *a, int *b) {
	return PyArg_ParseTuple(arg, "ii", a, b);
}

static PyObject* cgo_PyInit_gopkg(void) {
	static PyMethodDef methods[] = {
		{"sum", Py_gopkg_sum, METH_VARARGS, "Add two numbers."},
		{NULL, NULL, 0, NULL},
	};
	static struct PyModuleDef module = {
		PyModuleDef_HEAD_INIT, "gopkg", NULL, -1, methods,
	};
	return PyModule_Create(&module);
}