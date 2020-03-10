package gpx

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// FromStdin reads all data from Stdin and converts it into a GPX file structure
func FromStdin() (GPX, error) {
	// Declare return value
	var gpxContents GPX

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

// FromFile reads the contents of the supplied filepath and returns a structure of type GPX in case of success
func FromFile(inputPath *string) (GPX, error) {
	// Declare return value
	var gpxContents GPX
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

// GetName returns the name of the route, first trying to read it from the GPX metadata,
// and if not present here, from the first route. If both are empty, it will return
// "Unnamed Route"
func (gpx GPX) GetName() string {
	if gpx.Metadata.Name != "" {
		return gpx.Metadata.Name
	}

	// No route name present in metadata, try to read it from the first route
	// in the file.
	if len(gpx.Routes) >= 1 {
		if gpx.Routes[0].Name != "" {
			return gpx.Routes[0].Name
		}
	}

	// Fallback
	return "Unnamed Route"
}
