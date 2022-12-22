package main

import (
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"

	"image/draw"
	"image/png"

	"github.com/google/uuid"
)

type Layer struct {
	AssetsFolder string
	Position     image.Point
	NextLayer    *Layer
}

func main() {
	quotesLayer := Layer{
		AssetsFolder: "./quotes",
		Position:     image.Point{668, 100},
	}

	gophersLayer := Layer{
		AssetsFolder: "./gophers",
		Position:     image.Point{256, 256},
		NextLayer:    &quotesLayer,
	}

	backgroundsLayer := Layer{
		AssetsFolder: "./backgrounds",
		Position:     image.Point{0, 0},
		NextLayer:    &gophersLayer,
	}

	baseImage := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	collection, err := addLayer([]*image.RGBA{baseImage}, &backgroundsLayer)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, img := range collection {
		out, err := os.Create(fmt.Sprintf("./collection/%s.png", uuid.NewString()))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		if err := png.Encode(out, img); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		out.Close()

	}

	fmt.Println("Done.")
}

func addLayer(prevImages []*image.RGBA, layer *Layer) ([]*image.RGBA, error) {
	if layer == nil {
		return prevImages, nil
	}

	layerImages := []image.Image{}
	err := filepath.Walk(layer.AssetsFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, fileErr := os.Open(path)
		if fileErr != nil {
			return fileErr
		}

		defer file.Close()

		img, _, decodeErr := image.Decode(file)
		if decodeErr != nil {
			return decodeErr
		}

		layerImages = append(layerImages, img)

		return nil
	})
	if err != nil {
		return prevImages, err
	}

	newImages := []*image.RGBA{}

	for _, prevImage := range prevImages {
		for _, layerImage := range layerImages {
			dst := image.NewRGBA(prevImage.Bounds())

			// blend layers
			draw.Draw(dst, prevImage.Bounds(), prevImage, image.Point{}, draw.Over)
			draw.Draw(dst, layerImage.Bounds().Add(layer.Position), layerImage, image.Point{}, draw.Over)

			newImages = append(newImages, dst)
		}
	}

	return addLayer(newImages, layer.NextLayer)
}
