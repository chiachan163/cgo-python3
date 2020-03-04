#include "Python.h"

/* ---- test ----*/
extern PyObject* PyInit_gopkg();
extern PyObject* Py_gopkg_sum(PyObject *, PyObject *);

static int cgo_PyArg_ParseTuple_ii(PyObject *arg, int *a, int *b);
static PyObject* cgo_PyInit_gopkg(void);