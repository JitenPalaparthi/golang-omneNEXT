package main

import "fmt"

type Person struct {
	Id            uint
	Name          string
	Email, Mobile string
}

func main() {
	var p1 Person
	p1.Id = 100
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outlook.com"
	p1.Mobile = "9618558500"

	p2 := Person{101, "Jitenp", "Jiten.Palaparthi@Gmail.com", "9618558500"}
	p3 := Person{Id: 102, Name: "Jiten"}

	// pointer

	p4 := new(Person)
	p4.Id = 100
	p4.Name = "Jiten"
	p4.Email = "JitenP@Outlook.com"
	p4.Mobile = "9618558500"

	p5 := &Person{101, "Jitenp", "Jiten.Palaparthi@Gmail.com", "9618558500"}

	fmt.Println(p1, "\n", p2, "\n", p3, "\n", p4, "\n", p5)

	fmt.Println("Id:", p1.Id)
	fmt.Println("Name:", p1.Name)

	(*p5).Mobile = "9191919191"
	p5.Email = "Jp@Spanlet.com"

	// not related to struct, but related to arrays and array pointer
	arr := [3]int{10, 20, 30}
	var arrptr *[3]int = &arr
	(*arrptr)[0] = 100
	arrptr[1] = 200
}
