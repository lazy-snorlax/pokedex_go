package main

import (
	"fmt"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func generateImg(imgUrl string) {
	flags := aic_package.DefaultFlags()
	flags.Colored = true
	flags.Dimensions = []int{50, 20}

	// Conversion for an image
	asciiArt, err := aic_package.Convert(imgUrl, flags)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", asciiArt)
}
