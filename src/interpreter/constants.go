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
	If       = "If"
	Str      = "Str"
	Int      = "Int"
	Bool     = "Bool"
	Print    = "Print"
	Let      = "Let"
	Function = "Function"
	Var      = "Var"
)

const (
	Kind  = "kind"
	Value = "value"
)
