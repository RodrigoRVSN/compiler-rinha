package interpreter

import (
	"fmt"
	"reflect"
)

func getNode(ast interface{}, key string) interface{} {
	return ast.(map[string]interface{})[key]
}

func Interpret(ast interface{}, context map[string]interface{}) interface{} {
	switch node := getNode(ast, "kind"); node {
	case STRING, INTEGER, BOOLEAN:
		return getNode(ast, "value")

	case PRINT:
		printValue := getNode(ast, "value")
		value := Interpret(printValue, context)
		fmt.Println(value)

	case LET:
		letName := getNode(getNode(ast, "name"), "text").(string)
		value := Interpret(getNode(ast, "value"), context)

		context[letName] = value

		return Interpret(getNode(ast, "next"), context)

	case FUNCTION, TUPLE:
		return ast

	case IF:
		condition := Interpret(getNode(ast, "condition"), context)

		if condition.(bool) {
			then := getNode(ast, "then")
			return Interpret(then, context)
		}

		otherwise := getNode(ast, "otherwise")
		return Interpret(otherwise, context)

	case VARIABLE:
		varName := getNode(ast, "text").(string)
		if value, wasFound := context[varName]; wasFound {
			if _, isObject := value.(map[string]interface{}); isObject {
				return Interpret(value, context)
			}
			return value
		}
		panic(fmt.Sprintf("Variable '%s' doesn't exists!", varName))

	case CALL:
		function := Interpret(getNode(ast, "callee"), context)

		if getNode(function, "kind") != FUNCTION {
			return nil
		}

		arguments := getNode(ast, "arguments").([]interface{})
		parameters := getNode(function, "parameters").([]interface{})

		for index := range arguments {
			if _, ok := arguments[index].(map[string]interface{}); ok {
				if getNode(arguments[index], "kind") == nil {
					arguments[index] = getNode(arguments[index], "value")
				} else {
					arguments[index] = Interpret(arguments[index], context)
				}
			}
		}

		for index := range parameters {
			if _, ok := parameters[index].(map[string]interface{}); ok {
				if getNode(parameters[index], "kind") == nil {
					parameters[index] = getNode(parameters[index], "text")
				} else {
					parameters[index] = Interpret(parameters[index], context)
				}
			}
		}

		for index := range parameters {
			context[parameters[index].(string)] = arguments[index]
		}

		return Interpret(getNode(function, "value"), context)

	case BINARY:
		left := Interpret(getNode(ast, "lhs"), context)
		right := Interpret(getNode(ast, "rhs"), context)

		if reflect.TypeOf(left).Kind() == reflect.Float64 {
			left = int32(left.(float64))
		}

		if reflect.TypeOf(right).Kind() == reflect.Float64 {
			right = int32(right.(float64))
		}

		switch operator := getNode(ast, "op").(string); operator {
		case "Lt":
			return left.(int32) < right.(int32)
		case "Add":
			return left.(int32) + right.(int32)
		case "Sub":
			return left.(int32) - right.(int32)
		default:
			panic(fmt.Sprintf("Unknown operator: <%s>", operator))
		}

	default:
		panic(fmt.Sprintf("This node '%s' was not found my friendo", node))
	}

	return "XD"
}
