package main

func main() {

	day := uint(5)

	switch day {
	case 1:
		println("Sunday")
	case 2:
		{
			println("Monday")
		}
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("Noday")
	}

	char := 'A'

	switch char {

	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is a vovel")
	default:
		println(string(char), "is either a consonent or a special char")
	}

	num := -51

	switch { // empty switch
	case num >= 0 && num < 50:
		println(num, "is between 0-50")
	case num >= 50 && num < 100:
		println(num, "is between 50 to 100")
	case num >= 10:
		println(num, "is greater than 100")
	default:
		println(num, "is a negative number")
	}

	// fallthrough

	num = 4

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	println("false negative while using fallthrough")
	// when ever you remove break in many other programming languages , you have to add fallthrough in Golang
	num = 4
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}
}
