package main

import (
	"fmt"
	"os"

	cgo_python3 "github.com/chiachan163/cgo-python3"
)

func init() {
	err := cgo_python3.Initialize()
	if err != nil {
		panic(err.Error())
	}
	//defer cgo_python3.Finalize()
}

func main() {
	//fmt.Println(cgo_python3.ShowVersion())

	input := os.Args[1]
	fooModule := cgo_python3.PyImport_ImportModule("foo")
	if fooModule == nil {
		panic("Error importing module!")
	}

	//helloFunc := fooModule.GetAttrString("hello")
	helloFunc := fooModule.GetAttrString("hello_recall")
	if helloFunc == nil {
		panic("Error importing function!")
	}
	//rec := fooModule.CallMethod("hello_recall", cgo_python3.Py_BuildValue("s", "foo"))
	rec := helloFunc.CallFunctionObjArgs("0", cgo_python3.Py_BuildValue("(s)", input))
	if rec != nil {
		fmt.Println(cgo_python3.PyBytes_AsString(rec))
	}
	//helloFunc.CallFunction()
}
