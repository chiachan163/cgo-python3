#include "Python.h"

#include "frameobject.h"
#include "marshal.h"
/* stdlib */
#include <stdlib.h>
#include <string.h>

/* go-python */
#define _gopy_max_varargs 8 /* maximum number of varargs accepted by go-python */

/* ---- none ----*/
PyObject* _gopy_pynone(void);

/* ---- object ----*/
PyObject* _gopy_pyfalse(void);
PyObject* _gopy_pytrue(void);
PyObject* _gopy_PyObject_CallFunction(PyObject *o, int len, char* types, void *args);
PyObject* _gopy_PyObject_CallMethod(PyObject *o, char *method, int len, char* pyfmt, void *cargs);
PyObject* _gopy_Py_BuildValue(char* format,int len, void *cargs);

/* ---- arg ---- */
