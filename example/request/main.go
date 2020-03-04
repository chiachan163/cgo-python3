package main

import (
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

	//input := os.Args[1]
	fooModule := cgo_python3.PyImport_ImportModule("foo")
	if fooModule == nil {
		panic("Error importing module!")
	}

	//helloFunc := fooModule.GetAttrString("hello_recall")
	//if helloFunc == nil {
	//	panic("Error importing function!")
	//}
	//
	//rec := helloFunc.CallFunctionObjArgs("0", cgo_python3.PyString_FromString(input))
	//if rec != nil {
	//	fmt.Println(cgo_python3.PyString_AsString(rec))
	//}
}
