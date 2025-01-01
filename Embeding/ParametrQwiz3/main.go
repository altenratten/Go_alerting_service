package main

import "fmt"

// напишите функцию Not, которая будет изменять массив [5][10]int
// ...

func Not(arr *[5][10]int) {
	for i, v := range *arr {
		for j, item := range v {
			if item == 1 {
				(*arr)[i][j] = 0
			} else {
				(*arr)[i][j] = 1
			}
		}
	}
}

func main() {
	arr := [5][10]int{
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
	}
	// вставьте здесь вызов Not
	// ...
	Not(&arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
