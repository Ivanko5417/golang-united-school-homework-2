package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myCustomInt int

const SidesTriangle myCustomInt = 3
const SidesSquare myCustomInt = 4
const SidesCircle myCustomInt = 0

func CalcSquare(sideLen float64, sidesNum myCustomInt) float64 {
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
