package main

import "fmt"

type Example struct {
	flag bool
}

func main() {
	var ex Example
	ex = Example{
		flag: true,
	}
	fmt.Println(ex.flag)
}
