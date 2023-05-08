package distance

const (
	Base         = 1.0               // centimeter
	Inch         = Base * 2.54       // 2.54 centimeters
	Foot         = Inch * 12         // 30.48 centimeters
	Yard         = Foot * 3          // 91.44 centimeters
	Mile         = Base * 100 * 1609 // 1609 meters = 160900 centimeters
	NauticalMile = Base * 100 * 1852 // 1852 meters = 185200 centimeters
)

// CalcInches compute inch to centimeters
func CalcInches(n float64) float64 {
	return Inch * n
}

// CalcFeet compute feet to centimeters
func CalcFeet(n float64) float64 {
	return Foot * n
}

// CalcYards compute yards to centimeters
func CalcYards(n float64) float64 {
	return Yard * n
}

// CalcMiles compute miles to centimeters
func CalcMiles(n float64) float64 {
	return Mile * n
}

// CalcNauticalMiles compute nautical miles to centimeters
func CalcNauticalMiles(n float64) float64 {
	return NauticalMile * n
}
