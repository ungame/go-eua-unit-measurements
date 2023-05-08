package weight

const (
	Base  = 1.0                 // gram
	Once  = Base * 28.35        // 28.35 grams
	Pound = Base * 453.6        // 453.6 grams
	Ton   = Base * 1000 * 907.2 // 907200 grams = 907.2 kg
)

func CalcOnce(n float64) float64 {
	return Once * n
}

func CalcPounds(n float64) float64 {
	return Pound * n
}

func CalcTons(n float64) float64 {
	return Ton * n
}
