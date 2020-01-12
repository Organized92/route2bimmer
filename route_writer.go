package main

import (
	"route2bimmer/structs"
	"strconv"
)

// MapGPXtoRouteNav converts the GPX contents to a BMW compatible XML file for folder "Nav"
func MapGPXtoRouteNav(gpx structs.GPX, routeID int64) structs.DeliveryPackage {

	var deliveryPackage structs.DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = "0.0"
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = "0.0"
	deliveryPackage.LanguageCodeDesc = "../definitions/language.xml"
	deliveryPackage.CountryCodeDesc = "../definitions/country.xml"
	deliveryPackage.SupplierCodeDesc = "../definitions/supplier.xml"
	deliveryPackage.XYType = "WGS84"
	deliveryPackage.CategoryCodeDesc = "../definitions/category.xml"
	deliveryPackage.CharSet = "UTF-8"
	deliveryPackage.UpdateType = "BulkUpdate"
	deliveryPackage.Coverage = "0"
	deliveryPackage.Category = "4096"
	deliveryPackage.MajorVersion = "0"
	deliveryPackage.MinorVersion = "0"

	// Guided Tour
	var guidedTour structs.GuidedTour
	guidedTour.Access = "WEEKDAYS"
	guidedTour.Use = "ONFOOT"
	guidedTour.ID = int(routeID)
	guidedTour.TripType = 6

	// Country
	var country structs.Country
	country.CountryCode = 3
	country.Name.LanguageCode = "ENG"
	country.Name.Value = "Germany"
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName structs.TourName
	tourName.LanguageCode = "ENG"
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = "km"
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(CalcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = "h"
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		totalDurationH = totalDurationH + float64(CalcTotalTrackDuration(track))/60/60
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction structs.TourIntroduction
	introduction.LanguageCode = "ENG"
	introduction.Text = "-"
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description structs.TourDescription
	description.LanguageCode = "ENG"
	description.Text = "-"
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture structs.TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Entry Point
	var entryPoint structs.EntryPoint
	entryPoint.Route = "1"
	entryPoint.Value = "0"
	guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

	// Route
	var route structs.Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = 2
	route.Criteria = 0
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint structs.RouteWayPoint
		waypoint.ID = strconv.FormatInt(int64(index), 10)
		if index == 0 || index == len(gpx.Waypoints)-1 {
			waypoint.Importance = "always"
		} else {
			waypoint.Importance = "optional"
		}

		// Location
		var location structs.WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude
		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription structs.TourDescription
		wpDescription.LanguageCode = "ENG"
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	return deliveryPackage
}

// MapGPXtoRouteNavigation converts the GPX contents to a BMW compatible XML file for folder "Navigation"
func MapGPXtoRouteNavigation(gpx structs.GPX, routeID int64) structs.DeliveryPackage {

	var deliveryPackage structs.DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = "0.0"
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = "0.0"
	deliveryPackage.LanguageCodeDesc = "../definitions/language.xml"
	deliveryPackage.CountryCodeDesc = "../definitions/country.xml"
	deliveryPackage.SupplierCodeDesc = "../definitions/supplier.xml"
	deliveryPackage.XYType = "WGS84"
	deliveryPackage.CategoryCodeDesc = "../definitions/category.xml"
	deliveryPackage.CharSet = "UTF-8"
	deliveryPackage.UpdateType = "BulkUpdate"
	deliveryPackage.Coverage = "0"
	deliveryPackage.Category = "4096"
	deliveryPackage.MajorVersion = "0"
	deliveryPackage.MinorVersion = "0"

	// Guided Tour
	var guidedTour structs.GuidedTour
	guidedTour.Access = "WEEKDAYS"
	guidedTour.Use = "ONFOOT"
	guidedTour.ID = int(routeID)
	guidedTour.TripType = 6

	// Country
	var country structs.Country
	country.CountryCode = 3
	country.Name.LanguageCode = "ENG"
	country.Name.Value = "Germany"
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName structs.TourName
	tourName.LanguageCode = "ENG"
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = "km"
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(CalcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = "h"
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		totalDurationH = totalDurationH + float64(CalcTotalTrackDuration(track))/60/60
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction structs.TourIntroduction
	introduction.LanguageCode = "ENG"
	introduction.Text = "-"
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description structs.TourDescription
	description.LanguageCode = "ENG"
	description.Text = "-"
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture structs.TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Entry Point
	var entryPoint structs.EntryPoint
	entryPoint.Route = "1"
	entryPoint.Value = "0_0"
	guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

	// Route
	var route structs.Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = 2
	route.Criteria = 0
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint structs.RouteWayPoint
		waypoint.ID = "0_" + strconv.FormatInt(int64(index), 10)
		if index == 0 || index == len(gpx.Waypoints)-1 {
			waypoint.Importance = "always"
		} else {
			waypoint.Importance = "optional"
		}

		// Location
		var location structs.WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude
		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription structs.TourDescription
		wpDescription.LanguageCode = "ENG"
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	return deliveryPackage
}
