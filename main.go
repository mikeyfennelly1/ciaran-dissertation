package main

import (
	"github.com/mikeyfennelly1/ciaran-dissertation/geoimage"
)

const img_dir = "./exif-samples/jpg/gps"

func main() {
	geoimage.Exif_dir_processing(img_dir)
}
