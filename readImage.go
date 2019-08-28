package main

import (
	"io/ioutil"
)

func readImage(path string) ([]byte, error) {
	b, _ := ioutil.ReadFile(path)
	// img, _, err := image.Decode(bytes.NewReader(b))

	return b, nil
}
