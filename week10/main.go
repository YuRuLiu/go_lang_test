/*
# Week10 練習

題目：
撰寫一程式可對目標API 進行連線，需可自行調整併發數量

記錄：
單一連線花費時間
單一最大花費時間
單一最小花費時間
壓測總數花費時間
平均花費時間

參考關鍵字：cron, slice, (t Time) Sub, goroutine
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var num int = 10
var time_data []time.Duration

func main() {
	//建立WaitGroup
	var wg sync.WaitGroup

	//總共N次
	wg.Add(num)

	for i := 0; i < num; i++ {
		go getAPI(&wg)
	}

	//等待所有工作完成
	wg.Wait()

	fmt.Println("API：api/bank_info/currency")
	fmt.Println("次數：", num)
	fmt.Println("單一連線花費時間：", time_data)
	fmt.Println("單一最大花費時間：", getMaxTime(time_data))
	fmt.Println("單一最小花費時間：", getMinTime(time_data))
	fmt.Println("壓測總數花費時間：", getTotalTime(time_data))
	fmt.Println("平均花費時間：", getAvgTime(time_data, num))
}

func getBank() string {
	url := "http://bb.dev.d2:80/api/bank_info/currency"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func getAPI(wg *sync.WaitGroup) {
	t1 := time.Now()
	getBank()
	time_data = append(time_data, time.Since(t1))

	//工作完成 回報WaitGroup -1
	wg.Done()
}

// 單一最小花費時間
func getMinTime(num []time.Duration) time.Duration {
	min := num[0]

	for _, value := range num {
		if value <= min {
			min = value
		}
	}

	return min
}

// 單一最大花費時間
func getMaxTime(num []time.Duration) time.Duration {
	max := num[0]

	for _, value := range num {
		if value >= max {
			max = value
		}
	}

	return max
}

// 壓測總數花費時間
func getTotalTime(total []time.Duration) time.Duration {
	var total_time time.Duration
	for _, value := range total {
		total_time += value
	}

	return total_time
}

// 取平均花費時間
func getAvgTime(total []time.Duration, num int) time.Duration {
	total_time := getTotalTime(total)
	nano_time := time.Duration.Nanoseconds(total_time)
	avg := nano_time / int64(num)
	avg_time := time.Duration(avg)

	return avg_time
}
