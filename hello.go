package main

import "fmt"
import "github.com/prosconi/gomnist/guesser"
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
    
    recognizer := guesser.Guesser()

    var correct, trials int
    
    for label, image, err := data.Next(); err == nil ; label, image, err = data.Next() {
        trials++
        guess := recognizer.Recognize(image)
        
        if guess == label { correct++ }
    }
    
    fmt.Printf("Guessed %v correct of %v: %v%%", correct, trials, float32(correct)*100/float32(trials))
}