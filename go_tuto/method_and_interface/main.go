package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method
// (v Vertex)をレシーバという
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 副作用をもたせたいときは、*Vertexとする
// レシーバ型はコピーに対して、実行されるため、*をつけないと副作用がなくなる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func main() {
	v := Vertex{3, 4}
	v2 := v.Scale2(10)
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v2.Abs())
}
