package main

func move(position Coordinates, heading Heading, distance int) (newPosition Coordinates) {
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

func turnRight(heading Heading, degrees int) (newHeading Heading) {
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

func turnLeft(heading Heading, degrees int) (newHeading Heading) {
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

func changeHeading(heading Heading, command Command) (newHeading Heading) {
	if command.Direction == "L" {
		newHeading = turnLeft(heading, command.Distance)
	} else {
		newHeading = turnRight(heading, command.Distance)
	}

	return
}

func steer(position, waypoint Coordinates, heading Heading, command Command) (newPosition, newWaypoint Coordinates, newHeading Heading) {
	newWaypoint = waypoint

	switch command.Direction {
	case "F":
		newPosition = move(position, heading, command.Distance)
		newHeading = heading
	case "R", "L":
		newPosition.X, newPosition.Y = position.X, position.Y
		newHeading = changeHeading(heading, command)
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

// Local Variables:
// compile-command: "go build"
// End:
