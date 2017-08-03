package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/h2non/filetype.v1"
)

var counter int

type pixel struct {
	r, g, b, a uint32
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	images := getImages("./images/")

	// range over the [] holding the []pixel - eg, give me each image
	// range over the []pixel hold the pixels - eg, give me each pixel
	for i, img := range images {
		for j, pixel := range img {
			fmt.Println("Image", i, "\t pixel", j, "\t r g b a:", pixel)
			if j == 10 {
				break
			}
		}
	}
	fmt.Println("\nPixels counted:\t\t", counter, "\nRGBA values counted:\t", counter*4)
}

func getImages(dir string) [][]pixel {
	var images [][]pixel

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		buf, _ := ioutil.ReadFile(path)
		kind, err := filetype.Match(buf)
		if err != nil {
			fmt.Printf("Unknown: %s", err)
		}

		if kind.Extension != "unknown" {
			img := loadImage(path, kind.Extension)
			pixels := getPixels(img)
			images = append(images, pixels)
		}

		return nil
	})

	return images
}

func loadImage(filename string, fileType string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Open Error: ", err, " ", filename)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal("Decode Error: ", err, " ", filename)
	}

	return img
}

func getPixels(img image.Image) []pixel {
	bounds := img.Bounds()
	fmt.Println(bounds.Dx(), " x ", bounds.Dy())
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dy()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
		counter++
	}

	return pixels
}
