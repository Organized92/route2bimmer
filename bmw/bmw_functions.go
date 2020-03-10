package bmw

import (
	"encoding/xml"
	"strconv"

	"github.com/Organized92/route2bimmer/gpx"
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
const conTourTripType string = "6"
const conCountryCode int = 3
const conCountryName string = "Germany"
const conUnitDistance string = "km"
const conUnitDuration string = "h"
const conTextDefault string = "-"
const conRouteCostModel int = 2
const conRouteCriteria int = 0

// NavFromGPX maps GPX data into the BMW format for "Nav" folder
func NavFromGPX(gpx gpx.GPX, routeID int64) (DeliveryPackage, error) {
	var deliveryPackage DeliveryPackage
	var err error

	// Basic data
	fillDeliveryPackage(&deliveryPackage, gpx, routeID)

	// Fill guided tour with data
	deliveryPackage.GuidedTour, err = getGuidedToursNav(gpx, routeID)

	return deliveryPackage, err
}

// NavigationFromGPX maps GPX data into the BMW format for "Nav" folder
func NavigationFromGPX(gpx gpx.GPX, routeID int64) (DeliveryPackage, error) {
	var deliveryPackage DeliveryPackage
	var err error

	// Basic data
	fillDeliveryPackage(&deliveryPackage, gpx, routeID)

	// Fill guided tour with data
	deliveryPackage.GuidedTour, err = getGuidedToursNavigation(gpx, routeID)

	return deliveryPackage, err
}

func fillDeliveryPackage(bmw *DeliveryPackage, gpx gpx.GPX, routeID int64) {
	bmw.VersionNo = conVersionZeroDotZero
	if gpx.Metadata.Time != "" {
		bmw.CreationTime = gpx.Metadata.Time
	}
	bmw.MapVersion = conVersionZeroDotZero
	bmw.LanguageCodeDesc = conLanguageCodeDesc
	bmw.CountryCodeDesc = conCountryCodeDesc
	bmw.SupplierCodeDesc = conSupplierCodeDesc
	bmw.XYType = conXYType
	bmw.CategoryCodeDesc = conCategoryCodeDesc
	bmw.CharSet = conCharSet
	bmw.UpdateType = conUpdateType
	bmw.Coverage = conCoverage
	bmw.Category = conCategory
	bmw.MajorVersion = conMajorVersion
	bmw.MinorVersion = conMinorVersion
}

func getGuidedToursNav(gpx gpx.GPX, routeID int64) ([]GuidedTour, error) {
	var guidedTours []GuidedTour
	var guidedTour GuidedTour
	var err error

	guidedTour.Access = conTourAccess
	guidedTour.Use = conTourUse
	guidedTour.ID = strconv.FormatInt(routeID, 10)
	guidedTour.TripType = conTourTripType

	guidedTour.Countries, err = getCountries(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Names, err = getNames(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Length, err = getLength(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Duration, err = getDuration(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Introductions, err = getIntroductions(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Descriptions, err = getDescriptions(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Pictures, err = getPictures(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Routes, err = getRoutesNav(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.EntryPoints, err = getEntryPoints(guidedTour.Routes, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTours = append(guidedTours, guidedTour)
	return guidedTours, err
}

func getGuidedToursNavigation(gpx gpx.GPX, routeID int64) ([]GuidedTour, error) {
	var guidedTours []GuidedTour
	var guidedTour GuidedTour
	var err error

	guidedTour.Access = conTourAccess
	guidedTour.Use = conTourUse
	guidedTour.ID = strconv.FormatInt(routeID, 10)
	guidedTour.TripType = conTourTripType

	guidedTour.Countries, err = getCountries(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Names, err = getNames(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Length, err = getLength(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Duration, err = getDuration(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Introductions, err = getIntroductions(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Descriptions, err = getDescriptions(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Pictures, err = getPictures(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.Routes, err = getRoutesNavigation(gpx, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTour.EntryPoints, err = getEntryPoints(guidedTour.Routes, routeID)
	if err != nil {
		return guidedTours, err
	}

	guidedTours = append(guidedTours, guidedTour)
	return guidedTours, err
}

func getCountries(gpx gpx.GPX, routeID int64) ([]Country, error) {
	var countries []Country
	var country Country
	var err error

	country.CountryCode = conCountryCode
	country.Name.LanguageCode = conLanguageCodeEnglish
	country.Name.Value = conCountryName

	countries = append(countries, country)
	return countries, err
}

func getNames(gpx gpx.GPX, routeID int64) ([]TourName, error) {
	var names []TourName
	var name TourName
	var err error

	name.LanguageCode = conLanguageCodeEnglish
	name.Text = gpx.Metadata.Name

	names = append(names, name)
	return names, err
}

func getLength(gpx gpx.GPX, routeID int64) (TourLength, error) {
	var length TourLength
	var totalDistanceKm float64
	var err error

	length.Unit = conUnitDistance
	for _, track := range gpx.Tracks {
		totalDistanceKm = totalDistanceKm + float64(track.CalcTotalDistance())/1000
	}
	length.Value = totalDistanceKm
	return length, err
}

func getDuration(gpx gpx.GPX, routeID int64) (TourDuration, error) {
	var duration TourDuration
	var totalDurationH float64
	var err error

	duration.Unit = conUnitDuration
	for _, track := range gpx.Tracks {
		partialDurationH, err := track.CalcTotalDuration()
		if err != nil {
			return duration, err
		}
		totalDurationH = totalDurationH + float64(partialDurationH)/3600
	}
	duration.Value = totalDurationH
	return duration, err
}

func getIntroductions(gpx gpx.GPX, routeID int64) ([]TourIntroduction, error) {
	var introductions []TourIntroduction
	var introduction TourIntroduction
	var err error

	introduction.LanguageCode = conLanguageCodeEnglish
	introduction.Text = conTextDefault

	introductions = append(introductions, introduction)
	return introductions, err
}

func getDescriptions(gpx gpx.GPX, routeID int64) ([]TourDescription, error) {
	var descriptions []TourDescription
	var description TourDescription
	var err error

	description.LanguageCode = conLanguageCodeEnglish
	description.Text = conTextDefault

	descriptions = append(descriptions, description)
	return descriptions, err
}

func getPictures(gpx gpx.GPX, routeID int64) ([]TourPicture, error) {
	var pictures []TourPicture
	var picture TourPicture
	var err error

	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
	picture.Encoding = "JPEG"
	picture.Width = 252
	picture.Height = 172

	pictures = append(pictures, picture)
	return pictures, err
}

func getRoutesNav(gpx gpx.GPX, routeID int64) ([]Route, error) {
	var routes []Route
	var err error

	// The GPX file may contain multiple routes. We have to loop over them
	for _, gpxRoute := range gpx.Routes {
		var route Route

		route.RouteID = strconv.FormatInt(routeID, 10)
		route.CostModel = conRouteCostModel
		route.Criteria = conRouteCriteria
		route.AgoraCString = ""

		// Route length
		route.Length, err = getRouteLength(gpx, gpxRoute)
		if err != nil {
			return routes, err
		}

		// Route duration
		route.Duration, err = getRouteDuration(gpx, gpxRoute)
		if err != nil {
			return routes, err
		}

		// Loop over the routes Waypoints
		for rteWptIndex, gpxWaypoint := range gpxRoute.RouteWaypoints {
			var waypoint RouteWayPoint
			waypoint.ID = strconv.FormatInt(int64(rteWptIndex), 10)

			// If this is the first or the last waypoint in the route,
			// it has to be importance = always, otherwise optional
			if rteWptIndex == 0 || rteWptIndex == len(gpxRoute.RouteWaypoints)-1 {
				waypoint.Importance = conImportanceAlways
			} else {
				waypoint.Importance = conImportanceOptional
			}

			// Waypoint description
			if gpxWaypoint.Description != "" {
				var description TourDescription
				description.LanguageCode = conLanguageCodeEnglish
				description.Text = gpxWaypoint.Description
				waypoint.Descriptions = append(waypoint.Descriptions, description)
			}

			// Location
			var location WayPointLocation
			location.GeoPosition.Latitude = gpxWaypoint.Latitude
			location.GeoPosition.Longitude = gpxWaypoint.Longitude

			// If this is a importance = always waypoint, we have to include
			// address data to this location
			if waypoint.Importance == conImportanceAlways {
				var address WayPointAddress
				address.ParsedAddress.ParsedStreetAddress.ParsedStreetName.StreetName = gpxWaypoint.Name
				address.ParsedAddress.ParsedPlace.PlaceLevel4 = gpxWaypoint.Name
				location.Address = &address
			}

			waypoint.Locations = append(waypoint.Locations, location)

			route.WayPoint = append(route.WayPoint, waypoint)
		}

		routes = append(routes, route)
	}

	return routes, err
}

func getRoutesNavigation(gpx gpx.GPX, routeID int64) ([]Route, error) {
	var routes []Route
	var err error

	// The GPX file may contain multiple routes. We have to loop over them
	for rteIndex, gpxRoute := range gpx.Routes {
		var route Route

		route.RouteID = strconv.FormatInt(routeID, 10)
		route.CostModel = conRouteCostModel
		route.Criteria = conRouteCriteria
		route.AgoraCString = ""

		// Route length
		route.Length, err = getRouteLength(gpx, gpxRoute)
		if err != nil {
			return routes, err
		}

		// Route duration
		route.Duration, err = getRouteDuration(gpx, gpxRoute)
		if err != nil {
			return routes, err
		}

		// Loop over the routes Waypoints
		for rteWptIndex, gpxWaypoint := range gpxRoute.RouteWaypoints {
			var waypoint RouteWayPoint
			waypoint.ID = strconv.FormatInt(int64(rteIndex), 10) + "_" + strconv.FormatInt(int64(rteWptIndex), 10)

			// If this is the first or the last waypoint in the route,
			// it has to be importance = always, otherwise optional
			if rteWptIndex == 0 || rteWptIndex == len(gpxRoute.RouteWaypoints)-1 {
				waypoint.Importance = conImportanceAlways
			} else {
				waypoint.Importance = conImportanceOptional
			}

			// Waypoint description
			if gpxWaypoint.Description != "" {
				var description TourDescription
				description.LanguageCode = conLanguageCodeEnglish
				description.Text = gpxWaypoint.Description
				waypoint.Descriptions = append(waypoint.Descriptions, description)
			}

			// Location
			var location WayPointLocation
			location.GeoPosition.Latitude = gpxWaypoint.Latitude
			location.GeoPosition.Longitude = gpxWaypoint.Longitude

			// If this is a importance = always waypoint, we have to include
			// address data to this location
			if waypoint.Importance == conImportanceAlways {
				var address WayPointAddress
				address.ParsedAddress.ParsedStreetAddress.ParsedStreetName.StreetName = gpxWaypoint.Name
				address.ParsedAddress.ParsedPlace.PlaceLevel4 = gpxWaypoint.Name
				location.Address = &address
			}

			waypoint.Locations = append(waypoint.Locations, location)

			route.WayPoint = append(route.WayPoint, waypoint)
		}

		routes = append(routes, route)
	}

	return routes, err
}

func getEntryPoints(routes []Route, routeID int64) ([]EntryPoint, error) {
	var entryPoints []EntryPoint
	var err error

	// Every route has it's entry points, so we loop over them
	for rteIndex, route := range routes {
		// Every waypoint with importance = always needs to be an entry point
		for _, waypoint := range route.WayPoint {
			if waypoint.Importance == conImportanceAlways {
				var entryPoint EntryPoint
				entryPoint.Route = strconv.FormatInt(int64(rteIndex)+1, 10)
				entryPoint.Value = waypoint.ID
				entryPoints = append(entryPoints, entryPoint)
			}
		}
	}

	return entryPoints, err
}

func getRouteLength(gpx gpx.GPX, route gpx.Route) (TourLength, error) {
	var length TourLength
	var err error

	length.Unit = conUnitDistance
	// We have the track that corresponds to this route. If the GPX file
	// contains only one track, we use this one. If it has more than one,
	// we have to identify the correct track by name. This requires that
	// route name and track name are equal.
	if len(gpx.Tracks) == 1 {
		length.Value = float64(gpx.Tracks[0].CalcTotalDistance()) / 1000
	} else {
		// More than one? Or no track at all?
		if len(gpx.Tracks) > 1 {

			// There is more than one track. We have to find the correct one.
			for _, track := range gpx.Tracks {
				if track.Name == route.Name {
					length.Value = float64(track.CalcTotalDistance()) / 1000
				}
			}

		} else {
			// No track at all
			length.Value = 0
		}
	}

	return length, err
}

func getRouteDuration(gpx gpx.GPX, route gpx.Route) (TourDuration, error) {
	var duration TourDuration
	var calcDuration int64
	var err error

	duration.Unit = conUnitDuration
	// We have the track that corresponds to this route. If the GPX file
	// contains only one track, we use this one. If it has more than one,
	// we have to identify the correct track by name. This requires that
	// route name and track name are equal.
	if len(gpx.Tracks) == 1 {
		calcDuration, err = gpx.Tracks[0].CalcTotalDuration()
		duration.Value = float64(calcDuration) / 3600
	} else {
		// More than one? Or no track at all?
		if len(gpx.Tracks) > 1 {

			// There is more than one track. We have to find the correct one.
			for _, track := range gpx.Tracks {
				if track.Name == route.Name {
					calcDuration, err = track.CalcTotalDuration()
					duration.Value = float64(calcDuration) / 3600
				}
			}

		} else {
			// No track at all
			duration.Value = 0
		}
	}

	return duration, err
}

//
//
//
//
//
//

// NavigationFromGPX maps GPX data into the BMW format for "Navigation" folder
// func NavigationFromGPX(gpx gpx.GPX, routeID int64) (DeliveryPackage, error) {
// 	var deliveryPackage DeliveryPackage
//
// 	// Basic data
// 	deliveryPackage.VersionNo = conVersionZeroDotZero
// 	deliveryPackage.CreationTime = gpx.Metadata.Time
// 	deliveryPackage.MapVersion = conVersionZeroDotZero
// 	deliveryPackage.LanguageCodeDesc = conLanguageCodeDesc
// 	deliveryPackage.CountryCodeDesc = conCountryCodeDesc
// 	deliveryPackage.SupplierCodeDesc = conSupplierCodeDesc
// 	deliveryPackage.XYType = conXYType
// 	deliveryPackage.CategoryCodeDesc = conCategoryCodeDesc
// 	deliveryPackage.CharSet = conCharSet
// 	deliveryPackage.UpdateType = conUpdateType
// 	deliveryPackage.Coverage = conCoverage
// 	deliveryPackage.Category = conCategory
// 	deliveryPackage.MajorVersion = conMajorVersion
// 	deliveryPackage.MinorVersion = conMinorVersion
//
// 	// Guided Tour
// 	var guidedTour GuidedTour
// 	guidedTour.Access = conTourAccess
// 	guidedTour.Use = conTourUse
// 	guidedTour.ID = int(routeID)
// 	guidedTour.TripType = conTourTripType
//
// 	// Country
// 	var country Country
// 	country.CountryCode = conCountryCode
// 	country.Name.LanguageCode = conLanguageCodeEnglish
// 	country.Name.Value = conCountryName
// 	guidedTour.Countries = append(guidedTour.Countries, country)
//
// 	// Names
// 	var tourName TourName
// 	tourName.LanguageCode = conLanguageCodeEnglish
// 	tourName.Text = gpx.Metadata.Name
// 	guidedTour.Names = append(guidedTour.Names, tourName)
//
// 	// Length - sum up the distances of all tracks
// 	guidedTour.Length.Unit = conUnitDistance
// 	var totalDistanceKm float64
// 	for _, track := range gpx.Tracks {
// 		totalDistanceKm = totalDistanceKm + float64(track.CalcTotalDistance())/1000
// 	}
// 	guidedTour.Length.Value = totalDistanceKm
//
// 	// Duration - sum up the durations of all tracks
// 	guidedTour.Duration.Unit = conUnitDuration
// 	var totalDurationH float64
// 	for _, track := range gpx.Tracks {
// 		partialDurationH, err := track.CalcTotalDuration()
// 		if err != nil {
// 			return deliveryPackage, err
// 		}
// 		totalDurationH = totalDurationH + float64(partialDurationH)/3600
// 	}
// 	guidedTour.Duration.Value = totalDurationH
//
// 	// Introductions
// 	var introduction TourIntroduction
// 	introduction.LanguageCode = conLanguageCodeEnglish
// 	introduction.Text = conTextDefault
// 	guidedTour.Introductions = append(guidedTour.Introductions, introduction)
//
// 	// Descriptions
// 	var description TourDescription
// 	description.LanguageCode = conLanguageCodeEnglish
// 	description.Text = conTextDefault
// 	guidedTour.Descriptions = append(guidedTour.Descriptions, description)
//
// 	// Pictures
// 	var picture TourPicture
// 	picture.Reference = "routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg"
// 	picture.Encoding = "JPEG"
// 	picture.Width = 252
// 	picture.Height = 172
// 	guidedTour.Pictures = append(guidedTour.Pictures, picture)
//
// 	// Route
// 	var route Route
// 	route.RouteID = int(routeID)
// 	route.Length = guidedTour.Length
// 	route.Duration = guidedTour.Duration
// 	route.CostModel = conRouteCostModel
// 	route.Criteria = conRouteCriteria
// 	route.AgoraCString = ""
//
// 	// // Waypoints
// 	// for index, gpxWaypoint := range gpx.Waypoints {
// 	// 	var waypoint structs.RouteWayPoint
// 	// 	waypoint.ID = "0_" + strconv.FormatInt(int64(index), 10)
// 	// 	if index == 0 || index == len(gpx.Waypoints)-1 {
// 	//
// 	// 		waypoint.Importance = conImportanceAlways
// 	//
// 	// 		// Places with importance = always have to be EntryPoints
// 	// 		var entryPoint structs.EntryPoint
// 	// 		entryPoint.Route = "1"
// 	// 		entryPoint.Value = waypoint.ID
// 	// 		guidedTour.EntryPoints = append(guidedTour.EntryPoints, entryPoint)
// 	//
// 	// 	} else {
// 	// 		waypoint.Importance = conImportanceOptional
// 	// 	}
// 	//
// 	// 	// Location
// 	// 	var location structs.WayPointLocation
// 	// 	location.GeoPosition.Latitude = gpxWaypoint.Latitude
// 	// 	location.GeoPosition.Longitude = gpxWaypoint.Longitude
// 	//
// 	// 	// If this is a importance = always waypoint, we have to add address data
// 	// 	if waypoint.Importance == conImportanceAlways {
// 	// 		var addr structs.WayPointAddress
// 	// 		addr.ParsedAddress.ParsedStreetAddress.ParsedStreetName.StreetName = gpxWaypoint.Name
// 	// 		addr.ParsedAddress.ParsedPlace.PlaceLevel4 = gpxWaypoint.Name
// 	// 		location.Address = &addr
// 	// 	}
// 	//
// 	// 	waypoint.Locations = append(waypoint.Locations, location)
// 	//
// 	// 	// Description
// 	// 	var wpDescription structs.TourDescription
// 	// 	wpDescription.LanguageCode = conLanguageCodeEnglish
// 	// 	wpDescription.Text = gpxWaypoint.Name
// 	// 	waypoint.Descriptions = append(waypoint.Descriptions, wpDescription)
// 	//
// 	// 	route.WayPoint = append(route.WayPoint, waypoint)
// 	// }
//
// 	guidedTour.Routes = append(guidedTour.Routes, route)
// 	deliveryPackage.GuidedTour = append(deliveryPackage.GuidedTour, guidedTour)
//
// 	var err error
// 	return deliveryPackage, err
// }
//

// ToXML converts the BMW structure to some real xml Text
func (bmw DeliveryPackage) ToXML() ([]byte, error) {
	buffer, err := xml.MarshalIndent(bmw, "", "  ")
	return buffer, err
}
