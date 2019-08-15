package smarp

import (
	"math"

	"github.com/gdamore/tcell"
)

func NewScreen() (*Screen, error) {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)

	s, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := s.Init(); err != nil {
		return nil, err
	}

	return &Screen{Screen: s}, nil
}

type Screen struct {
	Screen tcell.Screen
}

func (s *Screen) Set(p *Particle) {
	st := tcell.StyleDefault
	gl := p.Char()

	s.Screen.SetContent(round(p.C.X), round(p.C.Y), gl, nil, st)
}

func (s *Screen) PollEvent() tcell.Event {
	return s.Screen.PollEvent()
}

func (s *Screen) Show() {
	s.Screen.Show()
}

func (s *Screen) Sync() {
	s.Screen.Sync()
}

func (s *Screen) Size() (int, int) {
	return s.Screen.Size()
}

func (s *Screen) Clear() {
	s.Screen.Clear()
}

func (s *Screen) ClearResources() {
	s.Screen.Fini()
}

func round(f float64) int {
	return int(math.Round(f))
}
