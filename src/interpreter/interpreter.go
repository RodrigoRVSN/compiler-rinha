package interpreter

import "fmt"

const (
	Kind  = "kind"
	Value = "value"
)

func getNode(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

func Interprete(ast interface{}) interface{} {
	switch node := getNode(ast, Kind); node {

	case "Str":
		return getNode(ast, Value)

	case "Print":
		printValue := getNode(ast, Value)
		value := getNode(printValue, Value)
		fmt.Println(value)

	default:
		panic("This node was not found!")
	}

	return "XD"
}
