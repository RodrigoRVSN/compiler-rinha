package utils

import (
	"encoding/json"
	"os"
)

var ast struct {
	Expression map[string]interface{} `json:"expression"`
}

func ExtractAST(jsonPath string) map[string]interface{} {
	rawData, error := os.ReadFile(jsonPath)

	if error != nil {
		panic(error)
	}

	if error := json.Unmarshal(rawData, &ast); error != nil {
		panic(error)
	}

	return ast.Expression
}
