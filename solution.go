package main

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"fmt"
	"math"
)

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func main() {
	fmt.Println("Area of shapes")
	CalcSquare(8, SidesTriangle)
	CalcSquare(10, SidesSquare)
	CalcSquare(12, SidesCircle)
}

func CalcSquare(sideLen float64, sidesNum uint64) float64 {

	var result float64
	switch sidesNum {
	case SidesTriangle:
		result = math.Sqrt(3.0) / 4.0 * math.Pow(sideLen, 2)
		fmt.Println("The area of a triangle is equal to:", result)
	case SidesSquare:
		result = sideLen * sideLen
		fmt.Println("The area of a square is equal to:", result)
	case SidesCircle:
		result = math.Pi * sideLen * sideLen
		fmt.Println("The area of a circle is equal to:", result)
	default:
		fmt.Println("This value is not supported")
	}

	return result
}
