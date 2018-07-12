package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	sum := 0
	pre := 1
	for i := 1; i <= 20; i++ {
		pre = pre * i
		sum = sum + pre
	}
	fmt.Println("階乘和：", sum)
	fmt.Println("時間：", time.Since(t1))
}
