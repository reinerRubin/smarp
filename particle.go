package smarp

type (
	Particles []*Particle
	Particle  struct {
		C     *Coord
		Speed *Vector
		Mass  float64

		CharToShow rune
	}
)

func (p *Particle) GetMass() float64 {
	return p.Mass
}

func (p *Particle) Char() rune {
	return p.CharToShow
}

func (p *Particle) Tick() {
	p.C.ApplyVector(p.Speed)
}
