package main

import (
	"fmt"
	"math"

	"github.com/prog-1/packages/areas"
)

func main() {
	fmt.Println("figure   : got, want")
	fmt.Println("--------------------")
	fmt.Println("circle   :", areas.Circle(2), 2*2*math.Pi)
	fmt.Println("square   :", areas.Square(2), 2*2)
	fmt.Println("rectangle:", areas.Rectangle(2, 3), 2*3)
}
