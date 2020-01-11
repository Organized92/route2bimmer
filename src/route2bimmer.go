package main

import (
	"flag"
	"fmt"
	"strconv"
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

		// Output the waypoint coordinates
		for i := 0; i < len(gpxContents.Waypoints); i++ {
			fmt.Println("Waypoint name: " + gpxContents.Waypoints[i].Name)
			fmt.Println("Latitude: " + strconv.FormatFloat(gpxContents.Waypoints[i].Latitude, 'f', -1, 64))
			fmt.Println("Longitute: " + strconv.FormatFloat(gpxContents.Waypoints[i].Longitute, 'f', -1, 64))
		}

		// Calculate the distance and duration of the tracks, if provided
		for i := 0; i < len(gpxContents.Tracks); i++ {
			var totalDistance = calcTotalTrackDistance(gpxContents.Tracks[i]) / 1000
			var totalDurationMin = calcTotalTrackDuration(gpxContents.Tracks[i]) / 60
			fmt.Println("Distance of track \"" + gpxContents.Tracks[i].Name + "\": " + strconv.FormatFloat(totalDistance, 'f', 3, 64) + " km")
			fmt.Println("Duration of track \"" + gpxContents.Tracks[i].Name + "\": " + strconv.FormatInt(totalDurationMin, 10) + " min")
		}
	} else {
		fmt.Println(err)
	}
}
