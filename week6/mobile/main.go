package main

import (
	"fmt"
	"time"
)

type CellPhone interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

//他是圓形的
type IWatch struct {
	version string
	radius  int
	battery time.Duration
}

func (w IWatch) Name() string {
	return w.version
}

func (w IWatch) Size() int {
	return w.radius * 3
}

func (w IWatch) TalkTime() time.Duration {
	return w.battery * time.Minute
}

func main() {
	iwatch := IWatch{"Apple Watch Hermès", 30, 24}

	ShowPhone(iwatch)
}

func ShowPhone(c CellPhone) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}
