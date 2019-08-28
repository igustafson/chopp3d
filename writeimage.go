package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
)

func writeFile(b []byte) (err error) {
	buf := new(bytes.Buffer)
	image := image.

	err = jpeg.Encode(buf, i, nil)
	if err != nil {
		return
	}

	return ioutil.WriteFile("output.jpeg", buf.Bytes(), 0777)
}