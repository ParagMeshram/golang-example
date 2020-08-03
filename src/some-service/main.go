package main

import (
	"fmt"
	"some-common/common"
	"some-common/goodbye"
)

func main() {
	fmt.Println("Hey there")
	fmt.Println(goodbye.Goodbye())
	fmt.Println(common.Square(12))
}
