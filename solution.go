package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myCustomInt int

func CalcSquare(sideLen float64, sidesNum myCustomInt) float64 {
	sideLenSquare := sideLen * sideLen
	if sidesNum == 2 {
		return sideLenSquare
	}
	if sidesNum == 3 {
		return sideLenSquare * math.Sqrt(3) / 4
	}
	if sidesNum == 0 {
		return math.Pi * sideLenSquare
	}
	return sideLenSquare
}
