package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"os"
	"strconv"
	"strings"

	"github.com/Organized92/route2bimmer/bmw"
	"github.com/Organized92/route2bimmer/gpx"
)

type fileData struct {
	filename string
	content  []byte
	mode     int64
}

func main() {

	// Variable for error handling
	var err error

	// ***************************************************************************
	// Command line arguments
	// ***************************************************************************
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
			log.Fatalln("Please specify the input GPX file. Use -h for more information.")
		}
		if *inputPtr != "" && *outputPtr == "" {
			log.Fatalln("Please specify the output ZIP file. Use -h for more information.")
		}
		directio = false
	}

	// ***************************************************************************
	// Read and interpret GPX file
	// ***************************************************************************
	var gpxFile gpx.GPX
	if directio == true {
		// Read from STDIN
		gpxFile, err = gpx.FromStdin()
		if err != nil {
			log.Println("Could not read fom STDIN!")
			log.Fatalln(err)
		}
	} else {
		// Read from file
		gpxFile, err = gpx.FromFile(inputPtr)
		if err != nil {
			log.Println("Could not read the GPX file!")
			log.Fatalln(err)
		}
	}

	// Generate random ID for this route
	routeID := generateRandomID()

	// ***************************************************************************
	// Generate contents for XML file in folder "Nav" and "Navigation"
	// ***************************************************************************
	routeNav, err := bmw.NavFromGPX(gpxFile, routeID)
	if err != nil {
		log.Println("BMW route XML could not be generated (Nav)!")
		log.Fatalln(err)
	}

	//routeNavigation, err := bmw.NavigationFromGPX(gpxFile, routeID)
	routeNavigation, err := bmw.NavigationFromGPX(gpxFile, routeID)
	if err != nil {
		log.Println("BMW route XML could not be generated (Nav)!")
		log.Fatalln(err)
	}

	// ***************************************************************************
	// Marshal contents into XML text for both files
	// ***************************************************************************
	xmlNav, err := routeNav.ToXML()
	if err != nil {
		log.Println("GPX contents could not be converted to BMW route format (Nav)!")
		log.Fatalln(err)
	}

	xmlNavigation, err := routeNavigation.ToXML()
	if err != nil {
		log.Println("GPX contents could not be converted to BMW route format (Navigation)!")
		log.Fatalln(err)
	}

	// We have to replace one XML tag so that it contains a newline
	xmlNav = replaceAgoraCString(xmlNav)

	// ***************************************************************************
	// Read image file data
	// ***************************************************************************
	var thumbnail []byte
	thumbnail, err = ioutil.ReadFile("routepicture.jpg")
	if err != nil {
		log.Println("Default route picture could not be loaded!")
		log.Fatalln(err)
	}

	// ***************************************************************************
	// Create TAR archive
	// ***************************************************************************
	var filesNav = []fileData{
		{strconv.FormatInt(routeID, 10) + ".xml", xmlNav, 0700},
		{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail, 0700},
	}
	var filesNavigation = []fileData{
		{strconv.FormatInt(routeID, 10) + ".xml", xmlNavigation, 0700},
		{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail, 0700},
	}

	// Create TAR archives
	bufNav, err := filesToTarBuffer(filesNav)
	if err != nil {
		log.Println("Could not create the tarball file (Nav)!")
		log.Fatalln(err)
	}

	bufNavigation, err := filesToTarBuffer(filesNavigation)
	if err != nil {
		log.Println("Could not create the tarball file (Navigation)!")
		log.Fatalln(err)
	}

	// ***************************************************************************
	// Compress the TAR files using GZIP
	// ***************************************************************************
	bufNavGzip, err := compressGzip(bufNav)
	if err != nil {
		log.Println("Could not gzip the tarball file (Nav)!")
		log.Fatalln(err)
	}

	bufNavigationGzip, err := compressGzip(bufNavigation)
	if err != nil {
		log.Println("Could not gzip the tarball file (Navigation)!")
		log.Fatalln(err)
	}

	// ***************************************************************************
	// Create ZIP archive
	// ***************************************************************************
	// File structure
	var filesZip = []fileData{
		{"BMWData/Nav/" + strconv.FormatInt(routeID, 10) + ".tar.gz", bufNavGzip.Bytes(), 0700},
		{"BMWData/Navigation/Routes/" + strconv.FormatInt(routeID, 10) + ".tar.gz", bufNavigationGzip.Bytes(), 0700},
	}

	// Create the ZIP file containing the folder structure and the tar.gz-files
	bufZip, err := filesToZipBuffer(filesZip)
	if err != nil {
		log.Println("Could not create the zip file!")
		log.Fatalln(err)
	}

	// ***************************************************************************
	// ZIP output
	// ***************************************************************************
	// Check if we have to write the zip file to stdout or into a file
	if directio == true {
		// write to stdout
		bufZip.WriteTo(os.Stdout)
	} else {
		// write the zip file onto the harddrive
		err = ioutil.WriteFile(*outputPtr, bufZip.Bytes(), 0644)
		if err != nil {
			log.Println("Could not write the zip file to the harddrive!")
			log.Fatalln(err)
		}
	}
}

// generateRandomID generates a random 7-digit number used as an ID for this route
func generateRandomID() int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(9999999-1000000) + 1000000)
}

// replaceAgoraCString corrects the xml element "AgoraCString"
func replaceAgoraCString(input []byte) []byte {
	var s = string(input[:])
	s = strings.ReplaceAll(s, "<AgoraCString></AgoraCString>", "<AgoraCString>\n          </AgoraCString>")
	var output = []byte(s)
	return output
}

// filesToTarBuffer writes the data contained in "files" in tar format into buffer
func filesToTarBuffer(files []fileData) (bytes.Buffer, error) {
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
			return buffer, err
		}
		// Write file data
		if _, err := tarWriter.Write(file.content); err != nil {
			return buffer, err
		}
	}

	// After all files have been written, close the tar
	err := tarWriter.Close()
	return buffer, err
}

// compressGzip compresses the input data using gzip
func compressGzip(data bytes.Buffer) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	// Create the gzip writer
	gzipWriter := gzip.NewWriter(&buffer)

	// Write data
	_, err := gzipWriter.Write(data.Bytes())
	if err != nil {
		return buffer, err
	}

	// Close the writer
	err = gzipWriter.Close()

	// Return compressed data
	return buffer, err
}

// filesToZipBuffer writes the data contained in "files" in zip format into buffer
func filesToZipBuffer(files []fileData) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	// Create zip writer
	zipWriter := zip.NewWriter(&buffer)

	// Loop over the files to add them into the archive
	for _, file := range files {
		// Create the file
		f, err := zipWriter.Create(file.filename)
		if err != nil {
			return buffer, err
		}
		// Write data into the file
		if _, err := f.Write(file.content); err != nil {
			return buffer, err
		}
	}

	// Close the zip writer
	err := zipWriter.Close()

	// Return the zip file
	return buffer, err
}
