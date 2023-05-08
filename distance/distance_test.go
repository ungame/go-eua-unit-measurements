package distance

import "testing"

func TestCalcInches(t *testing.T) {
	got := CalcInches(1)
	if got != 2.54 {
		t.Errorf("CalcInches(1) = %f; want = 2.54", got)
	}
}

func TestCalcFeet(t *testing.T) {
	got := CalcFeet(1)
	if got != 30.48 {
		t.Errorf("CalcFeet(1) = %f; want = 30.48", got)
	}
}

func TestCalcYards(t *testing.T) {
	got := CalcYards(1)
	if got != 91.44 {
		t.Errorf("CalcYard(1) = %f; want = 91.44", got)
	}
}

func TestCalcMiles(t *testing.T) {
	got := CalcMiles(1)
	if got != 160900 {
		t.Errorf("CalcMiles(1) = %f; want = 160900", got)
	}
}

func TestCalcNauticalMiles(t *testing.T) {
	got := CalcNauticalMiles(1)
	if got != 185200 {
		t.Errorf("CalcNauticalMiles(1) = %f; want = 185200", got)
	}
}
