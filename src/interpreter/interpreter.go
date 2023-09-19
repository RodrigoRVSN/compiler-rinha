package interpreter

import "fmt"

func getNode(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

func Interprete(ast interface{}, context map[string]interface{}) interface{} {
	switch node := getNode(ast, Kind); node {

	case Str, Int:
		return getNode(ast, Value)

	case Print:
		printValue := getNode(ast, Value)
		value := getNode(printValue, Value)
		fmt.Println(value)

	case Let:
		contextCopy := context

		letName := getNode(getNode(ast, "name"), "text").(string)
		value := Interprete(getNode(ast, "value"), contextCopy)

		contextCopy[letName] = value

		return Interprete(getNode(ast, "next"), contextCopy)

	case Function:
		return ast

	case If:
		value := getNode(ast, Value)
		ifValue := getNode(value, Value)
		fmt.Println(ifValue)

	case Var:
		varName := getNode(ast, "text").(string)
		if value, wasFound := context[varName]; wasFound{
			return value
		}
		panic(fmt.Sprintf("Variable '%s' doesn't exists!", varName))

	default:
		panic(fmt.Sprintf("This node '%s' was not found my friendo", node))
	}

	return "XD"
}
