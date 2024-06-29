package object

import (
	"bytes"
	"fmt"
	"monkey/internal/ast"
	"strings"
)

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return T_INTEGER
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return T_BOOL
}

type String struct {
	Value string
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return T_STRING }

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() ObjectType {
	return T_NULL
}

type Function struct {
	Params []*ast.ID
	Body   *ast.BlockStatement
	Env    *Environment
}

func (f *Function) Type() ObjectType { return T_FUNCTION }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := make([]string, len(f.Params))
	for i := range f.Params {
		params[i] = f.Params[i].String()
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Type() ObjectType {
	return T_RETURN_VALUE
}

type BuiltinFunc func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunc
}

func (b *Builtin) Type() ObjectType { return T_BUILTIN }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return T_ERROR }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
