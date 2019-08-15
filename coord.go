package smarp

import (
	"math"
)

type Coord struct {
	X, Y float64
}

func (a *Coord) Distance(o *Coord) float64 {
	x := math.Abs(o.X - a.X)
	y := math.Abs(o.Y - a.Y)

	return math.Sqrt(x*x + y*y)

}

func (a *Coord) ApplyVector(v *Vector) {
	a.X += v.C.X
	a.Y += v.C.Y
}

func (a *Coord) Plus(o *Coord) *Coord {
	ox := a.X + o.X
	oy := a.Y + o.Y

	return &Coord{X: ox, Y: oy}
}

func (a *Coord) Minus(o *Coord) *Coord {
	ox := a.X - o.X
	oy := a.Y - o.Y

	return &Coord{X: ox, Y: oy}
}

func AdjustX(x, maxx float64) float64 {
	if x < 0 {
		return maxx
	}

	if x > maxx {
		return 0
	}

	return x
}

func (a *Coord) Adjust(maxx, maxy float64) {
	a.X = AdjustX(a.X, maxx)
	a.Y = AdjustX(a.Y, maxy)
}
