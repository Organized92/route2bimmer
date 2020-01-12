package main

import (
	"encoding/xml"
	"strconv"
)

type DeliveryPackage struct {
	XMLName            xml.Name     `xml:"DeliveryPackage"`
	VersionNo          string       `xml:"VersionNo,attr"`
	CreationTime       string       `xml:"CreationTime,attr"`
	MapVersion         string       `xml:"MapVersion,attr"`
	LanguageCodeDesc   string       `xml:"Language_Code_Desc,attr"`
	CountryCodeDesc    string       `xml:"Country_Code_Desc,attr"`
	SupplierCodeDesc   string       `xml:"Supplier_Code_Desc,attr"`
	XY_Type            string       `xml:"XY_Type,attr"`
	Category_Code_Desc string       `xml:"Category_Code_Desc,attr"`
	CharSet            string       `xml:"Char_Set,attr"`
	UpdateType         string       `xml:"UpdateType,attr"`
	Coverage           string       `xml:"Coverage,attr"`
	Category           string       `xml:"Category,attr"`
	MajorVersion       string       `xml:"MajorVersion,attr"`
	MinorVersion       string       `xml:"MinorVersion,attr"`
	GuidedTour         []GuidedTour `xml:"GuidedTour"`
}

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

type Country struct {
	XMLName     xml.Name    `xml:"Country"`
	CountryCode int         `xml:"CountryCode"`
	Name        CountryName `xml:"Name"`
}

type CountryName struct {
	XMLName      xml.Name `xml:"Name"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Value        string   `xml:",chardata"`
}

type TourName struct {
	XMLName      xml.Name `xml:"Name"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

type TourLength struct {
	XMLName xml.Name `xml:"Length"`
	Unit    string   `xml:"Unit,attr"`
	Value   float64  `xml:",chardata"`
}

type TourDuration struct {
	XMLName xml.Name `xml:"Duration"`
	Unit    string   `xml:"Unit,attr"`
	Value   float64  `xml:",chardata"`
}

type TourIntroduction struct {
	XMLName      xml.Name `xml:"Introduction"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

type TourDescription struct {
	XMLName      xml.Name `xml:"Description"`
	LanguageCode string   `xml:"Language_Code,attr"`
	Text         string   `xml:"Text"`
}

type TourPicture struct {
	XMLName   xml.Name `xml:"Picture"`
	Reference string   `xml:"Reference"`
	Encoding  string   `xml:"Encoding"`
	Width     int      `xml:"Width"`
	Height    int      `xml:"Height"`
}

type EntryPoint struct {
	XMLName xml.Name `xml:"EntryPoint"`
	Route   string   `xml:"Route,attr"`
	Value   string   `xml:",chardata"`
}

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

type RouteWayPoint struct {
	XMLName      xml.Name           `xml:"WayPoint"`
	ID           string             `xml:"Id"`
	Locations    []WayPointLocation `xml:"Locations>Location"`
	Importance   string             `xml:"Importance"`
	Descriptions []TourDescription  `xml:"Descriptions>Description"`
}

type WayPointLocation struct {
	XMLName     xml.Name            `xml:"Location"`
	GeoPosition WayPointGeoPosition `xml:"GeoPosition"`
}

type WayPointGeoPosition struct {
	XMLName   xml.Name `xml:"GeoPosition"`
	Latitude  float64  `xml:"Latitude"`
	Longitude float64  `xml:"Longitude"`
}

func mapGPXtoRouteNav(gpx GPX, routeID int64) DeliveryPackage {

	var deliveryPackage DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = "0.0"
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = "0.0"
	deliveryPackage.LanguageCodeDesc = "../definitions/language.xml"
	deliveryPackage.CountryCodeDesc = "../definitions/country.xml"
	deliveryPackage.SupplierCodeDesc = "../definitions/supplier.xml"
	deliveryPackage.XY_Type = "WGS84"
	deliveryPackage.Category_Code_Desc = "../definitions/category.xml"
	deliveryPackage.CharSet = "UTF-8"
	deliveryPackage.UpdateType = "BulkUpdate"
	deliveryPackage.Coverage = "0"
	deliveryPackage.Category = "4096"
	deliveryPackage.MajorVersion = "0"
	deliveryPackage.MinorVersion = "0"

	// Guided Tour
	var guidedTour GuidedTour
	guidedTour.Access = "WEEKDAYS"
	guidedTour.Use = "ONFOOT"
	guidedTour.ID = int(routeID)
	guidedTour.TripType = 6

	// Country
	var country Country
	country.CountryCode = 3
	country.Name.LanguageCode = "ENG"
	country.Name.Value = "Germany"
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName TourName
	tourName.LanguageCode = "ENG"
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = "km"
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(calcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = "h"
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		totalDurationH = totalDurationH + float64(calcTotalTrackDuration(track))/60/60
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction TourIntroduction
	introduction.LanguageCode = "ENG"
	introduction.Text = "-"
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description TourDescription
	description.LanguageCode = "ENG"
	description.Text = "-"
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Entry Point
	var entryPoint EntryPoint
	entryPoint.Route = "1"
	entryPoint.Value = "0"
	guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

	// Route
	var route Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = 2
	route.Criteria = 0
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint RouteWayPoint
		waypoint.ID = strconv.FormatInt(int64(index), 10)
		if index == 0 {
			waypoint.Importance = "always"
		} else {
			waypoint.Importance = "optional"
		}

		// Location
		var location WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude
		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription TourDescription
		wpDescription.LanguageCode = "ENG"
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	return deliveryPackage
}


func mapGPXtoRouteNavigation(gpx GPX, routeID int64) DeliveryPackage {

	var deliveryPackage DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = "0.0"
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = "0.0"
	deliveryPackage.LanguageCodeDesc = "../definitions/language.xml"
	deliveryPackage.CountryCodeDesc = "../definitions/country.xml"
	deliveryPackage.SupplierCodeDesc = "../definitions/supplier.xml"
	deliveryPackage.XY_Type = "WGS84"
	deliveryPackage.Category_Code_Desc = "../definitions/category.xml"
	deliveryPackage.CharSet = "UTF-8"
	deliveryPackage.UpdateType = "BulkUpdate"
	deliveryPackage.Coverage = "0"
	deliveryPackage.Category = "4096"
	deliveryPackage.MajorVersion = "0"
	deliveryPackage.MinorVersion = "0"

	// Guided Tour
	var guidedTour GuidedTour
	guidedTour.Access = "WEEKDAYS"
	guidedTour.Use = "ONFOOT"
	guidedTour.ID = int(routeID)
	guidedTour.TripType = 6

	// Country
	var country Country
	country.CountryCode = 3
	country.Name.LanguageCode = "ENG"
	country.Name.Value = "Germany"
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName TourName
	tourName.LanguageCode = "ENG"
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = "km"
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(calcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = "h"
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		totalDurationH = totalDurationH + float64(calcTotalTrackDuration(track))/60/60
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction TourIntroduction
	introduction.LanguageCode = "ENG"
	introduction.Text = "-"
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description TourDescription
	description.LanguageCode = "ENG"
	description.Text = "-"
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Entry Point
	var entryPoint EntryPoint
	entryPoint.Route = "1"
	entryPoint.Value = "0_0"
	guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

	// Route
	var route Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = 2
	route.Criteria = 0
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint RouteWayPoint
		waypoint.ID = "0_" + strconv.FormatInt(int64(index), 10)
		if index == 0 {
			waypoint.Importance = "always"
		} else {
			waypoint.Importance = "optional"
		}

		// Location
		var location WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude
		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription TourDescription
		wpDescription.LanguageCode = "ENG"
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	return deliveryPackage
}
