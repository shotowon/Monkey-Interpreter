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
			{"true", true},
			{"false", false},
			{"!true", false},
			{"!false", true},
			{"!!true", true},
			{"!!false", false},
			{"1 < 2", true},
			{"1 > 2", false},
			{"1 == 1", true},
			{"1 != 1", false},
			{"true == true", true},
			{"false == false", true},
			{"true == false", false},
			{"true != false", true},
			{"false != true", true},
			{"(1 < 2) == true", true},
			{"(1 < 2) == false", false},
			{"(1 > 2) == true", false},
			{"(1 > 2) == false", true},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			testBooleanObject(t, eval, tt.expected)
		}
	})

	t.Run("test eval if expression", func(t *testing.T) {
		type ifTest struct {
			input    string
			expected interface{}
		}

		tests := []ifTest{
			{"if (true) { 10 }", 10},
			{"if (false) { 10 }", nil},
			{"if (1) { 10 }", 10},
			{"if (1 < 2) { 10 }", 10},
			{"if (1 > 2) { 10 }", nil},
			{"if (1 > 2) { 10 } else { 20 }", 20},
			{"if (1 < 2) { 10 } else { 20 }", 10},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			int, ok := tt.expected.(int)
			if ok {
				testIntegerObject(t, eval, int64(int))
			} else {
				testNullObject(t, eval)
			}
		}
	})
	t.Run("test eval return <int> statement", func(t *testing.T) {
		type returnTest struct {
			input    string
			expected int64
		}

		tests := []returnTest{
			{"return 10;", 10},
			{"return 10;", 10},
			{"return 10; 9;", 10},
			{"return 2 * 5; 9;", 10},
			{"9; return 2 * 5; 9;", 10},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			testIntegerObject(t, eval, tt.expected)
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

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != eval.NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}

	return true
}
