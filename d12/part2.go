package main

import (
	"fmt"
	"math"
)

func rotateWaypoint(waypoint Coordinates, command Command) (newWaypoint Coordinates) {
	var s, c int
	switch command.Distance {
	case 90:
		s = 1
		c = 0
	case 180:
		s = 0
		c = -1
	case 270:
		s = -1
		c = 0
	}

	if command.Direction == "L" {
		s *= -1
	}

	newWaypoint.X = waypoint.X*c - waypoint.Y*s
	newWaypoint.Y = waypoint.X*s + waypoint.Y*c

	return
}

func steerByWaypoint(position, waypoint Coordinates, command Command) (newPosition,
	newWaypoint Coordinates) {
	switch command.Direction {
	case "F":
		newPosition.X = position.X + command.Distance*waypoint.X
		newPosition.Y = position.Y + command.Distance*waypoint.Y
		newWaypoint = waypoint
	case "R", "L":
		newWaypoint = rotateWaypoint(waypoint, command)
		newPosition = position
	case "N":
		newWaypoint.X, newWaypoint.Y = waypoint.X, waypoint.Y-command.Distance
		newPosition = position
	case "E":
		newWaypoint.X, newWaypoint.Y = waypoint.X+command.Distance, waypoint.Y
		newPosition = position
	case "S":
		newWaypoint.X, newWaypoint.Y = waypoint.X, waypoint.Y+command.Distance
		newPosition = position
	case "W":
		newWaypoint.X, newWaypoint.Y = waypoint.X-command.Distance, waypoint.Y
		newPosition = position
	}

	return
}

func part2(commands []Command) {
	position := Coordinates{X: 0, Y: 0}
	waypoint := Coordinates{X: 10, Y: -1}

	for _, command := range commands {
		position, waypoint = steerByWaypoint(position, waypoint, command)
	}

	manhattan := math.Abs(float64(position.X)) + math.Abs(float64(position.Y))
	fmt.Println("Part 2 =", manhattan)
}

// Local Variables:
// compile-command: "go build"
// End:
