package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"route2bimmer/structs"
)

// ReadGPXFromStdin read the contents of the GPX file from stdin and returns a structure of type GPX in case of success
func ReadGPXFromStdin() (structs.GPX, error) {
	// Declare return value
	var gpxContents structs.GPX

	// Read data from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return gpxContents, err
	}

	// Unmarshal the byteArray which contains the GPX file into the gpxContents
	err = xml.Unmarshal(data, &gpxContents)

	// Return the contents of the GPX file
	return gpxContents, err
}

// ReadGPXFromFile reads the contents of the supplied filepath and returns a structure of type GPX in case of success
func ReadGPXFromFile(inputPath *string) (structs.GPX, error) {
	// Declare return value
	var gpxContents structs.GPX
	var err error

	// Open the GPX file
	gpxFile, err := os.Open(*inputPath)
	if err != nil {
		return gpxContents, err
	}
	defer gpxFile.Close()

	// Read the GPX file as a byte array
	byteValue, err := ioutil.ReadAll(gpxFile)
	if err != nil {
		return gpxContents, err
	}

	// Unmarshal the byteArray which contains the GPX file into the gpxContents
	err = xml.Unmarshal(byteValue, &gpxContents)

	// Return the contents of the GPX file
	return gpxContents, err
}
