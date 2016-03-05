package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "\u0060"
	prefix := `package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "\u0060"
	prefix := `
	suffix := `
	return prefix + tick + prefix + tick + "\n	suffix := " + tick + suffix + tick + suffix
}`
	return prefix + tick + prefix + tick + "\n	suffix := " + tick + suffix + tick + suffix
}