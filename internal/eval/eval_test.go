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
			{input: "10", expected: 10},
			{input: "-5", expected: -5},
			{input: "-51", expected: -51},
			{input: "-51 + 10", expected: -41},
			{input: "51 + 10", expected: 61},
			{input: "-51 + 61", expected: 10},
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
			{input: "!true", expected: false},
			{input: "!false", expected: true},
			{input: "!!true", expected: true},
			{input: "!!false", expected: false},
			{input: "1 < 2", expected: true},
			{input: "1 > 2", expected: false},
			{input: "1 <= 2", expected: true},
			{input: "1 >= 2", expected: false},
			{input: "1 == 1", expected: true},
			{input: "1 != 1", expected: false},
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
