package main

import "math/rand/v2"

func main() {

	c1 := New("Jiten", "JitenP@Outlook.Com", "Medical College Trv", "695011", "active", "current")
	c1.Display()

	c1.SocialMap = map[string]string{"x": "jitenp", "linkedin": "linkedin.com/jpalaparthi/"}
	println("Staus:", c1.Address.Status)
	c1.Print()

}

func New(name, email, line1, pincode, status, astatus string) *Customer {
	return &Customer{Id: rand.IntN(1000), Name: name, Email: email, Status: status, Address: struct {
		Line1   string
		Pincode string
		Status  string
	}{Line1: line1, Pincode: pincode, Status: astatus}}
}

type Customer struct {
	Id                  int
	Name, Email, Status string
	Address             struct { // embedded struct
		Line1, Pincode, Status string
	}
	SocialMedia // promoted field
}

type SocialMedia struct {
	SocialMap map[string]string
}

func (s *SocialMedia) Print() {
	for k, v := range s.SocialMap {
		println(k, "-->", v)
	}
}

func (c *Customer) Display() {
	println("ID:", c.Id)
	println("Name:", c.Name)
	println("Email:", c.Email)
	println("Status:", c.Status)
	println("Line1:", c.Address.Line1)
	println("Pincode:", c.Address.Pincode)
	println("Status:", c.Address.Status)

}

// Inhe
// Poly
// Direct Encap
// Constr
