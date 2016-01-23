package main

import "fmt"
import "github.com/prosconi/gomnist/mnist"

func createPixels(numberOfCols, numberOfRows int32) [][]byte {
    pixels := make([][]byte, numberOfRows)
    for i := range pixels {
        pixels[i] = make([]byte, numberOfCols)
    }
    return pixels
}

func main() {
    imageFile := "t10k-images.idx3-ubyte"
    labelFile := "t10k-labels.idx1-ubyte"
    data := mnist.Open(imageFile, labelFile)
    
    label, image, _ := data.Next()
    
    fmt.Printf("%v - %v", label, image.Bounds())
}