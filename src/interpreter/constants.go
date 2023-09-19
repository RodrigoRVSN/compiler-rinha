package interpreter

type Term = string

/*
type File struct {
	name       string
	expression Term
	location   Location
}

type Location struct {
	start    int
	end      int
	filename string
}

type Var struct {
	kind string
	text string
	location Location
} */

const (
	IF       = "If"
	BINARY   = "Binary"
	CALL     = "Call"
	STRING   = "Str"
	INTEGER  = "Int"
	BOOLEAN  = "Bool"
	PRINT    = "Print"
	LET      = "Let"
	FUNCTION = "Function"
	TUPLE    = "Tuple"
	VARIABLE = "Var"
)
