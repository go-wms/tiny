package image

import (
	"fmt"
	"github.com/go-wms/tidy/file"
	"github.com/nfnt/resize"
	"golang.org/x/sync/errgroup"
	"image"
	"image/jpeg"
	"os"
)

//Reduce Batch compression of pictures
func Reduce(imageUrls []string) (paths []string, err error) {
	group := new(errgroup.Group)
	for _, imageUrl := range imageUrls {
		group.Go(func() error {
			// 保存图片，获取图片路径
			filePath,err := file.Save(imageUrl)
			if err != nil {
				return err
			}
			// 打开图片，并压缩，返回文件路径
			newFilePath,err := ReduceImage(filePath)
			if err != nil {
				return err
			}
			fmt.Println(newFilePath)
			paths = append(paths, newFilePath)
			return nil
		})
	}
	if err := group.Wait(); err!=nil {
		return nil, err
	}
	return
}

func ReduceImage(filePath string) (string,error) {
	openFile, err := os.Open("tmp/"+filePath)
	fmt.Println(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer func() {
		_ = openFile.Close()
	}()

	img, _, err := image.Decode(openFile)
	if err != nil {
		return "", err
	}
	width := img.Bounds().Dx()

	m := resize.Resize(uint(width), 0, img, resize.Lanczos3)
	newFilePath := "resize/" + filePath
	fmt.Println(newFilePath)
	out, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = out.Close()
	}()

	// write new image to open
	err = jpeg.Encode(out, m, nil)
	return newFilePath, err
}
