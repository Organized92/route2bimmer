package main

import (
	"math"
	"time"
)

type meterCoordinates struct {
	x float64
	y float64
	z float64
}

// calcTotalTrackDuration calculates the amount of time you will need for this route, based on the time data provided in the track data of the GPX file
// this will return the duration in seconds
func calcTotalTrackDuration(track Track) int64 {

	var totalDuration int64

	// We have to loop over the track segments
	for _, segment := range track.Segments {

		// We need the first and the last track point, the ones in between are irrelevant
		var timeFirstPoint, err1 = time.Parse(time.RFC3339, segment.Points[0].Time)
		var timeLastPoint, err2 = time.Parse(time.RFC3339, segment.Points[len(segment.Points)-1].Time)

		if err1 == nil && err2 == nil {

			// Calculate the duration
			var durationNs = timeLastPoint.Sub(timeFirstPoint)

			// Add to total duration
			totalDuration = totalDuration + int64(math.Round(durationNs.Seconds()))

		}

	}

	return totalDuration

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
