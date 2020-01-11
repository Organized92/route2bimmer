package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
)

func main() {

	// command line arguments
	inputPtr := flag.String("input", "", "path to input file")
	outputPtr := flag.String("output", "route.zip", "path to output zip file")
	flag.Parse()

	fmt.Println("input:  ", *inputPtr)
	fmt.Println("output: ", *outputPtr)

	var gpxContents, err = readGPXFile(inputPtr)
	if err == nil {
		var route = mapGPXtoRoute(gpxContents, 1234567)

		var output, err2 = xml.MarshalIndent(route, "  ", "    ")
		if err == nil {
			os.Stdout.Write(output)
		} else {
			fmt.Println(err2)
		}
	}
}
