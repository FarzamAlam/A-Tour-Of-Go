package main

import (
	"fmt"
	"math"
)

type MyFloat float64

type Abser interface{
	Abs()float64
}

type Vertex struct {
	X, Y float64
}

