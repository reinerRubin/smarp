package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/reinerRubin/smarp"
)

const tickTime = time.Millisecond * 10

func draw(s *smarp.Screen, e *smarp.Engine) {
	for _, p := range e.Particles {
		s.Set(p)
	}

	s.Show()
}

func main() {
	screen, err := smarp.NewScreen()
	if err != nil {
		fmt.Printf("cant init screen: %s, err")
		os.Exit(1)
	}
	defer screen.ClearResources()
	screen.Clear()

	// logFile, err := os.Create("/tmp/log")
	// if err != nil {
	//	os.Exit(1)
	// }
	// log.SetOutput(logFile)
	// defer logFile.Close()

	e := &smarp.Engine{
		Maxx: func() float64 { x, _ := screen.Size(); return float64(x) },
		Maxy: func() float64 { _, y := screen.Size(); return float64(y) },
		Particles: smarp.Particles{
			&smarp.Particle{
				C:          &smarp.Coord{X: 30, Y: 20},
				Speed:      &smarp.Vector{C: &smarp.Coord{}},
				CharToShow: '1',
				Mass:       0.01,
			},
			&smarp.Particle{
				C:          &smarp.Coord{X: 22, Y: 32},
				Speed:      &smarp.Vector{C: &smarp.Coord{}},
				CharToShow: '4',
				Mass:       0.03,
			},
			&smarp.Particle{
				C:          &smarp.Coord{X: 40, Y: 16},
				Speed:      smarp.NewVectorForHuman(0, 0.3),
				CharToShow: '2',
				Mass:       0.02,
			},
			&smarp.Particle{
				C:          &smarp.Coord{X: 30, Y: 26},
				Speed:      smarp.NewVectorForHuman(math.Pi, 0.2),
				CharToShow: '3',
				Mass:       0.008,
			},
		},
	}

	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	cnt := 0
	dur := time.Duration(0)
loop:
	for {
		screen.Clear()
		select {
		case <-quit:
			break loop
		case <-time.After(tickTime):
		}
		start := time.Now()
		draw(screen, e)
		e.Tick()

		cnt++
		dur += time.Now().Sub(start)
	}
}
