package main

import (
	"errors"
	"fmt"
)

type MyMap map[string]any

func (m MyMap) Delete(key string) error {
	if m == nil {
		return errors.New("input map is nil")
	}

	if _, ok := m[key]; !ok {
		return fmt.Errorf("key-->%v does not exist", key)
	}

	delete(m, key)
	return nil
}

func (m MyMap) GetKeysNdValues() ([]string, []any) {
	keys := make([]string, len(m))
	values := make([]any, len(m))
	count := 0
	for k, v := range m {
		keys[count] = k
		values[count] = v
		count++
	}
	return keys, values
}

func (m MyMap) Display() {
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func main() {

	m1 := make(MyMap) // can directly use built in functions/constructs on a user defiend type if applicable
	m1["560086"] = "Bangaluru-1"
	m1["560096"] = "Bangaluru-2"
	m1["560034"] = "Bangaluru-3"
	m1["12345"] = "No Where"
	m1.Display()

	if err := m1.Delete("12345"); err != nil {
		println(err.Error())
	} else {
		println("Successfully deleted")
	}

	keys, values := m1.GetKeysNdValues()
	fmt.Println(keys)
	fmt.Println(values)

	m2 := make(map[string]any)
	m2["560086"] = "Bangaluru-1"
	m2["560096"] = "Bangaluru-2"
	m2["560034"] = "Bangaluru-3"
	m2["12345"] = "No Where"

	MyMap(m2).Display()
	if err := MyMap(m2).Delete("12345"); err != nil {
		println(err.Error())
	} else {
		println("Successfully deleted")
	}

	keys, values = MyMap(m2).GetKeysNdValues()
	fmt.Println(keys)
	fmt.Println(values)
}
