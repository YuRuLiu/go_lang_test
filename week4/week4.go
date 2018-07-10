package main

import (
	"fmt"
	"sync"
	"time"
)

// 伐木廠
type loggingFactory struct {
	paperPulp float32 // 每秒產生的紙漿量
}

// 造紙廠
type paperFactory struct {
	paperPulp float32 // 每秒消耗的紙漿量
	paper     int     // 每秒製造的紙張數
}

// 印刷廠
type printingFactory struct {
	paper int // 每秒印刷的紙張數
}

func main() {
	//建立WaitGroup
	var wg sync.WaitGroup

	// 所有紙漿
	var totalPaperPulp float32 = 0

	// 所有紙張
	var totalPaper int = 0

	// 所有印刷的紙
	var totalPrint int = 0

	// 伐木工廠1
	loggingFactory1 := loggingFactory{1}
	go generatePaperPulp(loggingFactory1.paperPulp, &totalPaperPulp)

	// 造紙廠1
	paperFactory1 := paperFactory{0.5, 5000}
	go generatePaper(&totalPaperPulp, paperFactory1, &totalPaper)
	// 造紙廠2
	paperFactory2 := paperFactory{0.3, 3000}
	go generatePaper(&totalPaperPulp, paperFactory2, &totalPaper)

	//WaitGroup +1
	wg.Add(1)

	// 印刷廠1
	printingFactory1 := printingFactory{6000}
	go printPaper(&totalPaper, &totalPrint, printingFactory1, &wg)

	//等待所有工作完成
	wg.Wait()
}

// 生產紙漿
func generatePaperPulp(paperPulp float32, totalPaperPulp *float32) {
	for {
		*totalPaperPulp = *totalPaperPulp + paperPulp
		time.Sleep(1 * time.Second)
		fmt.Println("紙漿：", *totalPaperPulp)
	}
}

// 生產紙張
func generatePaper(totalPaperPulp *float32, factory paperFactory, totalPaper *int) {
	for {
		if *totalPaperPulp > factory.paperPulp {
			*totalPaperPulp = *totalPaperPulp - factory.paperPulp
			*totalPaper = *totalPaper + factory.paper
			time.Sleep(1 * time.Second)
			fmt.Println("紙張：", *totalPaper)
		}
	}

}

// 印刷紙張
func printPaper(totalPaper *int, totalPrint *int, factory printingFactory, wg *sync.WaitGroup) {
	for {
		if *totalPaper > factory.paper {
			*totalPaper = *totalPaper - factory.paper
			*totalPrint = *totalPrint + factory.paper
			time.Sleep(1 * time.Second)
			fmt.Println("印：", *totalPrint)
		}

		if *totalPrint >= 60000 {
			//工作完成 回報WaitGroup -1
			wg.Done()
		}
	}
}
