package calculator

// Calculator ...
type Calculator interface {
	Calculate(ad1, ad2 int) (price int)
}

type calculator struct{}

func (c *calculator) Calculate(ad1, ad2 int) (price int) {
	return ad1 + ad2
}

// NewCalculator returns a new Calculator instance.
func NewCalculator() Calculator {
	return &calculator{}
}
