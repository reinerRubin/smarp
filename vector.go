package smarp

import "math"

type (
	Vector struct {
		C *Coord
	}

	Vectors []*Vector
)

func NewVectorForHuman(angle, value float64) *Vector {
	newX := value * math.Cos(angle)
	newY := value * math.Sin(angle)

	return &Vector{C: &Coord{X: newX, Y: newY}}
}

func (vv Vectors) Sum() *Vector {
	rv := &Vector{C: &Coord{}}

	for _, v := range vv {
		rv = rv.Plus(v)
	}

	return rv
}

func (v *Vector) Value() float64 {
	return math.Sqrt(v.C.X*v.C.X + v.C.Y*v.C.Y)
}

func (v *Vector) SetValue(newValue float64) *Vector {
	angle := v.Angle()

	newX := newValue * math.Cos(angle)
	newY := newValue * math.Sin(angle)

	return &Vector{C: &Coord{X: newX, Y: newY}}
}

func (v *Vector) Plus(o *Vector) *Vector {
	return &Vector{
		C: v.C.Plus(o.C),
	}
}

func (v *Vector) Minus(o *Vector) *Vector {
	return &Vector{
		C: v.C.Minus(o.C),
	}
}

func (v *Vector) Angle() float64 {
	tgk := v.C.Y / v.C.X
	angle := math.Atan(tgk)

	if v.C.X < 0 && v.C.Y < 0 {
		angle = math.Pi + angle
	}
	if v.C.X < 0 && v.C.Y > 0 {
		angle = -angle
		angle = math.Pi/2 + (math.Pi/2 - angle)
	}
	if v.C.Y < 0 && v.C.X > 0 {
		angle = -angle
		angle = math.Pi + math.Pi/2 + (math.Pi/2 - angle)
	}
	return angle
}
