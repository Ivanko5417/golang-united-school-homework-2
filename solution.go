package square

import "math"

type sidesNumber int

const (
	SidesCircle   sidesNumber = 0
	SidesTriangle sidesNumber = 3
	SidesSquare   sidesNumber = 4
)

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {
	sideLenSquare := sideLen * sideLen
	switch sidesNum {
	case SidesTriangle:
		return sideLenSquare * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLenSquare
	case SidesCircle:
		return math.Pi * sideLenSquare
	}
	return 0
}
