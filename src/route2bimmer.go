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
	"strconv"
)

type fileStructure struct {
	filename string
	content  []byte
}

func main() {

	var err error
	var gpxContents GPX
	var xmlNav, xmlNavigation, thumbnail []byte
	var routeID int64 = 1234567

	// command line arguments
	inputPtr := flag.String("input", "", "path to input file")
	outputPtr := flag.String("output", "route.zip", "path to output zip file")
	flag.Parse()

	// read GPX file into gpxContents
	gpxContents, err = readGPXFile(inputPtr)
	if err != nil {
		log.Fatalln(err)
	} else {

		// Generate contents for XML file in folder "Nav" and "Navigation"
		var routeNav = mapGPXtoRouteNav(gpxContents, routeID)
		var routeNavigation = mapGPXtoRouteNavigation(gpxContents, routeID)

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

		// Create tar archives containing the XML file and a picture
		var bufNav, bufNavigation bytes.Buffer
		tarWriterNav := tar.NewWriter(&bufNav)
		tarWriterNavigation := tar.NewWriter(&bufNavigation)

		// File structure
		var filesNav = []fileStructure{
			{strconv.FormatInt(routeID, 10) + ".xml", xmlNav},
			{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail},
		}
		var filesNavigation = []fileStructure{
			{strconv.FormatInt(routeID, 10) + ".xml", xmlNavigation},
			{"routepicture_" + strconv.FormatInt(routeID, 10) + ".jpg", thumbnail},
		}

		// Create tarballs
		for _, file := range filesNav {
			var fileHdr = &tar.Header{
				Name: file.filename,
				Mode: 0700,
				Size: int64(len(file.content)),
			}
			if err := tarWriterNav.WriteHeader(fileHdr); err != nil {
				log.Fatalln(err)
			}
			if _, err := tarWriterNav.Write(file.content); err != nil {
				log.Fatalln(err)
			}
		}

		for _, file := range filesNavigation {
			var fileHdr = &tar.Header{
				Name: file.filename,
				Mode: 0700,
				Size: int64(len(file.content)),
			}
			if err := tarWriterNavigation.WriteHeader(fileHdr); err != nil {
				log.Fatalln(err)
			}
			if _, err := tarWriterNavigation.Write(file.content); err != nil {
				log.Fatalln(err)
			}
		}

		// Close tarballs
		if err := tarWriterNav.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := tarWriterNavigation.Close(); err != nil {
			log.Fatalln(err)
		}

		// compress the tarballs using GZIP
		var bufNavGzip, bufNavigationGzip bytes.Buffer
		gzWriterNav := gzip.NewWriter(&bufNavGzip)
		gzWriterNavigation := gzip.NewWriter(&bufNavigationGzip)

		if _, err := gzWriterNav.Write(bufNav.Bytes()); err != nil {
			log.Fatalln(err)
		}

		if _, err := gzWriterNavigation.Write(bufNavigation.Bytes()); err != nil {
			log.Fatalln(err)
		}

		gzWriterNav.Close()
		gzWriterNavigation.Close()

		// create the ZIP file containing the folder structure and the tar.gz-files
		var bufZip bytes.Buffer
		zipWriter := zip.NewWriter(&bufZip)

		var filesZip = []fileStructure{
			{"BMWData/Nav/" + gpxContents.Metadata.Name + ".tar.gz", bufNavGzip.Bytes()},
			{"BMWData/Navigation/Routes/" + gpxContents.Metadata.Name + ".tar.gz", bufNavigationGzip.Bytes()},
		}

		for _, file := range filesZip {
			f, err := zipWriter.Create(file.filename)
			if err != nil {
				log.Fatalln(err)
			}
			if _, err := f.Write(file.content); err != nil {
				log.Fatalln(err)
			}
		}

		zipWriter.Close()

		err = ioutil.WriteFile(*outputPtr, bufZip.Bytes(), 0644)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(*outputPtr)

	}
}
