package main

import (
	"fmt"
	"time"

	"github.com/rodrigorvsn/compiler-rinha/src/interpreter"
	"github.com/rodrigorvsn/compiler-rinha/src/utils"
)

func main() {
	start := time.Now()

	var ast = utils.ExtractAST(utils.GetAbsolutePath("../../files/fib.json"))

	context := make(map[string]interface{})
	interpreter.Interpret(ast, context)

	difference := time.Since(start)
	fmt.Println("The code was executed in", difference)
}
