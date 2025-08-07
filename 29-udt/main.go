package main

func main() {
	var address1 struct {
		Line1, Pincode, Status string
		Display                func()
	}

	address1 = struct {
		Line1, Pincode, Status string
		Display                func()
	}{
		Line1:   "line-1 address",
		Pincode: "124210",
		Status:  "active",
		Display: func() {
			println("Line1:", address1.Line1)
			println("Pincode:", address1.Pincode)
			println("Status:", address1.Status)
		},
	}
	address1.Display()

	var address2 struct {
		Line1, Pincode, Status string
	}

	address2 = struct {
		Line1, Pincode, Status string
	}{
		Line1:   "line-1 address",
		Pincode: "124210",
		Status:  "active",
	}

	Display(address2)

	// epty struct variable
	var empty struct{} = struct{}{}
	println(empty)
}

func Display(addr struct {
	Line1, Pincode, Status string
}) {
	println("Line1:", addr.Line1)
	println("Pincode:", addr.Pincode)
	println("Status:", addr.Status)
}

// type Address struct {
// 	Line1, Pincode, Status string
// 	Display                func()
// }
