package pdf2picture

import (
	"os"

	"github.com/h2non/bimg"
)

func Something() {
	image, imageErr := bimg.NewImage([]byte{}).Convert(bimg.PNG)
	if imageErr != nil {
		panic(imageErr)
	}

	file, _ := os.Create("./pic.png")
	defer file.Close()

	file.Write(image)
}
