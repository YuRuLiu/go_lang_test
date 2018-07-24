package main

import (
	"fmt"
	"go_package/phone"
)

type iphone struct {
	browserName string
	wifiName    string
	appName     string
}

type htc struct {
	browserName string
	wifiName    string
	appName     string
}

func (i iphone) Browser() string {
	return i.browserName
}

func (i iphone) Wifi() string {
	return i.wifiName
}

func (i iphone) App() string {
	return i.appName
}

func (h htc) Browser() string {
	return h.browserName
}

func (h htc) Wifi() string {
	return h.wifiName
}

func (h htc) App() string {
	return h.appName
}

func main() {
	iphone7 := iphone{"safari", "on", "Youtube"}
	fmt.Println("iphone7:")
	phone.Internet(iphone7)
	phone.WatchVedio(iphone7)

	u12 := htc{"chrome", "on", "Line TV"}
	fmt.Println("HTC U12:")
	phone.Internet(u12)
	phone.WatchVedio(u12)
}
