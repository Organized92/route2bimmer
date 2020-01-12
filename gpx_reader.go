package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"route2bimmer/structs"
)

// ReadGPXFile reads the contents of the supplied filepath and returns a structure of type GPX in case of success
func ReadGPXFile(inputPath *string) (structs.GPX, error) {

	// Declare return value
	var gpxContents structs.GPX

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
