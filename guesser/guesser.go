package guesser

import "image"
import "math/rand"
import "time"

//guesser is a Recognizer that guesses a random number between 0 and 9.
type guesser struct { r *rand.Rand }

func Guesser() guesser {
    s := rand.NewSource(time.Now().UnixNano())
    r := rand.New(s)
    return guesser { r }
}

func (g guesser) Recognize(i image.Image) byte {
    return byte(g.r.Int()%10)
}