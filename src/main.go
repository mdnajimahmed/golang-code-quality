package main

import (
	"github.com/mandatorySuicide/golang-code-quality/src/app"
)

func main() {
	//fmt.Println("MAIN INIT")
	app.SetupServer().Run()
}
