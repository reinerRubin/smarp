package smarp

import (
	"math"
)

type Coord struct {
	X, Y float64
}

func (a *Coord) Angle(o *Coord) float64 {
	x := o.X - a.X
	y := o.Y - a.Y

	tgk := y / x
	angle := math.Atan(tgk)

	if x < 0 && y < 0 {
		angle = math.Pi + angle
	}
	if x < 0 && y > 0 {
		angle = -angle
		angle = math.Pi/2 + (math.Pi/2 - angle)
	}
	if y < 0 && x > 0 {
		angle = -angle
		angle = math.Pi + math.Pi/2 + (math.Pi/2 - angle)
	}

	return angle
}

func (a *Coord) Distance(o *Coord) float64 {
	x := math.Abs(o.X - a.X)
	y := math.Abs(o.Y - a.Y)

	return math.Sqrt(x*x + y + y)

}

func (a *Coord) ApplyVector(v *Vector) {
	a.X += math.Cos(v.Angle) * v.Value
	a.Y += math.Sin(v.Angle) * v.Value
}

func (a *Coord) Adjust(maxx, maxy float64) {
	a.X = AdjustX(a.X, maxx)
	a.Y = AdjustX(a.Y, maxy)
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
