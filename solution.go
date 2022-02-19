package square

import "math"

type MyCustomInt int

const SidesTriangle MyCustomInt = 3
const SidesSquare MyCustomInt = 4
const SidesCircle MyCustomInt = 0

func CalcSquare(sideLen float64, sidesNum MyCustomInt) float64 {
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
