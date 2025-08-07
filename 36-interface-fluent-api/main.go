package main

func main() {

	//db.Select().Where().Orderby().Filter().Exec()
	//	resukt := Data.Add(10).Sub(10).Mul(5).Add(400).Div(5).Get()

	r := NewCalc(100).Add(10).Mul(2).Add(20).Sub(10).Div(2).Get()

	println(r)
}

type Calc struct {
	data int
}

func NewCalc(d int) *Calc {
	return &Calc{d}
}

func (c *Calc) Add(d int) ICalc {
	c.data += d
	return c
}

func (c *Calc) Sub(d int) ICalc {
	c.data -= d
	return c
}

func (c *Calc) Mul(d int) ICalc {
	c.data *= d
	return c
}
func (c *Calc) Div(d int) ICalc {
	c.data /= d
	return c
}
func (c *Calc) Get() int {
	return c.data
}

type ICalc interface {
	Add(int) ICalc
	Sub(int) ICalc
	Mul(int) ICalc
	Div(int) ICalc
	Get() int
}
