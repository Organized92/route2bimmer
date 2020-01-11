package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// GPX is the main structure for GPX files
type GPX struct {
	XMLName   xml.Name   `xml:"gpx"`
	Metadata  Metadata   `xml:"metadata"`
	Waypoints []Waypoint `xml:"wpt"`
	Tracks    []Track    `xml:"trk"`
}

// Metadata contains some metadata from the GPX file
type Metadata struct {
	XMLName xml.Name `xml:"metadata"`
	Name    string   `xml:"name"`
	Time    string   `xml:"time"`
}

// Waypoint contains details for a GPX waypoint
type Waypoint struct {
	XMLName     xml.Name `xml:"wpt"`
	Latitude    float64  `xml:"lat,attr"`
	Longitude   float64  `xml:"lon,attr"`
	Name        string   `xml:"name"`
	Description string   `xml:"desc"`
}

// Track contains details for a GPX track
type Track struct {
	XMLName  xml.Name       `xml:"trk"`
	Name     string         `xml:"name"`
	Segments []TrackSegment `xml:"trkseg"`
}

// TrackSegment contain details for a GPX track segment
type TrackSegment struct {
	XMLName xml.Name     `xml:"trkseg"`
	Points  []TrackPoint `xml:"trkpt"`
}

// TrackPoint contains details for a GPX trackpoint
type TrackPoint struct {
	XMLName   xml.Name `xml:"trkpt"`
	Latitude  float64  `xml:"lat,attr"`
	Longitude float64  `xml:"lon,attr"`
	Elevation float64  `xml:"elevation"`
	Time      string   `xml:"time"`
}

// readGPXFile reads the contents of the supplied filepath and returns a structure of type GPX in case of success
func readGPXFile(inputPath *string) (GPX, error) {

	// Declare return value
	var gpxContents GPX

	// Open the GPX file
	gpxFile, err := os.Open(*inputPath)
	if err != nil {
		return gpxContents, err
	}

	// Defer the closing of the GPX file so we can parse it
	defer gpxFile.Close()

	// Read the GPX file as a byte array
	byteValue, _ := ioutil.ReadAll(gpxFile)

	// Unmarshal the byteArray which contains the GPX file into the gpxContents
	xml.Unmarshal(byteValue, &gpxContents)

	// Return the contents of the GPX file
	return gpxContents, nil
}
