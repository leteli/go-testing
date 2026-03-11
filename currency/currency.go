package currency

type Converter interface {
	Convert(amount float64, from, to string) float64
}

func PriceIn(amount float64, from, to string, c Converter) float64 {
	return c.Convert(amount, from, to)
}
