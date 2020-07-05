package main

import (
	"fmt"
	"image/gif"
)

func main() {
	anim := gif.GIF{LoopCount: 64}
	fmt.Printf("type of anim is %T\n", anim)
}
