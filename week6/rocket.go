package main

import (
	"fmt"
	"go_package/rocket"
	"time"
)

type rocketA struct {
	name string
}

type rocketB struct {
	name string
}

func (r rocketA) Launch() {
	for i := 10; i > 0; i-- {
		fmt.Println(r.name, "倒數：", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(r.name, "發射！！！！！")
}

func (r rocketB) Launch() {
	for i := 5; i > 0; i-- {
		fmt.Println(r.name, "倒數：", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println(r.name, "發射！！！！！")
}

func main() {
	rocket1 := rocketA{"A火箭"}
	rocket2 := rocketB{"B火箭"}
	go rocket.Launch(rocket1)
	go rocket.Launch(rocket2)

	time.Sleep(20 * time.Second)
}
