package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("目前書單：")
	getlist()
	fmt.Println("訂書囉！")
	booking()
	fmt.Println("最新書單：")
	getlist()
}

func getlist() {
	url := "http://localhost:1323/booklist"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func booking() {
	url := "http://localhost:1323/booking"

	payload := strings.NewReader("name=betsy&bookName=白夜行&author=東野圭吾")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
