package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	// 	b, _ := ioutil.ReadFile("./stock_image.jpg")
	// 	fmt.Println(b)
	file, _ := os.Open("./stock_image.jpg")
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(img.Bounds())
}
