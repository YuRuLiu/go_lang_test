package main

import (
	"fmt"
	"time"
)

type loggingFactory struct {
	paperPulp float32 // 每秒產生的紙漿量
}

type paperFactory struct {
	paperPulp float32 // 每秒消耗的紙漿量
	paper     int     // 每秒製造的紙張數
}

type printingFactory struct {
	paper int // 每秒印刷的紙張數
}

func main() {
	// 伐木工廠1
	loggingFactory1 := loggingFactory{1}
	var totalPaperPulp float32 = 0
	generatePaperPulp(loggingFactory1.paperPulp, &totalPaperPulp)
}

// 生產紙漿
func generatePaperPulp(paperPulp float32, totalPaperPulp *float32) {
	for {
		*totalPaperPulp = *totalPaperPulp + paperPulp
		time.Sleep(1 * time.Second)
		fmt.Println("紙漿：", *totalPaperPulp)
	}
}
