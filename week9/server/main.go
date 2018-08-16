package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type Weather struct {
	region string `gorm:"column:'region';default:'taipei_city'"`
	info   string `gorm:"column:'info';default:''"`
}

func main() {
	e := echo.New()
	e.GET("/weather", getWeather)
	e.Logger.Fatal(e.Start(":1323"))
}

func getWeather(c echo.Context) error {
	// 地區
	region := c.QueryParam("region")

	source := "Redis:"
	weather := getRedis(region)
	if weather == "" {
		source = "DB"
		weather = getDB(region)
		setRedis(region, weather)
		if weather == "" {
			source = "API:"
			weather = getAPI(region)
			setDB(region, weather)
		}
	}

	show := source + string(weather)

	return c.String(http.StatusOK, show)
}

func getAPI(region string) string {
	url := "http://weather.json.tw/api?region=" + region

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func getRedis(key string) string {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return ""
	}
	defer c.Close()

	weather, err := redis.String(c.Do("GET", key))
	if err != nil {
		return ""
	} else {
		return weather
	}
}

func setRedis(key string, value string) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", key, value, "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func getDB(region string) string {
	var weather Weather

	db, err := gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connect to DB error", err)
		return ""
	}
	defer db.Close()

	db.Where("region = ?", region).Find(&weather)

	return weather.info
}

func setDB(region string, info string) {
	weather := Weather{region, info}

	db, err := gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connect to DB error", err)
		return
	}
	defer db.Close()

	db.Create(&weather)
}
