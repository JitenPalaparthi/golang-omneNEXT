package main

func main() {
	a, b := 10, 20
	c := a + b/2*(a+5) + (b-2)*10 + 20 //Is this an atomic operation?
	// 10 + 20/2*(10+5)+(20-2)*10+20
	// 10 + 20/2*15+(20-2)*10+20
	// 10 + 20/2*15+18*10+20
	// 10 + 10 * 15+18*10+20
	// 10+ 150 + 18* 10+ 20
	// 10 + 150 + 180 + 20
	// 360
	println(c)
}

// What is an atomic operation?
// How does that expression is evaluated?
