package main

func main() {

	cd1 := ColourCode{100, 200, "RED"}
	cd1.Display()

	cd2 := ColourCode{int: 100, string: "RED"}
	cd2.Display()

}

type integer = int // just alias

// anonymous fields
type ColourCode struct {
	int
	integer
	string
}

func (c *ColourCode) Display() {
	println("Code:", c.int)
	println("Code2:", c.integer)
	println("Name:", c.string)
}
