package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"route2bimmer/structs"
	"strconv"
)

type fileStructure struct {
	filename string
	content  []byte
	mode     int64
}

func main() {

	var err error
	var gpxContents structs.GPX
	var xmlNav, xmlNavigation, thumbnail []byte
	var routeID int64 = 1234567

	// command line arguments
	inputPtr := flag.String("input", "", "path to input file")
	outputPtr := flag.String("output", "route.zip", "path to output zip file")
	flag.Parse()

	// read GPX file into gpxContents
	gpxContents, err = ReadGPXFile(inputPtr)
	if err != nil {
		log.Fatalln(err)
	} else {

		// Generate contents for XML file in folder "Nav" and "Navigation"
		var routeNav = MapGPXtoRouteNav(gpxContents, routeID)
		var routeNavigation = MapGPXtoRouteNavigation(gpxContents, routeID)

		// Marshal contents into XML text for both files
		xmlNav, err = xml.MarshalIndent(routeNav, "  ", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		xmlNavigation, err = xml.MarshalIndent(routeNavigation, "  ", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		// Read image file data
		thumbnail, err = ioutil.ReadFile("routepicture.jpg")
		if err != nil {
			log.Fatalln(err)
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

		// Write the zip file onto the harddrive
		err = ioutil.WriteFile(*outputPtr, bufZip.Bytes(), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
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
