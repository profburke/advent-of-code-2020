package main

import "fmt"

type Coordinates struct {
	X int
	Y int
	Z int
}

type Grid map[Coordinates]bool

type CA3 struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
	G      Grid
}

func (ca *CA3) NeighborCount(cell Coordinates) (count int) {
	count = 0

	for x := cell.X - 1; x <= cell.X+1; x++ {
		for y := cell.Y - 1; y <= cell.Y+1; y++ {
			for z := cell.Z - 1; z <= cell.Z+1; z++ {
				if x == cell.X && y == cell.Y && z == cell.Z {
					continue
				}

				c := Coordinates{X: x, Y: y, Z: z}
				if _, found := ca.G[c]; found {
					count++
				}
			}
		}
	}

	return
}

func (ca *CA3) Step() {
	newGrid := make(Grid)

	for x := ca.X1 - 1; x <= ca.X2+1; x++ {
		for y := ca.Y1 - 1; y <= ca.Y2+1; y++ {
			for z := ca.Z1 - 1; z <= ca.Z2+1; z++ {
				c := Coordinates{X: x, Y: y, Z: z}
				_, active := ca.G[c]
				count := ca.NeighborCount(c)

				if active && (count == 2 || count == 3) {
					newGrid[c] = true
				}

				if !active && (count == 3) {
					newGrid[c] = true
				}
			}
		}
	}

	ca.G = newGrid

	// strictly speaking, updating the bounds is only necessary
	// if there is a newly active cell on the border...
	// but it doesn't "hurt" to do this, and it's simpler
	ca.X1--
	ca.Y1--
	ca.Z1--
	ca.X2++
	ca.Y2++
	ca.Z2++
}

func (ca *CA3) ActiveCount() (count int) {
	for x := ca.X1; x <= ca.X2; x++ {
		for y := ca.Y1; y <= ca.Y2; y++ {
			for z := ca.Z1; z <= ca.Z2; z++ {
				c := Coordinates{X: x, Y: y, Z: z}
				if _, found := ca.G[c]; found {
					count++
				}
			}
		}
	}

	return
}

func (ca *CA3) String() (result string) {
	result = ""

	for z := ca.Z1; z <= ca.Z2; z++ {
		result += fmt.Sprintf("z=%d\n", z)
		for x := ca.X1; x <= ca.X2; x++ {
			for y := ca.Y1; y <= ca.Y2; y++ {
				c := Coordinates{X: x, Y: y, Z: z}
				if _, active := ca.G[c]; active {
					result += "#"
				} else {
					result += "."
				}
			}
			result += "\n"
		}
		result += "\n\n"
	}

	return
}

// Local Variables:
// compile-command: "go build"
// End:
