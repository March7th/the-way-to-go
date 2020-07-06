package main

import (
	"fmt"
	"image/gif"
	"log"
)

func main() {
	anim := gif.GIF{LoopCount: 64}
	fmt.Printf("type of anim is %T\n", anim)

	fmt.Print("just for test\t")
	fmt.Println("just for test")

	log.Print("**log test1**")
	log.Print("**log test2**\n")
	log.Print("**log test3**\n\n")
	log.Println("**log test4**")

	log.Fatal("**Fatal test**")
}
