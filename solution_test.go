package square

import "testing"

func Test_SidesTriangle(t *testing.T) {
	expected := 10.825317547305485
	res := CalcSquare(5, 3)
	if expected != res {
		t.Errorf("result is not correct, expected result %f", expected)
	}
}

func Test_SidesSquare(t *testing.T) {
	expected := 25.0
	res := CalcSquare(5, 4)
	if expected != res {
		t.Errorf("result is not correct, expected result %f", expected)
	}
}

func Test_SidesCircle(t *testing.T) {
	expected := 78.53981633974483
	res := CalcSquare(5, 0)
	if expected != res {
		t.Errorf("result is not correct, expected result %f", expected)
	}
}

func Test_SideIncorrect(t *testing.T) {
	expected := 0.0
	res := CalcSquare(5, 5)
	if expected != res {
		t.Errorf("result is not correct, expected result %f", expected)
	}
}
