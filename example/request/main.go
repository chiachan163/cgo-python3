package main

import (
	"fmt"
	cgo_python "github.com/chiachan163/cgo-python2"
	//cgo_python "github.com/sbinet/go-python"
	"os"
)

func init() {
	err := cgo_python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	input := os.Args[1]
	fooModule := cgo_python.PyImport_ImportModule("foo")
	if fooModule == nil {
		panic("Error importing module!")
	}

	helloFunc := fooModule.GetAttrString("hello_recall")
	if helloFunc == nil {
		panic("Error importing function!")
	}

	rec := helloFunc.CallFunctionObjArgs("0", cgo_python.PyString_FromString(input))
	if rec != nil {
		fmt.Println(cgo_python.PyString_AsString(rec))
	}
}
