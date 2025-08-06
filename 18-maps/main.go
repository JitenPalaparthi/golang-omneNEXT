package main

import (
	"errors"
	"fmt"
)

func main() {

	var m1 map[string]any

	if m1 == nil {
		println("nil map")
		m1 = make(map[string]any)
	}

	m1["560086"] = "Bangalore-1"
	m1["560096"] = "Bangalore-2"
	m1["522001"] = "Guntur-1"
	m1["695001"] = "Trivandrum-1"
	// m1["1"] = true
	// m1["2"] = "Hello World"

	v := m1["560086"]
	fmt.Println(v)

	for k, v := range m1 {
		fmt.Println("Key:", k, "Value:", v)
	}

	// var any1 any = 100

	// v2, ok := any1.(int)
	// if ok {

	// }
	v1, ok := m1["5600906"]
	if ok {
		fmt.Println(v1)
	} else {
		println("key does not exist")
	}

	//delete(m1, "522001")
	if err := Delete(m1, "522001"); err != nil {
		fmt.Println(err.Error())
	} else {
		println("key successfully deleted")
	}
	fmt.Println(m1)
	if err := Delete(m1, "522001"); err != nil {
		fmt.Println(err.Error())
	} else {
		println("key successfully deleted")
	}

	clear(m1)
	fmt.Println(m1, len(m1))
}

func Delete(m map[string]any, key string) error {
	if m == nil {
		return errors.New("input map is nil")
	}

	if _, ok := m[key]; !ok {
		return fmt.Errorf("key-->%v does not exist", key)
	}

	delete(m, key)
	return nil
}

// Create 4 functions add, sub, mul ,div
// create a function greet
// create function sq

// create a map[string]any

// add, sub, mul, div,greet,sq
// assign respective functions to the map as values
// execute them in a range loop
