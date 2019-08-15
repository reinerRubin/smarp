package smarp

import "math"

type (
	Vector struct {
		Angle float64
		Value float64
	}

	Vectors []*Vector
)

func (vv Vectors) Sum() *Vector {
	rv := &Vector{}

	for _, v := range vv {
		rv = rv.Plus(v)
	}

	return rv
}

func (a *Vector) Plus(o *Vector) *Vector {
	ox := math.Cos(o.Angle)*o.Value + math.Cos(a.Angle)*a.Value
	oy := math.Sin(o.Angle)*o.Value + math.Sin(a.Angle)*a.Value

	zero := &Coord{}
	newAngle := zero.Angle(&Coord{X: ox, Y: oy})
	return &Vector{
		Value: math.Sqrt(ox*ox + oy*oy),
		Angle: newAngle,
	}
}
