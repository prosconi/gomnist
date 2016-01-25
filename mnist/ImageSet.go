package mnist

import "fmt"
import "os"
import "bytes"
import "encoding/binary"

type imageSet struct {
    imageFile, labelFile *os.File
    count, rows, cols int32
}

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

func Open(imageFilename, labelFilename string) imageSet {
    var imageset imageSet
    var magic, countImg, countLbl int32
    imageset.imageFile,_ = os.Open(imageFilename)
    magic,_ = readInt32(imageset.imageFile)
    if magic != 2051 { panic(fmt.Sprintf("oh shit: %v isn't 2051", magic)) }
    
    imageset.labelFile,_ = os.Open(labelFilename)
    magic,_ = readInt32(imageset.labelFile)
    if magic != 2049 { panic(fmt.Sprintf("oh shit: %v isn't 2049", magic)) }
    
    countImg,_ = readInt32(imageset.imageFile);
    countLbl,_ = readInt32(imageset.labelFile);
    if countImg != countLbl {
        panic(fmt.Sprintf("oh shit: %v images and %v labels", countImg, countLbl))
    } else {
        imageset.count = countImg
    }
    
    imageset.rows,_ = readInt32(imageset.imageFile);
    imageset.cols,_ = readInt32(imageset.imageFile);
    
    return imageset
}

func (is *imageSet) Next() (byte, mnistImage, error){
    nextLabel := make([]byte, 1)
    nextImage := mnistImage {
        pixels: make([]byte, is.rows*is.cols),
        w: is.cols,
        h: is.rows,
    }
    
    is.labelFile.Read(nextLabel)
    _,imgErr := is.imageFile.Read(nextImage.pixels)  
    
    return nextLabel[0], nextImage, imgErr
}