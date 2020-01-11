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

	var gpxContents, err = readGPXFile(inputPtr)
	if err == nil {
		fmt.Println("GPX name: " + gpxContents.Metadata.Name)
		fmt.Println("GPX time: " + gpxContents.Metadata.Time)
		for i := 0; i < len(gpxContents.Waypoints); i++ {
			fmt.Println("Waypoint name: " + gpxContents.Waypoints[i].Name)
			fmt.Println("Latitude: " + gpxContents.Waypoints[i].Latitude)
			fmt.Println("Longitute: " + gpxContents.Waypoints[i].Longitute)
		}
	} else {
		fmt.Println(err)
	}
}
