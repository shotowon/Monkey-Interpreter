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

			// eval program
			{
				`if (10 > 1) {
					if (10 > 1) {
						return 10;
					}
					129
					return 1;
				}`,
				10,
			},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)
			testIntegerObject(t, eval, tt.expected)
		}
	})
	t.Run("test function object", func(t *testing.T) {
		input := `fn(x) { x + 2; };`

		eval := testEval(input)
		fn, ok := eval.(*object.Function)
		if !ok {
			t.Fatalf("eval is not *object.Function. got=%T (%+v)", eval, eval)
		}

		if len(fn.Params) != 1 {
			t.Fatalf("function has wrong number of params. got=%d expected=%d", len(fn.Params), 1)
		}

		if fn.Params[0].String() != "x" {
			t.Fatalf("params is not 'x'. got=%q", fn.Params[0])
		}

		expectedBody := `{ (x + 2) }`

		if fn.Body.String() != expectedBody {
			t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
		}
	})
	t.Run("test function call eval", func(t *testing.T) {
		type callTest struct {
			input    string
			expected int64
		}

		tests := []callTest{
			{"let identity = fn(x) { x; }; identity(5);", 5},
			{"let identity = fn(x) { return x; }; identity(5);", 5},
			{"let double = fn(x) { x * 2; }; double(5);", 10},
			{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
			{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
			{"fn(x) { x; }(5)", 5},
		}

		for _, tt := range tests {
			testIntegerObject(t, testEval(tt.input), tt.expected)
		}
	})
	t.Run("test closures", func(t *testing.T) {
		input := `
		let newAdder = fn(x) {
			fn(y) { x + y };
		};
		let addTwo = newAdder(2);
		addTwo(2);`
		testIntegerObject(t, testEval(input), 4)
	})
	t.Run("test error handling", func(t *testing.T) {
		type errTest struct {
			input           string
			expectedMessage string
		}

		tests := []errTest{
			{
				"5 + true;",
				"type mismatch: INTEGER + BOOL",
			},
			{
				"5 + true; 5;",
				"type mismatch: INTEGER + BOOL",
			},
			{
				"-true",
				"unknown operator: -BOOL",
			},
			{
				"true + false;",
				"unknown operator: BOOL + BOOL",
			},
			{
				"5; true + false; 5",
				"unknown operator: BOOL + BOOL",
			},
			{
				"if (10 > 1) { true + false; }",
				"unknown operator: BOOL + BOOL",
			},
			{
				`
				132
				if (10 > 1) {
					if (10 > 1) {
						return true + false;
					}
					return 1;
				}
				`,
				"unknown operator: BOOL + BOOL",
			},
			{
				"foobar",
				"identifier not found: foobar",
			},
		}

		for _, tt := range tests {
			eval := testEval(tt.input)

			errObj, ok := eval.(*object.Error)
			if !ok {
				t.Errorf("no error object returned. got=%T(%+v)", eval, eval)
				continue
			}

			if errObj.Message != tt.expectedMessage {
				t.Errorf("wrong error message. got=%q expected=%q", errObj.Message, tt.expectedMessage)
			}
		}
	})

	t.Run("test let statement", func(t *testing.T) {
		type letTest struct {
			input    string
			expected int64
		}

		tests := []letTest{
			{"let a = 5; a;", 5},
			{"let a = 5 * 5; a;", 25},
			{"let a = 5; let b = a; b;", 5},
			{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		}

		for _, tt := range tests {
			testIntegerObject(t, testEval(tt.input), tt.expected)
		}
	})

	t.Run("test string literal", func(t *testing.T) {
		input := `"Hello World"`

		eval := testEval(input)
		str, ok := eval.(*object.String)
		if !ok {
			t.Fatalf("eval object is not *object.String. got=%T", eval)
		}

		if str.Value != "Hello World" {
			t.Errorf("string has wrong value. got=%q", str.Value)
		}
	})
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnv()

	return eval.Eval(program, env)
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
