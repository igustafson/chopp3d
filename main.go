package main

import (
	"fmt"
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
}
