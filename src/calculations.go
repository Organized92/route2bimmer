package main

import (
	"math"
)

type meterCoordinates struct {
	x float64
	y float64
	z float64
}

// calcTotalTrackDistance calculates the distance of a GPX track, also considering the elevation provided in the GPX file.
func calcTotalTrackDistance(track Track) float64 {

	var totalDistance float64
	var distance float64

	// We have to loop over the track segments
	for _, segment := range track.Segments {

		// Loop over the track points to calculate the distance
		// We skip the first entry so that we will not run out of bounds at the end of the array
		for i := 1; i < len(segment.Points); i++ {

			// calculate the meter based coordinates
			var coord = coordinatesToMeters(segment.Points[i].Elevation, segment.Points[i].Latitude, segment.Points[i].Longitude)
			var coordBefore = coordinatesToMeters(segment.Points[i-1].Elevation, segment.Points[i-1].Latitude, segment.Points[i-1].Longitude)

			// calculate the distance between these two points and add it to total distance of the track
			distance = math.Sqrt(math.Pow(coord.x-coordBefore.x, 2) + math.Pow(coord.y-coordBefore.y, 2) + math.Pow(coord.z-coordBefore.z, 2))
			totalDistance = totalDistance + distance
		}
	}

	return totalDistance
}

// coordinatesToMeters converts the coordinates to meter-like coordinates, so that we can calculate the distance easier
func coordinatesToMeters(elevation float64, latitude float64, longitude float64) meterCoordinates {
	// constant for earths medium radius in meters, source: wikipedia
	const earthRadiusInMeters float64 = 6371000.785

	// return variable
	var coordinates meterCoordinates

	// Thanks to Ignacio Vazquez-Abrams for his answer of this stackoverflow question:
	// https://stackoverflow.com/questions/29827636/distance-between-two-points-including-elevation
	var r = earthRadiusInMeters + elevation
	var theta = latitude * math.Pi / 180
	var phi = longitude * math.Pi / 180

	// calculate the meter based coordinates and return them
	coordinates.x = r * math.Cos(theta) * math.Cos(phi)
	coordinates.y = r * math.Cos(theta) * math.Sin(phi)
	coordinates.z = r * math.Sin(theta)

	return coordinates
}
