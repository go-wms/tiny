package main

import (
	"fmt"
	"github.com/go-wms/tidy/image"
)

func main() {
	imageUrls := []string{"https://ae01.alicdn.com/kf/HTB1qVq4FxWYBuNjy1zkq6xGGpXaD.jpg"}
	a,_ := image.Reduce(imageUrls)
	fmt.Println(a)
	//image.ReduceImage("tmp/HTB1qVq4FxWYBuNjy1zkq6xGGpXaD.jpg")
}
