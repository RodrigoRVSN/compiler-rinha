package main

import (
	"fmt"
	"time"

	"github.com/rodrigorvsn/compiler-rinha/src/interpreter"
	"github.com/rodrigorvsn/compiler-rinha/src/utils"
)

func main() {
	start := time.Now()

	var ast = utils.ExtractAST(utils.GetAbsolutePath("../../files/print.json"))
	interpreter.Interprete(ast)

	difference := time.Since(start)
	fmt.Println("Difference", difference)
}
