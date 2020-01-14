package main

import (
	"route2bimmer/structs"
	"strconv"
)

const conImportanceAlways string = "always"
const conImportanceOptional string = "optional"
const conLanguageCodeEnglish string = "ENG"
const conVersionZeroDotZero string = "0.0"
const conLanguageCodeDesc string = "../definitions/language.xml"
const conCountryCodeDesc string = "../definitions/country.xml"
const conSupplierCodeDesc string = "../definitions/supplier.xml"
const conCategoryCodeDesc string = "../definitions/category.xml"
const conXYType string = "WGS84"
const conCharSet string = "UTF-8"
const conUpdateType string = "BulkUpdate"
const conCoverage string = "0"
const conCategory string = "4096"
const conMajorVersion string = "0"
const conMinorVersion string = "0"
const conTourAccess string = "WEEKDAYS"
const conTourUse string = "ONFOOT"
const conTourTripType int = 6
const conCountryCode int = 3
const conCountryName string = "Germany"
const conUnitDistance string = "km"
const conUnitDuration string = "h"
const conTextDefault string = "-"
const conRouteCostModel int = 2
const conRouteCriteria int = 0

// MapGPXtoRouteNav converts the GPX contents to a BMW compatible XML file for folder "Nav"
func MapGPXtoRouteNav(gpx structs.GPX, routeID int64) (structs.DeliveryPackage, error) {

	var deliveryPackage structs.DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = conVersionZeroDotZero
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = conVersionZeroDotZero
	deliveryPackage.LanguageCodeDesc = conLanguageCodeDesc
	deliveryPackage.CountryCodeDesc = conCountryCodeDesc
	deliveryPackage.SupplierCodeDesc = conSupplierCodeDesc
	deliveryPackage.XYType = conXYType
	deliveryPackage.CategoryCodeDesc = conCategoryCodeDesc
	deliveryPackage.CharSet = conCharSet
	deliveryPackage.UpdateType = conUpdateType
	deliveryPackage.Coverage = conCoverage
	deliveryPackage.Category = conCategory
	deliveryPackage.MajorVersion = conMajorVersion
	deliveryPackage.MinorVersion = conMinorVersion

	// Guided Tour
	var guidedTour structs.GuidedTour
	guidedTour.Access = conTourAccess
	guidedTour.Use = conTourUse
	guidedTour.ID = int(routeID)
	guidedTour.TripType = conTourTripType

	// Country
	var country structs.Country
	country.CountryCode = conCountryCode
	country.Name.LanguageCode = conLanguageCodeEnglish
	country.Name.Value = conCountryName
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName structs.TourName
	tourName.LanguageCode = conLanguageCodeEnglish
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = conUnitDistance
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(CalcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = conUnitDuration
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		partialDurationH, err := CalcTotalTrackDuration(track)
		if err != nil {
			return deliveryPackage, err
		}
		totalDurationH = totalDurationH + float64(partialDurationH)/3600
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction structs.TourIntroduction
	introduction.LanguageCode = conLanguageCodeEnglish
	introduction.Text = conTextDefault
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description structs.TourDescription
	description.LanguageCode = conLanguageCodeEnglish
	description.Text = conTextDefault
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture structs.TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Route
	var route structs.Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = conRouteCostModel
	route.Criteria = conRouteCriteria
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint structs.RouteWayPoint
		waypoint.ID = strconv.FormatInt(int64(index), 10)
		if index == 0 || index == len(gpx.Waypoints)-1 {

			waypoint.Importance = conImportanceAlways

			// Places with importance = always have to be EntryPoints
			var entryPoint structs.EntryPoint
			entryPoint.Route = "1"
			entryPoint.Value = waypoint.ID
			guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

		} else {
			waypoint.Importance = conImportanceOptional
		}

		// Location
		var location structs.WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude

		// If this is a importance = always waypoint, we have to add address data
		if waypoint.Importance == conImportanceAlways {
			var addr structs.WayPointAddress
			addr.ParsedAddress.ParsedStreetAddress.ParsedStreetName.StreetName = gpxWaypoint.Name
			addr.ParsedAddress.ParsedPlace.PlaceLevel4 = gpxWaypoint.Name
			location.Address = &addr
		}

		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription structs.TourDescription
		wpDescription.LanguageCode = conLanguageCodeEnglish
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	var err error
	return deliveryPackage, err
}

// MapGPXtoRouteNavigation converts the GPX contents to a BMW compatible XML file for folder "Navigation"
func MapGPXtoRouteNavigation(gpx structs.GPX, routeID int64) (structs.DeliveryPackage, error) {

	var deliveryPackage structs.DeliveryPackage

	// Basic data
	deliveryPackage.VersionNo = conVersionZeroDotZero
	deliveryPackage.CreationTime = gpx.Metadata.Time
	deliveryPackage.MapVersion = conVersionZeroDotZero
	deliveryPackage.LanguageCodeDesc = conLanguageCodeDesc
	deliveryPackage.CountryCodeDesc = conCountryCodeDesc
	deliveryPackage.SupplierCodeDesc = conSupplierCodeDesc
	deliveryPackage.XYType = conXYType
	deliveryPackage.CategoryCodeDesc = conCategoryCodeDesc
	deliveryPackage.CharSet = conCharSet
	deliveryPackage.UpdateType = conUpdateType
	deliveryPackage.Coverage = conCoverage
	deliveryPackage.Category = conCategory
	deliveryPackage.MajorVersion = conMajorVersion
	deliveryPackage.MinorVersion = conMinorVersion

	// Guided Tour
	var guidedTour structs.GuidedTour
	guidedTour.Access = conTourAccess
	guidedTour.Use = conTourUse
	guidedTour.ID = int(routeID)
	guidedTour.TripType = conTourTripType

	// Country
	var country structs.Country
	country.CountryCode = conCountryCode
	country.Name.LanguageCode = conLanguageCodeEnglish
	country.Name.Value = conCountryName
	guidedTour.Countries = append(guidedTour.Countries, country)

	// Names
	var tourName structs.TourName
	tourName.LanguageCode = conLanguageCodeEnglish
	tourName.Text = gpx.Metadata.Name
	guidedTour.Names = append(guidedTour.Names, tourName)

	// Length - sum up the distances of all tracks
	guidedTour.Length.Unit = conUnitDistance
	var totalDistanceKm float64
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(CalcTotalTrackDistance(track))/1000
	}
	guidedTour.Length.Value = totalDistanceKm

	// Duration - sum up the durations of all tracks
	guidedTour.Duration.Unit = conUnitDuration
	var totalDurationH float64
	for _, track := range gpx.Tracks {
		partialDurationH, err := CalcTotalTrackDuration(track)
		if err != nil {
			return deliveryPackage, err
		}
		totalDurationH = totalDurationH + float64(partialDurationH)/3600
	}
	guidedTour.Duration.Value = totalDurationH

	// Introductions
	var introduction structs.TourIntroduction
	introduction.LanguageCode = conLanguageCodeEnglish
	introduction.Text = conTextDefault
	guidedTour.Introductions = append(guidedTour.Introductions, introduction)

	// Descriptions
	var description structs.TourDescription
	description.LanguageCode = conLanguageCodeEnglish
	description.Text = conTextDefault
	guidedTour.Descriptions = append(guidedTour.Descriptions, description)

	// Pictures
	var picture structs.TourPicture
	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172
	guidedTour.Pictures = append(guidedTour.Pictures, picture)

	// Route
	var route structs.Route
	route.RouteID = int(routeID)
	route.Length = guidedTour.Length
	route.Duration = guidedTour.Duration
	route.CostModel = conRouteCostModel
	route.Criteria = conRouteCriteria
	route.AgoraCString = ""

	// Waypoints
	for index, gpxWaypoint := range gpx.Waypoints {
		var waypoint structs.RouteWayPoint
		waypoint.ID = "0_" + strconv.FormatInt(int64(index), 10)
		if index == 0 || index == len(gpx.Waypoints)-1 {

			waypoint.Importance = conImportanceAlways

			// Places with importance = always have to be EntryPoints
			var entryPoint structs.EntryPoint
			entryPoint.Route = "1"
			entryPoint.Value = waypoint.ID
			guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)

		} else {
			waypoint.Importance = conImportanceOptional
		}

		// Location
		var location structs.WayPointLocation
		location.GeoPosition.Latitude = gpxWaypoint.Latitude
		location.GeoPosition.Longitude = gpxWaypoint.Longitude

		// If this is a importance = always waypoint, we have to add address data
		if waypoint.Importance == conImportanceAlways {
			var addr structs.WayPointAddress
			addr.ParsedAddress.ParsedStreetAddress.ParsedStreetName.StreetName = gpxWaypoint.Name
			addr.ParsedAddress.ParsedPlace.PlaceLevel4 = gpxWaypoint.Name
			location.Address = &addr
		}

		waypoint.Locations = append(waypoint.Locations, location)

		// Description
		var wpDescription structs.TourDescription
		wpDescription.LanguageCode = conLanguageCodeEnglish
		wpDescription.Text = gpxWaypoint.Name
		waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)

		route.WayPoint = append(route.WayPoint, waypoint)
	}

	guidedTour.Routes = append(guidedTour.Routes, route)
	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)

	var err error
	return deliveryPackage, err
}
