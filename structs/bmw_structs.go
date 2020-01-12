package structs

import "encoding/xml"

// DeliveryPackage is the root node of the BMW route format
type DeliveryPackage struct {
	XMLName          xml.Name     `xml:"DeliveryPackage"`
	VersionNo        string       `xml:"VersionNo,attr"`
	CreationTime     string       `xml:"CreationTime,attr"`
	MapVersion       string       `xml:"MapVersion,attr"`
	LanguageCodeDesc string       `xml:"Language_Code_Desc,attr"`
	CountryCodeDesc  string       `xml:"Country_Code_Desc,attr"`
	SupplierCodeDesc string       `xml:"Supplier_Code_Desc,attr"`
	XYType           string       `xml:"XY_Type,attr"`
	CategoryCodeDesc string       `xml:"Category_Code_Desc,attr"`
	CharSet          string       `xml:"Char_Set,attr"`
	UpdateType       string       `xml:"UpdateType,attr"`
	Coverage         string       `xml:"Coverage,attr"`
	Category         string       `xml:"Category,attr"`
	MajorVersion     string       `xml:"MajorVersion,attr"`
	MinorVersion     string       `xml:"MinorVersion,attr"`
	GuidedTour       []GuidedTour `xml:"GuidedTour"`
}

// GuidedTour contains all data for the route
type GuidedTour struct {
	XMLName       xml.Name           `xml:"GuidedTour"`
	Access        string             `xml:"access,attr"`
	Use           string             `xml:"use,attr"`
	ID            int                `xml:"Id"`
	TripType      int                `xml:"TripType"`
	Countries     []Country          `xml:"Countries>Country"`
	Names         []TourName         `xml:"Names>Name"`
	Length        TourLength         `xml:"Length"`
	Duration      TourDuration       `xml:"Duration"`
	Introductions []TourIntroduction `xml:"Introductions>Introduction"`
	Descriptions  []TourDescription  `xml:"Descriptions>Description"`
	Pictures      []TourPicture      `xml:"Pictures>Picture"`
	EntryPoints   []EntryPoint       `xml:"EntryPoints>EntryPoint"`
	Routes        []Route            `xml:"Routes>Route"`
}

// Country contains details about some countries
type Country struct {
	XMLName     xml.Name    `xml:"Country"`
	CountryCode int         `xml:"CountryCode"`
	Name        CountryName `xml:"Name"`
}

// CountryName contains a name and a language code of a country
type CountryName struct {
	XMLName      xml.Name `xml:"Name"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Value        string   `xml:",chardata"`
}

// TourName contains a name and a language code of the tour
type TourName struct {
	XMLName      xml.Name `xml:"Name"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

// TourLength contains the length and a unit of this tour
type TourLength struct {
	XMLName xml.Name `xml:"Length"`
	Unit    string   `xml:"Unit,attr"`
	Value   float64  `xml:",chardata"`
}

// TourDuration contains the driving duration and a unit of this tour
type TourDuration struct {
	XMLName xml.Name `xml:"Duration"`
	Unit    string   `xml:"Unit,attr"`
	Value   float64  `xml:",chardata"`
}

// TourIntroduction contains information about introductions?
type TourIntroduction struct {
	XMLName      xml.Name `xml:"Introduction"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

// TourDescription contains descriptions (text & language code) for a tour
type TourDescription struct {
	XMLName      xml.Name `xml:"Description"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

// TourPicture contains
type TourPicture struct {
	XMLName   xml.Name `xml:"Picture"`
	Reference string   `xml:"Reference"`
	Encoding  string   `xml:"Encoding"`
	Width     int      `xml:"Width"`
	Height    int      `xml:"Height"`
}

// EntryPoint contains information about where the tour can be started
type EntryPoint struct {
	XMLName xml.Name `xml:"EntryPoint"`
	Route   string   `xml:"Route,attr"`
	Value   string   `xml:",chardata"`
}

// Route contains information about the route
type Route struct {
	XMLName      xml.Name        `xml:"Route"`
	RouteID      int             `xml:"RouteID"`
	WayPoint     []RouteWayPoint `xml:"WayPoint"`
	Length       TourLength      `xml:"Length"`
	Duration     TourDuration    `xml:"Duration"`
	CostModel    int             `xml:"CostModel"`
	Criteria     int             `xml:"Criteria"`
	AgoraCString string          `xml:"AgoraCString"`
}

// RouteWayPoint is a way point of a route
type RouteWayPoint struct {
	XMLName      xml.Name           `xml:"WayPoint"`
	ID           string             `xml:"Id"`
	Locations    []WayPointLocation `xml:"Locations>Location"`
	Importance   string             `xml:"Importance"`
	Descriptions []TourDescription  `xml:"Descriptions>Description"`
}

// WayPointLocation contains information about where the waypoint is
type WayPointLocation struct {
	XMLName     xml.Name            `xml:"Location"`
	Address     *WayPointAddress    `xml:"Address,omitempty"`
	GeoPosition WayPointGeoPosition `xml:"GeoPosition"`
}

// WayPointAddress contains the address details for a waypoint
type WayPointAddress struct {
	XMLName       xml.Name      `xml:"Address"`
	ParsedAddress ParsedAddress `xml:"ParsedAddress"`
}

// ParsedAddress contains parsed address data
type ParsedAddress struct {
	XMLName             xml.Name            `xml:"ParsedAddress"`
	ParsedStreetAddress ParsedStreetAddress `xml:"ParsedStreetAddress"`
	ParsedPlace         ParsedPlace         `xml:"ParsedPlace"`
}

// ParsedStreetAddress contains parsed street address data
type ParsedStreetAddress struct {
	XMLName          xml.Name         `xml:"ParsedStreetAddress"`
	ParsedStreetName ParsedStreetName `xml:"ParsedStreetName"`
}

// ParsedStreetName contains a parsed street name
type ParsedStreetName struct {
	XMLName    xml.Name `xml:"ParsedStreetName"`
	StreetName string   `xml:"StreetName"`
}

// ParsedPlace contains a parsed place information
type ParsedPlace struct {
	XMLName     xml.Name `xml:"ParsedPlace"`
	PlaceLevel4 string   `xml:"PlaceLevel4"`
}

// WayPointGeoPosition contains information about the geographical position
type WayPointGeoPosition struct {
	XMLName   xml.Name `xml:"GeoPosition"`
	Latitude  float64  `xml:"Latitude"`
	Longitude float64  `xml:"Longitude"`
}
