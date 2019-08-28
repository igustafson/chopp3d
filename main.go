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

	imgWithMessage := HideMessage("this is a secret message", img)

	err = writeFile(imgWithMessage)
	if err != nil {
		fmt.Println(err)
	}

	imgWithMessage, err = readImage("./output.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}

	message := ReadMessage(imgWithMessage)

	fmt.Println(message)

	return
}
