package main

import "fmt"

func CheckYourself(ptr *int, x, val int) int {
	ret := ptr
	fmt.Println(ret, ptr, val)
	*ptr += val
	fmt.Println(*ptr)
	x += val
	fmt.Println(x, ret, *ret)
	return 2 * *ret
}

func main() {
	i := 6
	j := 11
	fmt.Println(CheckYourself(&i, j, 4), i, j)
}
