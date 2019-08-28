package main

import (
	"bytes"
	"fmt"
<<<<<<< HEAD
)

func main() {
	img, err := readImage("./stock_image.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Add a secret message

	err = writeFile(img)
	if err != nil {
		fmt.Println(err)
	}

	return
=======
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
>>>>>>> e72efdbd3465754daef68fb1d9cfc3da9fb1b1a3
}
