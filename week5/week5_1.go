package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

// 錯誤訊息：panic: send on closed channel
// 錯誤原因：一直送資料給被關閉的channel
// 正常輸出：不使用的通道可以不要顯示關閉，會自行回收
// 參考：https://www.jianshu.com/p/e89dfebe2c4a
