// +build !unix

package cgo_python3

/*
// macOS:
#cgo darwin pkg-config: python3

// linux
#cgo linux pkg-config: python3

// windows
// should generate libpython3.a from python3.lib

#include <Python.h>
*/
import "C"
