package main

import "image"

//Recognizers will recognize a digit from an input image, then return that digit.
type Recognizer interface {
    Recognize(image.Image) byte 
}