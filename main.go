package main

import (
	"fmt"
	"github.com/go-wms/tidy/image"
)

func main() {
	imageUrls := []string{"http://ae01.alicdn.com/kf/HTB1tEHzb.gQMeJjy0Ff762ddXXaF.png"}
	a, err := image.Reduce(imageUrls)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)
	//image.ReduceImage("tmp/HTB1qVq4FxWYBuNjy1zkq6xGGpXaD.jpg")
}
