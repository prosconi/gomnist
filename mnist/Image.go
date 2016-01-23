package mnist

import "image"
import "image/color"

type mnistImage struct {
    pixels []byte
    w, h int32
}

func (i *mnistImage) ColorModel() color.Model {
    return color.GrayModel
}

func (i *mnistImage) Bounds() image.Rectangle {
    return image.Rect(0, 0, int(i.w), int(i.h))
}

func (i *mnistImage) At(x, y int) color.Color {
    return color.Gray { i.pixels[x + int(i.w)*y] } 
}