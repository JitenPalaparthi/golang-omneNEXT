package rect

type Rect struct {
	L, B float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{l, b}
}

func NewRectDefaults() *Rect {
	return &Rect{1, 1}
}

func (r *Rect) Area() float64 {
	return float64(r.L * r.B)
}

func (r *Rect) Perimeter() float64 {
	return float64(2 * (r.L + r.B))
}

func (r *Rect) Display() {
	println("L:", r.L, "\nB:", r.B)
}

func (r *Rect) What() string {
	return "Rect"
}
