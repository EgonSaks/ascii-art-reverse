package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		reverse(args)
	} else if len(args) == 1 {
		output(args)
	} else if len(args) == 2 {
		output(args)
	} else {
		fmt.Println("Too many arguments")
	}
}
