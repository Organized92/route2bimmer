package main

import (
	"flag"
	"fmt"
)

func main() {

	// command line arguments
	inputPtr := flag.String("input", "", "path to input file")
	outputPtr := flag.String("output", "route.zip", "path to output zip file")
	flag.Parse()

	fmt.Println("input:  ", *inputPtr)
	fmt.Println("output: ", *outputPtr)
}
