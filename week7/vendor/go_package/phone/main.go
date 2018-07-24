package phone

type phone interface {
	Browser() string
	Wifi() string
	App() string
}

func Internet(p phone) {
	if p.Wifi() == "on" {
		println("使用" + p.Browser() + "上網")
	} else {
		println("無法連線")
	}
}

func WatchVedio(p phone) {
	if p.Wifi() == "on" {
		println("使用" + p.App() + "看影片")
	} else {
		println("無法連線")
	}
}
