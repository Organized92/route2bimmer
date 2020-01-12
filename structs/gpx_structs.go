package structs

import "encoding/xml"

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
