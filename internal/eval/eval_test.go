package eval_test

import (
	"monkey/internal/eval"
	"monkey/internal/lexer"
	"monkey/internal/object"
	"monkey/internal/parser"
	"testing"
)

func TestEval(t *testing.T) {
	t.Run("test eval integer expression", func(t *testing.T) {
		type testIntExpr struct {
			input    string
			expected int64
		}

		tests := []testIntExpr{
			{input: "5", expected: 5},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			testIntegerObject(t, eval, tt.expected)
		}
	})

	t.Run("test eval boolean expression", func(t *testing.T) {
		type testBoolExpr struct {
			input    string
			expected bool
		}

		tests := []testBoolExpr{
			{input: "true", expected: true},
			{input: "false", expected: false},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			testBooleanObject(t, eval, tt.expected)
		}
	})
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return eval.Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("obj is not *object.Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("obj has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("obj is not *object.Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("obj has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}