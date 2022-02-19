package square

import "math"

type SidesNumber int

const SidesTriangle SidesNumber = 3
const SidesSquare SidesNumber = 4
const SidesCircle SidesNumber = 0

func CalcSquare(sideLen float64, sidesNum SidesNumber) float64 {
	sideLenSquare := sideLen * sideLen
	switch sidesNum {
	case SidesTriangle:
		return sideLenSquare * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLenSquare
	case SidesCircle:
		return math.Pi * sideLenSquare
	default:
		return sideLenSquare
	}
}
