package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("./stock_image.jpg")
	fmt.Println(b)
	img, _, _ := image.Decode(bytes.NewReader(b))

	fmt.Println(img.Bounds())
}
