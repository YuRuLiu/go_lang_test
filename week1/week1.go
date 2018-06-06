package main

import (
	"fmt"
	"strconv"

	"../betsypackage"
)

func main() {
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9
	var x string = "I Love Golang_"

	g, _ := strToInt(d)
	h := betsypackage.IntToStr(a)

	fmt.Println(a + int(b))
	fmt.Println(a + int(b) + int(c))
	fmt.Println(float32(f) / e)
	fmt.Println(a + g)
	fmt.Println(x + h)
}

func strToInt(str string) (int, error) {
	res, error := strconv.Atoi(str)

	return res, error
}
