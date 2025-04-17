package geoimage

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func Exif_dir_processing(dirname string) {
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Error(err)
		}

		// If it's a file, print the path
		if !info.IsDir() && strings.HasSuffix(path, ".jpg") {
			gd, err := get_file_data(path)
			if err != nil {
				return err
			}
			gd.print()
		}
		return nil
	})

	if err != nil {
		logrus.Error(err)
	}
}

type imggeodata struct {
	fname     string
	longitude float64
	latitude  float64
}

func (gd *imggeodata) print() {
	fmt.Printf("%s Longitude: %f\n", gd.fname, gd.longitude)
	fmt.Printf("%s Latitude: %f\n", gd.fname, gd.latitude)
}

func get_file_data(fname string) (*imggeodata, error) {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	exifObj, err := exif.Decode(file)
	if err != nil {
		return nil, err
	}

	var latlong imggeodata
	latitude, longitude, err := exifObj.LatLong()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	latlong.longitude = longitude
	latlong.latitude = latitude
	latlong.fname = fname

	return &latlong, nil
}
