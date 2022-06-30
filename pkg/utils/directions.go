package utils

const NORTH = "north"
const SOUTH = "south"
const EAST = "east"
const WEST = "west"

type Direction struct {
	cardinalDirection string
}

func GetNorthCardinalDirection() *Direction {
	return &Direction{
		cardinalDirection: NORTH,
	}
}

func GetSouthCardinalDirection() *Direction {
	return &Direction{
		cardinalDirection: SOUTH,
	}
}

func GetEastCardinalDirection() *Direction {
	return &Direction{
		cardinalDirection: EAST,
	}
}

func GetWestCardinalDirection() *Direction {
	return &Direction{
		cardinalDirection: WEST,
	}
}

func GetDirectionByString(direction string) *Direction {
	switch direction {
	case NORTH:
		return GetNorthCardinalDirection()
	case SOUTH:
		return GetSouthCardinalDirection()
	case EAST:
		return GetEastCardinalDirection()
	case WEST:
		return GetWestCardinalDirection()
	default:
		return nil
	}
}

func (d *Direction) GetOppositeDirection() *Direction {
	switch d.cardinalDirection {
	case NORTH:
		return GetSouthCardinalDirection()
	case SOUTH:
		return GetNorthCardinalDirection()
	case EAST:
		return GetWestCardinalDirection()
	case WEST:
		return GetEastCardinalDirection()
	default:
		return nil
	}
}

func (d *Direction) String() any {
	return d.cardinalDirection
}
