package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "`"
	prefix := `package main

import "fmt"

func main() {
	fmt.Print(getCode())
}

func getCode() string {
	tick := "%s"
	prefix := %s%s%s
	return fmt.Sprintf(prefix, tick, tick, prefix, tick)
}`
	return fmt.Sprintf(prefix, tick, tick, prefix, tick)
}