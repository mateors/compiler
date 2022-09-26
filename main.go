package main

import (
	"fmt"
	"math"

	"github.com/mateors/compiler/code"
)

func main() {

	fmt.Println([]int{65534})
	fmt.Println([]byte{byte(0), 255, 254})
	fmt.Println(math.Pow(2, 16)) //2--16

	instruction := code.Make(0, 65534)
	fmt.Println(instruction) //[0 255 254]
}
