package main

import (
	"fmt"
	"github.com/mandatorySuicide/golang-code-quality/src/service"
	"github.com/mandatorySuicide/golang-code-quality/src/util"
)

func main() {
	bsElem := util.MakeRandInt(10)
	msElem := util.MakeRandInt(10)
	service.BubbleSort(bsElem, func(a, b int) int {
		return a - b
	})
	service.MergeSort(msElem, func(a, b int) int {
		return a - b
	})
	fmt.Println("Hello World")
}
