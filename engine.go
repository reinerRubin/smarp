package smarp

const Gravity = 1
const SpeedLimit = 2

type (
	Engine struct {
		Maxx, Maxy func() float64

		Particles Particles
	}
)

func (e *Engine) Tick() {
	// recalculate forces
	for _, tPartical := range e.Particles {
		forces := make(Vectors, 0)

		for _, oPartical := range e.Particles {
			if tPartical == oPartical {
				continue
			}

			f1, _ := ForcesBetween(tPartical, oPartical)
			forces = append(forces, f1)
		}

		ApplyForceToParticale(tPartical, forces.Sum())
	}

	// change positions
	for _, tPartical := range e.Particles {
		tPartical.Tick()
		tPartical.C.Adjust(e.Maxx(), e.Maxy())
	}
}

func ForcesBetween(p1, p2 *Particle) (f1, f2 *Vector) {
	distance := p1.C.Distance(p2.C)
	force := Gravity * float64(p1.GetMass()*p2.GetMass()) / distance * distance

	v1 := &Vector{C: p2.C.Minus(p1.C)}
	v2 := &Vector{C: p1.C.Minus(p2.C)}

	return v1.SetValue(force),
		v2.SetValue(force)
}

func ApplyForceToParticale(p *Particle, f *Vector) {
	a := f.Value() / p.GetMass()
	acceleration := f.SetValue(a)
	p.Speed = p.Speed.Plus(acceleration)
}
