package main

import (
	"fmt"
	"math"
)

func determineNewPosition(position Coordinates, heading Heading, distance int) (newPosition Coordinates) {
	switch heading {
	case N:
		newPosition.X = position.X
		newPosition.Y = position.Y - distance
	case E:
		newPosition.X = position.X + distance
		newPosition.Y = position.Y
	case S:
		newPosition.X = position.X
		newPosition.Y = position.Y + distance
	case W:
		newPosition.X = position.X - distance
		newPosition.Y = position.Y
	}

	return
}

func goRight(heading Heading, degrees int) (newHeading Heading) {
	switch heading {
	case N:
		switch degrees {
		case 90:
			newHeading = E
		case 180:
			newHeading = S
		case 270:
			newHeading = W
		}
	case E:
		switch degrees {
		case 90:
			newHeading = S
		case 180:
			newHeading = W
		case 270:
			newHeading = N
		}
	case S:
		switch degrees {
		case 90:
			newHeading = W
		case 180:
			newHeading = N
		case 270:
			newHeading = E
		}
	case W:
		switch degrees {
		case 90:
			newHeading = N
		case 180:
			newHeading = E
		case 270:
			newHeading = S
		}
	}

	return
}

func goLeft(heading Heading, degrees int) (newHeading Heading) {
	switch heading {
	case N:
		switch degrees {
		case 90:
			newHeading = W
		case 180:
			newHeading = S
		case 270:
			newHeading = E
		}
	case E:
		switch degrees {
		case 90:
			newHeading = N
		case 180:
			newHeading = W
		case 270:
			newHeading = S
		}
	case S:
		switch degrees {
		case 90:
			newHeading = E
		case 180:
			newHeading = N
		case 270:
			newHeading = W
		}
	case W:
		switch degrees {
		case 90:
			newHeading = S
		case 180:
			newHeading = E
		case 270:
			newHeading = N
		}
	}

	return
}

func determineNewHeading(heading Heading, command Command) (newHeading Heading) {
	if command.Direction == "L" {
		newHeading = goLeft(heading, command.Distance)
	} else {
		newHeading = goRight(heading, command.Distance)
	}

	return
}

func steer(position Coordinates, heading Heading, command Command) (newPosition Coordinates,
	newHeading Heading) {
	switch command.Direction {
	case "F":
		newPosition = determineNewPosition(position, heading, command.Distance)
		newHeading = heading
	case "R", "L":
		newPosition.X, newPosition.Y = position.X, position.Y
		newHeading = determineNewHeading(heading, command)
	case "N":
		newPosition.X, newPosition.Y = position.X, position.Y-command.Distance
		newHeading = heading
	case "E":
		newPosition.X, newPosition.Y = position.X+command.Distance, position.Y
		newHeading = heading
	case "S":
		newPosition.X, newPosition.Y = position.X, position.Y+command.Distance
		newHeading = heading
	case "W":
		newPosition.X, newPosition.Y = position.X-command.Distance, position.Y
		newHeading = heading
	}

	return
}

func part1(commands []Command) {
	position := Coordinates{X: 0, Y: 0}
	var heading = E

	for _, command := range commands {
		position, heading = steer(position, heading, command)
	}

	manhattan := math.Abs(float64(position.X)) + math.Abs(float64(position.Y))
	fmt.Println("Part 1 =", manhattan)
}

func main() {
	commands := readData()

	part1(commands)
	part2(commands)
}

// Local Variables:
// compile-command: "go build"
// End:
