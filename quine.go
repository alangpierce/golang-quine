package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "\u0060"
	code :=`package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "\u0060"
	code :=
	return code[:117] + tick + code + tick + code[117:]
}`
	return code[:117] + tick + code + tick + code[117:]
}