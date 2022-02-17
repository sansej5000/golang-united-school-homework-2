package square

import "math"

type side int

const (
	SidesTriangle side = 3
	SidesSquare   side = 4
	SidesCircle   side = 0
)

func CalcSquare(sideLen float64, sidesNum side) float64 {

	if sidesNum == SidesTriangle {
		return math.Sqrt(math.Pow(sideLen, 2)-math.Pow(sideLen/2, 2)) * sideLen / 2
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2.0)
	} else if sidesNum == SidesCircle {
		return math.Pi * math.Pow(sideLen, 2.0)
	} else {
		return 0
	}

}
