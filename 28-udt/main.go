package main

import "fmt"

func main() {

	var address1 struct {
		Line1, Pincode, Status string
	}

	address1 = struct {
		Line1, Pincode, Status string
	}{Line1: "line-1 address", Pincode: "124210", Status: "active"}

	address2 := struct {
		Line1, Pincode, Status string
	}{Line1: "line-1 address", Pincode: "124210", Status: "active"}

	var address3 struct {
		Line1, Pincode, Status string
	} = struct {
		Line1, Pincode, Status string
	}{Line1: "line-1 address", Pincode: "124210", Status: "active"}

	fmt.Println(address1, address2, address3)
}
