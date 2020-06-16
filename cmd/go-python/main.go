// a go wrapper around py-main
package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {

	numpy := python.PyImport_ImportModule("numpy")
	fmt.Println(numpy)
	sys := python.PyImport_ImportModule("sys")
	fmt.Println(sys)

	//rc := python.Py_Main(os.Args)
	//os.Exit(rc)
}
