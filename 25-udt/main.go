package main

import "math/rand/v2"

func main() {

	c1 := New("Jiten", "JitenP@Outlook.Com", "Medical College Trv", "695011")
	c1.Display()
	c1.Address.Display()
}

func New(name, email, line1, pincode string) *Customer {
	return &Customer{Id: rand.IntN(1000), Name: name, Email: email, Address: Address{Line1: line1, Pincode: pincode}}
}

type Customer struct {
	Id          int
	Name, Email string
	Address     Address // composition
}

func (c *Customer) Display() {
	println("ID:", c.Id)
	println("Name:", c.Name)
	println("Email:", c.Email)
}

type Address struct {
	Line1, Pincode string
}

func (a *Address) Display() {
	println("Line1:", a.Line1)
	println("Pincode:", a.Pincode)
}

// Inhe
// Poly
// Direct Encap
// Constr
