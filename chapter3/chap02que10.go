package main

import (
	"fmt"
)

type space3D interface {
	xaxis() int
	yaxis() int
	zaxis() int
}

type points struct {
	X, Y, Z int
}

func (p points) xaxis() int {
	return p.X
}

func (p points) yaxis() int {
	return p.Y
}

func (p points) zaxis() int {
	return p.Z
}

func findcoorinates(s space3D) {
	fmt.Println("X:", s.xaxis(), "Y:", s.yaxis(), "Z:", s.zaxis())
}

type xcoordinate int

func (p xcoordinate) xaxis() int {
	return int(p)
}

func (p xcoordinate) yaxis() int {
	return 0
}

func (p xcoordinate) zaxis() int {
	return 0
}

func main() {

	a := points{X: 1, Y: 2, Z: 3}
	fmt.Println(a)
	findcoorinates(a)
	b := xcoordinate(4)
	findcoorinates(b)
}