package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"route2bimmer/structs"
	"strconv"
	"strings"
)

type fileStructure struct {
	filename string
	content  []byte
	mode     int64
}

func main() {

	// Variable for error handling
	var err error

	// Parse command line arguments
	inputPtr := flag.String("input", "", "path to input file")
	outputPtr := flag.String("output", "", "path to output zip file")
	flag.Parse()

	// Check if we have to read the input data from stdin or from a file
	// Also we do some argument checks
	var directio bool
	if *inputPtr == "" && *outputPtr == "" {
		directio = true
	} else {
		if *inputPtr == "" && *outputPtr != "" {
			fmt.Println("Please specify the input GPX file. Use -h for more information.")
			return
		}
		if *inputPtr != "" && *outputPtr == "" {
			fmt.Println("Please specify the output ZIP file. Use -h for more information.")
			return
		}
		directio = false
	}

	var gpxContents structs.GPX
	if directio == true {
		// Read from stdin
		gpxContents, err = ReadGPXFromStdin()
	} else {
		// Read from file
		gpxContents, err = ReadGPXFromFile(inputPtr)
	}

	// Generate random ID for this route
	var routeID = generateRandomID()

	// Generate contents for XML file in folder "Nav" and "Navigation"
	var routeNav = MapGPXtoRouteNav(gpxContents, routeID)
	var routeNavigation = MapGPXtoRouteNavigation(gpxContents, routeID)

	// Marshal contents into XML text for both files
	var xmlNav = convertRouteToXML(routeNav)
	var xmlNavigation = convertRouteToXML(routeNavigation)

	// We have to replace one XML tag so that it contains a newline
	xmlNav = replaceAgoraCString(xmlNav)

	// Read image file data
	var thumbnail []byte
	thumbnail, err = ioutil.ReadFile("routepicture.jpg")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// File structure for tar archive
	var filesNav = []fileStructure{
		{strconv.FormatInt(routeID, 10) + ".xml", xmlNav, 0700},
		{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail, 0700},
	}
	var filesNavigation = []fileStructure{
		{strconv.FormatInt(routeID, 10) + ".xml", xmlNavigation, 0700},
		{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail, 0700},
	}

	// Create tarballs
	var bufNav = filesToTarBuffer(filesNav)
	var bufNavigation = filesToTarBuffer(filesNavigation)

	// compress the tarballs using GZIP
	var bufNavGzip = compressGzip(bufNav)
	var bufNavigationGzip = compressGzip(bufNavigation)

	// File structure for zip archive
	var filesZip = []fileStructure{
		{"BMWData/Nav/" + gpxContents.Metadata.Name + ".tar.gz", bufNavGzip.Bytes(), 0700},
		{"BMWData/Navigation/Routes/" + gpxContents.Metadata.Name + ".tar.gz", bufNavigationGzip.Bytes(), 0700},
	}

	// Create the ZIP file containing the folder structure and the tar.gz-files
	var bufZip = filesToZipBuffer(filesZip)

	// Check if we have to write the zip file to stdout or into a file
	if directio == true {
		// write to stdout
		bufZip.WriteTo(os.Stdout)
	} else {
		// write the zip file onto the harddrive
		err = ioutil.WriteFile(*outputPtr, bufZip.Bytes(), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// generateRandomID generates a random 7-digit number used as an ID for this route
func generateRandomID() int64 {
	return int64(rand.Intn(9999999-1000000) + 1000000)
}

// convertRouteToXML converts a xml structure to some real xml Text
func convertRouteToXML(route structs.DeliveryPackage) []byte {
	var buffer, err = xml.MarshalIndent(route, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	return buffer
}

// replaceAgoraCString corrects the xml element "AgoraCString"
func replaceAgoraCString(input []byte) []byte {
	var s = string(input[:])
	s = strings.ReplaceAll(s, "<AgoraCString></AgoraCString>", "<AgoraCString>\n          </AgoraCString>")
	var output = []byte(s)
	return output
}

// filesToTarBuffer writes the data contained in "files" in tar format into buffer
func filesToTarBuffer(files []fileStructure) bytes.Buffer {
	var buffer bytes.Buffer

	// Create tar writer
	tarWriter := tar.NewWriter(&buffer)

	// Loop over the supplied file list
	for _, file := range files {
		// Write file header
		var fileHdr = &tar.Header{
			Name: file.filename,
			Mode: file.mode,
			Size: int64(len(file.content)),
		}
		if err := tarWriter.WriteHeader(fileHdr); err != nil {
			log.Fatalln(err)
		}
		// Write file data
		if _, err := tarWriter.Write(file.content); err != nil {
			log.Fatalln(err)
		}
	}

	// After all files have been written, close the tar
	if err := tarWriter.Close(); err != nil {
		log.Fatalln(err)
	}

	return buffer
}

// compressGzip compresses the input data using gzip
func compressGzip(data bytes.Buffer) bytes.Buffer {
	var buffer bytes.Buffer

	// Create the gzip writer
	gzipWriter := gzip.NewWriter(&buffer)

	// Write data
	if _, err := gzipWriter.Write(data.Bytes()); err != nil {
		log.Fatalln(err)
	}

	// Close the writer
	gzipWriter.Close()

	// Return compressed data
	return buffer
}

// filesToZipBuffer writes the data contained in "files" in zip format into buffer
func filesToZipBuffer(files []fileStructure) bytes.Buffer {
	var buffer bytes.Buffer

	// Create zip writer
	zipWriter := zip.NewWriter(&buffer)

	// Loop over the files to add them into the archive
	for _, file := range files {
		// Create the file
		f, err := zipWriter.Create(file.filename)
		if err != nil {
			log.Fatalln(err)
		}
		// Write data into the file
		if _, err := f.Write(file.content); err != nil {
			log.Fatalln(err)
		}
	}

	// Close the zip writer
	zipWriter.Close()

	// Return the zip file
	return buffer
}
