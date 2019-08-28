package main

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
)

func writeFile(i image.Image) (err error) {
	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, i, nil)
	if err != nil {
		return
	}

	return ioutil.WriteFile("output.jpeg", buf.Bytes(), 0777)
}