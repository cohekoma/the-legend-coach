package main

type MenuItem struct {
	Label  string
	Target Screen
	X, Y   int
	W, H   int
}

// check if the mouse cursor is within the current menu item (meaning player is click the item)
func (item MenuItem) Contains(px, py int) bool {
	return px >= item.X &&
		px <= item.X+item.W &&
		py >= item.Y &&
		py <= item.Y+item.H
}
