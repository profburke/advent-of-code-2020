package main

func rotateWaypoint(waypoint Coordinates, command Command) (newWaypoint Coordinates) {
	// standard 2d rotation around origin but taking into account
	// fact that angle values are only ever 90, 180, 270 and keeping everything as int
	var s, c int
	switch command.Distance {
	case 90:
		s, c = 1, 0
	case 180:
		s, c = 0, -1
	case 270:
		s, c = -1, 0
	}

	if command.Direction == "L" {
		s *= -1
	}

	newWaypoint.X = waypoint.X*c - waypoint.Y*s
	newWaypoint.Y = waypoint.X*s + waypoint.Y*c

	return
}

func steerByWaypoint(position, waypoint Coordinates, heading Heading, command Command) (newPosition,
	newWaypoint Coordinates, newHeading Heading) {
	newHeading = heading

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

// Local Variables:
// compile-command: "go build"
// End:
