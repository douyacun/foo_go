package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

// https://github.com/milosgajdos/gocv-playground/blob/master/03_Colors_Thresholding/README.md
func main() {
	out := "/Users/admin/Desktop/gray.jpeg"
	src := gocv.IMRead("/Users/admin/Desktop/车牌2.jpeg", gocv.IMReadGrayScale)
	if src.Empty() {
		fmt.Println("不能读取图片...")
		return
	}

	gray := gocv.NewMat()
	defer gray.Close()

	bin := gocv.NewMat()
	defer bin.Close()

	inv := gocv.NewMat()
	defer inv.Close()

	//gocv.CvtColor(src, &gray, gocv.ColorBGRToGray)
	//gocv.Threshold(gray, &bin, 10.0, 255.0, gocv.ThresholdBinary)
	//gocv.Threshold(gray, &inv, 10.0, 255.0, gocv.ThresholdBinaryInv)
	gocv.AdaptiveThreshold(src, &bin, 255.0, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 5, 4.0)
	if ok := gocv.IMWrite("/Users/admin/Desktop/gray.jpeg", bin); !ok {
		fmt.Printf("Failed to write image: %s\n", out)
		return
	}

	//gocv.GaussianBlur(bin, &inv, image.Point{X: 3, Y: 3}, 0.8, 0.8, gocv.BorderDefault)
	gocv.MedianBlur(bin, &inv, 3)
	//gocv.IMWrite("/Users/admin/Desktop/bin.jpeg", bin)
	gocv.IMWrite("/Users/admin/Desktop/inv.jpeg", inv)
}
