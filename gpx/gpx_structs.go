package gpx

import "encoding/xml"

// GPX is the main structure for GPX files
type GPX struct {
	XMLName  xml.Name `xml:"gpx"`
	Metadata Metadata `xml:"metadata"`
	Routes   []Route  `xml:"rte"`
	Tracks   []Track  `xml:"trk"`
}

// Metadata contains some metadata from the GPX file
type Metadata struct {
	XMLName     xml.Name `xml:"metadata"`
	Name        string   `xml:"name"`
	Time        string   `xml:"time"`
	Description string   `xml:"desc"`
}

// Route contains details for a GPX route
type Route struct {
	XMLName        xml.Name        `xml:"rte"`
	Name           string          `xml:"name"`
	Description    string          `xml:"desc"`
	RouteWaypoints []RouteWaypoint `xml:"rtept"`
}

// RouteWaypoint contains details for a GPX route waypoint
type RouteWaypoint struct {
	XMLName     xml.Name `xml:"rtept"`
	Latitude    float64  `xml:"lat,attr"`
	Longitude   float64  `xml:"lon,attr"`
	Name        string   `xml:"name"`
	Description string   `xml:"desc"`
	Elevation   float64  `xml:"ele"`
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
	Elevation float64  `xml:"ele"`
	Time      string   `xml:"time"`
}
