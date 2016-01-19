package main

import "bytes"
import "fmt"
import "os"
import "encoding/binary"

func readInt32(buffer *os.File) (int32, error) {
    var i int32
    intBuffer := make([]byte, 4)
    buffer.Read(intBuffer)
    reader := bytes.NewReader(intBuffer)
    err := binary.Read(reader, binary.BigEndian, &i)
    if err != nil {
        panic(err)
    }
    return i, err
}

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
    stream, err := os.Open(imageFile)
    if err != nil {
        panic(err)
    }
    defer func() { stream.Close() }()

    magicNumber, _      := readInt32(stream)
    numberOfImages, _   := readInt32(stream)
    numberOfRows, _     := readInt32(stream)
    numberOfCols, _     := readInt32(stream)

    fmt.Println("Image file: " + imageFile)
    fmt.Println("Label file: " + labelFile)
    fmt.Printf("Magic number: %d", magicNumber)
    fmt.Println()
    fmt.Printf("Number of images: %d", numberOfImages)
    fmt.Println()
    fmt.Printf("Number of rows:  %d", numberOfRows)
    fmt.Println()
    fmt.Printf("Number of cols:  %d", numberOfCols)
    fmt.Println()
        
    singleByteBuffer := make([]byte, 1)
    count := 0
    for i := int32(0); i < numberOfImages; i++ {
        pixels := createPixels(numberOfCols, numberOfRows)
        for x := int32(0); x < numberOfRows; x++ {
            for y := int32(0); y < numberOfRows; y++ {
                stream.Read(singleByteBuffer)
                pixels[x][y] = singleByteBuffer[0]
                count++
            }
        }
    }
    fmt.Printf("Total number of pixels: %d", count)
}