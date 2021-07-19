package main

import (
	"fmt"
)
func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered program", v)
		}
	}()
	panic("test panic")
	fmt.Println("hellow")
}
