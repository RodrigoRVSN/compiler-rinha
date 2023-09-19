package main

import (
	"fmt"
	"time"

	"github.com/rodrigorvsn/compiler-rinha/src/interpreter"
	"github.com/rodrigorvsn/compiler-rinha/src/utils"
)

func main() {
	start := time.Now()

	var ast = utils.ExtractAST(utils.GetAbsolutePath("../../files/sum.json"))

	context := make(map[string]interface{})
	interpreter.Interprete(ast, context)

	difference := time.Since(start)
	fmt.Println("The code was executed in", difference)
}
