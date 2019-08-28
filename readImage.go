package main

import (
	"bytes"
	"image"
	"io/ioutil"
)

func readImage(path string) (image.Image, error) {
	b, _ := ioutil.ReadFile(path)
	img, _, err := image.Decode(bytes.NewReader(b))

	return img, err
}
