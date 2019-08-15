package smarp

import (
	"fmt"
	"testing"
)

func TestAngle(t *testing.T) {
	// c1 := &Coord{X: 0, Y: 0}
	// fmt.Println(c1.Angle(&Coord{X: 2, Y: 2}))   // 45
	// fmt.Println(c1.Angle(&Coord{X: 0, Y: 2}))   // 90
	// fmt.Println(c1.Angle(&Coord{X: -2, Y: -2})) // 225
	// fmt.Println(c1.Angle(&Coord{X: -2, Y: 2}))  // 135
	// fmt.Println(c1.Angle(&Coord{X: 2, Y: -2}))  // 315

	{
		c1 := &Coord{X: 2, Y: 2}
		c2 := &Coord{X: -2, Y: -2}
		fmt.Println(c1.Angle(c2)) // 135
		fmt.Println(c2.Angle(c1)) // 135
	}
}
