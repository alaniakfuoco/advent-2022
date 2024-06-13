package day17

import "advent/util"

type shape struct {
	bots     []util.Point
	other    []util.Point
	farLeft  int
	farRight int
	tippyTop int
}

func newFlat(y int) *shape {
	left := util.Point{X: 3, Y: y}
	mid1 := util.Point{X: 4, Y: y}
	mid2 := util.Point{X: 5, Y: y}
	right := util.Point{X: 6, Y: y}
	return &shape{
		bots:     []util.Point{left, mid1, mid2, right},
		farLeft:  3,
		farRight: 6,
		tippyTop: y,
	}
}

func newPlus(y int) *shape {
	left := util.Point{X: 3, Y: y + 1}
	mid := util.Point{X: 4, Y: y + 1}
	right := util.Point{X: 5, Y: y + 1}
	top := util.Point{X: 4, Y: y + 2}
	bot := util.Point{X: 4, Y: y}
	return &shape{
		bots:     []util.Point{left, right, bot},
		other:    []util.Point{mid, top},
		farLeft:  3,
		farRight: 5,
		tippyTop: y + 2,
	}
}

func newTall(y int) *shape {
	top := util.Point{X: 3, Y: y + 3}
	mid1 := util.Point{X: 3, Y: y + 2}
	mid2 := util.Point{X: 3, Y: y + 1}
	bot := util.Point{X: 3, Y: y}
	return &shape{
		bots:     []util.Point{bot},
		other:    []util.Point{top, mid1, mid2},
		farLeft:  3,
		farRight: 3,
		tippyTop: y + 3,
	}
}

func newL(y int) *shape {
	top := util.Point{X: 5, Y: y + 2}
	mid1 := util.Point{X: 5, Y: y + 1}
	corner := util.Point{X: 5, Y: y}
	mid2 := util.Point{X: 4, Y: y}
	bot := util.Point{X: 3, Y: y}
	return &shape{
		bots:     []util.Point{bot, mid2, corner},
		other:    []util.Point{top, mid1},
		farLeft:  3,
		farRight: 5,
		tippyTop: y + 2,
	}
}

func newSquare(y int) *shape {
	tr := util.Point{X: 4, Y: y + 1}
	tl := util.Point{X: 3, Y: y + 1}
	br := util.Point{X: 4, Y: y}
	bl := util.Point{X: 3, Y: y}
	return &shape{
		bots:     []util.Point{bl, br},
		other:    []util.Point{tl, tr},
		farLeft:  3,
		farRight: 4,
		tippyTop: y + 1,
	}
}

func (s *shape) moveDown() {
	for i := range s.bots {
		s.bots[i].Y -= 1
	}
	for i := range s.other {
		s.other[i].Y -= 1
	}
	s.tippyTop -= 1
}

func (s *shape) moveLeft() {
	if s.farLeft > 1 {
		for i := range s.bots {
			s.bots[i].X -= 1
		}
		for i := range s.other {
			s.other[i].X -= 1
		}
		s.farLeft -= 1
		s.farRight -= 1
	}

}

func (s *shape) moveRight() {
	if s.farRight < 7 {
		for i := range s.bots {
			s.bots[i].X += 1
		}
		for i := range s.other {
			s.other[i].X += 1
		}
		s.farRight += 1
		s.farLeft += 1
	}
}

func (s *shape) getAllPoints() []util.Point {
	return append(s.bots, s.other...)
}
