package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"math"
	"os"
	"runtime"
	"sync"

	"github.com/disintegration/imaging"
)

type ImageHistogram struct {
	RGBHistogram [3][256]float64
}

const width int = 600 // new width ../img/22157905_1565518940176364_8232237062615465984_n.jpg
const height int = 400

var stringsSlice []string

func main() {

	var sliceLock sync.Mutex

	var wg sync.WaitGroup

	args := os.Args

	folderPath := args[2]

	fileNames, err := getFilesInFolder(folderPath)

	if err != nil {
		log.Fatalf("Error reading folder: %v", err)
	}

	srcImage, err := imaging.Open(args[1])
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	img1 := resize(srcImage)

	wg.Add(len(fileNames))

	for _, fileName := range fileNames {
		go aux(folderPath, fileName, img1, &wg, &sliceLock)
	}

	wg.Wait()

	for _, filename := range stringsSlice {
		fmt.Println(filename)
	}
}

func aux(folderPath, fileName string, img1 image.Image, wg *sync.WaitGroup, sliceLock *sync.Mutex) {
	defer wg.Done()
	srcImage2, err := imaging.Open(folderPath + "/" + fileName)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	img2 := resize(srcImage2)
	diff := HistogramIntersection(img1, img2)
	if diff <= 1.0 {
		sliceLock.Lock()
		stringsSlice = append(stringsSlice, folderPath+"/"+fileName)
		sliceLock.Unlock()
	}
}

func resize(srcImage image.Image) image.Image {

	// Reduce the dimensions by resizing the image
	resizedImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)

	return resizedImage
}

func HistogramIntersection(image1, image2 image.Image) float64 {
	histogram1 := CalculateHistogram(image1)
	histogram2 := CalculateHistogram(image2)

	var distance float64

	for i := 0; i < 3; i++ {
		for j := 0; j < 256; j++ {
			distance += math.Abs(float64(histogram1.RGBHistogram[i][j] - histogram2.RGBHistogram[i][j]))
		}
	}

	return distance
}

func CalculateHistogram(img image.Image) *ImageHistogram {
	histogram := &ImageHistogram{}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r, g, b = r/256, g/256, b/256

			histogram.RGBHistogram[0][r]++
			histogram.RGBHistogram[1][g]++
			histogram.RGBHistogram[2][b]++
		}
	}

	totalPixels := float64(width * height)

	for c := 0; c < 3; c++ {
		for i := 0; i < 256; i++ {
			histogram.RGBHistogram[c][i] /= totalPixels
		}
	}

	return histogram
}

func getFilesInFolder(folderPath string) ([]string, error) {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}
