// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// boxes just displays random colored boxes on your terminal screen.
// Press ESC to exit the program.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/reinerRubin/smarp"
)

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

	logFile, err := os.Create("/tmp/log")
	if err != nil {
		os.Exit(1)
	}
	log.SetOutput(logFile)
	defer logFile.Close()

	e := &smarp.Engine{
		Maxx: func() float64 { x, _ := screen.Size(); return float64(x) },
		Maxy: func() float64 { _, y := screen.Size(); return float64(y) },
		Particles: smarp.Particles{
			&smarp.Particle{
				C:          &smarp.Coord{X: 20, Y: 20},
				Speed:      &smarp.Vector{},
				CharToShow: '1',
				Mass:       0.03,
			},
			&smarp.Particle{
				C:          &smarp.Coord{X: 40, Y: 16},
				Speed:      &smarp.Vector{Angle: 0, Value: 0.5},
				CharToShow: '2',
				Mass:       0.008,
			},
			// &smarp.Particle{
			//	C:          &smarp.Coord{X: 30, Y: 16},
			//	Speed:      &smarp.Vector{Angle: math.Pi, Value: 0.5},
			//	CharToShow: '3',
			// },
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
		case <-time.After(time.Millisecond * 10):
		}
		start := time.Now()
		draw(screen, e)
		e.Tick()

		cnt++
		dur += time.Now().Sub(start)
	}
}
