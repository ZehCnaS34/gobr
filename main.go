// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	// "encoding/base64"
	"fmt"
	"image"
	"log"
	// "strings"
	"flag"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	_ "image/jpeg"
)

type ImageFile string

func (i *ImageFile) Set(v string) error {
	*i = ImageFile(v)
	return nil
} 

func (i *ImageFile) String() string {
	v := *i
	return string(v)
}


var filename ImageFile

func init() {
	flag.Var(&filename, "file", "this is the file you want to perform on")

}

func main() {
	flag.Parse()
	fmt.Println(filename)

	// Decode the JPEG data. If reading from file, create a reader with
	reader, err := os.Open(filename.String())
	
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.

	var (
		pixelCount  uint32
		lightness   uint32
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			rp := r / 255
			gp := g / 255
			bp := b / 255

			cx := Max(rp, gp, bp)
			cn := Min(rp, gp, bp)

			lightness += (cx + cn) / 2
			pixelCount++
		}
	}

	fmt.Println(lightness/pixelCount)
}

func Max(is ...uint32) uint32 {
	max := is[0]
	for _, v := range is {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(is ...uint32) uint32 {
	min := is[0]
	for _, v := range is {
		if v > min {
			min = v
		}
	}
	return min
}
