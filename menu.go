package main

type MenuItem struct {
	Label  string
	Target Screen
	X, Y   int
	W, H   int
}

func (m MenuItem) Contains(px, py int) bool {
	return px >= m.X &&
		px <= m.X+m.W &&
		py >= m.Y &&
		py <= m.Y+m.H
}
