package main

import (
	"fmt"
)

func main() {
	var score [5]int // 成績
	var sum int = 0  // 總和
	var avg int = 0  // 平均

	score[0] = 100
	score[1] = 90
	score[2] = 80
	score[3] = 70
	score[4] = 60

	for i := 0; i < len(score); i++ {
		sum += score[i]
	}

	avg = sum / len(score)

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := getMin(x[0], x)

	fmt.Println("成績：", score)
	fmt.Println("總和：", sum)
	fmt.Println("平均：", avg)
	fmt.Println("x的最小值：", min)
}

// 取得最小值
func getMin(first int, num []int) int {
	for _, value := range num {
		if value < first {
			first = value
		}
	}

	return first
}
