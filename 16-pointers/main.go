package main

var Global int = 1000

func main() {
	var num1 int = 100
	var ptr1 *int = &num1
	ptr1 = &Global

	// ptr1 = ptr1 + 8

	// 4f4308 D main.Global
	println(ptr1)
	println(&num1)
	slice := make([]int, 100000, 1000000)
	println(slice)
	var arr [1000]int
	println(&arr)

	ptr2 := new(int) // stack
	*ptr2 = 1000

	var num2 uint8 = 127

	var ptr3 *uint8 = &num2

	println(ptr3)

	*ptr3 = 255

	num3 := 100
	println(Sq1(num1))
	println(Sq2(&num1))
	println((*int)(Sq3(num1)))
	println((*int)(Sq4(num3)))
	Sq5(&num3)
	println(num3)

	var ptr10 *int // what is the zero value of a pointer --> nil
	println(Sq2(nil))
	println(Sq2(ptr10))

	var ptr4 **int = &ptr1
	var ptr5 ***int = &ptr4
	println(***ptr5)

}

func Sq1(num int) int {
	return num * num
}

func Sq2(num *int) int {
	// if num != nil {
	// 	return 0
	// }
	return *num * *num
}

func Sq3(num int) *int {
	var ptr *int = &num // created inside dangling pointer
	*ptr = num * num
	return ptr
}

func Sq4(num int) *int {
	ptr := new(int) // heap
	*ptr = num * num
	return ptr
}

func Sq5(ptr *int) {
	*ptr = *ptr * *ptr
}
