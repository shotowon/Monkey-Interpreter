package parser_test

import (
	"fmt"
	"monkey/internal/ast"
	"monkey/internal/lexer"
	"monkey/internal/parser"
	"monkey/internal/token"
	"testing"
)

func TestStatementsParsing(t *testing.T) {
	t.Run("Return statement", func(t *testing.T) {
		input := `
		return 5;
		return 10;
		return 993322;
		`

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
		}

		for _, stmt := range program.Statements {
			returnStmt, ok := stmt.(*ast.ReturnStatement)
			if !ok {
				t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
				continue
			}

			if returnStmt.TokenLiteral() != "return" {
				t.Errorf("returnStmt.TokenLiteral() not 'return', got=%q", returnStmt.TokenLiteral())
			}
		}
	})

	t.Run("Let statement", func(t *testing.T) {
		type letStatementTesting struct {
			input         string
			expectedID    string
			expectedValue interface{}
		}

		tests := []letStatementTesting{
			{"let foo = 5;", "foo", 5},
			{"let bar = true;", "bar", true},
			{"let baz = y;", "baz", "y"},
		}

		for _, tt := range tests {
			l := lexer.New(tt.input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
			}

			stmt := program.Statements[0]
			if !testLetStatement(t, stmt, tt.expectedID) {
				return
			}

			val := stmt.(*ast.LetStatement).Value
			if !testLiteralExpr(t, val, tt.expectedValue) {
				return
			}
		}
	})

	t.Run("Expressions", func(t *testing.T) {
		t.Run("Test identifier expression parsing", func(t *testing.T) {
			input := `foobar;`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
			}

			id, ok := stmt.Expression.(*ast.ID)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.ID. got=%T", stmt.Expression)
			}

			if id.Value != "foobar" {
				t.Errorf("id.Value not %s. got=%s", "foobar", id.Value)
			}

			if id.TokenLiteral() != "foobar" {
				t.Errorf("id.TokenLiteral() not %s. got=%s", "foobar", id.Value)
			}
		})

		t.Run("Integer literal expression parsing", func(t *testing.T) {
			input := `5;`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
			}

			literal, ok := stmt.Expression.(*ast.IntegerLiteral)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.IntegerLiteral. got=%T", program.Statements[0])
			}

			if literal.Value != 5 {
				t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
			}

			if literal.TokenLiteral() != "5" {
				t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
			}
		})
		t.Run("Prefix expression parsing", func(t *testing.T) {
			type prefixTest struct {
				input    string
				operator string
				value    interface{}
			}

			prefixTests := []prefixTest{
				{"!15", "!", 15},
				{"-15", "-", 15},
				{"!true", "!", true},
				{"-false", "-", false},
			}

			for _, pTest := range prefixTests {
				l := lexer.New(pTest.input)
				p := parser.New(l)
				program := p.ParseProgram()
				checkParserErrors(t, p)

				if len(program.Statements) != 1 {
					t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
				}

				stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
				if !ok {
					t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement statements. got=%T", program.Statements[0])
				}

				exp, ok := stmt.Expression.(*ast.PrefixExpression)
				if !ok {
					t.Fatalf("stmt.Expression is not *ast.PrefixExpression statements. got=%T", stmt.Expression)
				}

				if exp.Operator != pTest.operator {
					t.Fatalf("exp.Operator is not %s statements. got=%T", pTest.operator, exp.Operator)
				}

				if !testLiteralExpr(t, exp.Right, pTest.value) {
					return
				}
			}
		})
		t.Run("Parsing infix expressions", func(t *testing.T) {
			type infixTest struct {
				input    string
				left     interface{}
				operator string
				right    interface{}
			}

			infixTests := []infixTest{
				{"5 + 5;", 5, "+", 5},
				{"5 - 5;", 5, "-", 5},
				{"5 * 5;", 5, "*", 5},
				{"5 / 5;", 5, "/", 5},
				{"5 > 5;", 5, ">", 5},
				{"5 < 5;", 5, "<", 5},
				{"5 == 5;", 5, "==", 5},
				{"5 != 5;", 5, "!=", 5},
				{"true == true", true, "==", true},
				{"true != false", true, "!=", false},
				{"false == false", false, "==", false},
			}

			for _, iTest := range infixTests {
				l := lexer.New(iTest.input)
				p := parser.New(l)
				program := p.ParseProgram()
				checkParserErrors(t, p)

				if len(program.Statements) != 1 {
					t.Fatalf("program.Statements does not container %d statements. got=%d", 1, len(program.Statements))
				}

				stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
				if !ok {
					t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement")
				}

				infix, ok := stmt.Expression.(*ast.InfixExpression)
				if !ok {
					t.Fatalf("stmt.Expression is not *ast.InfixExpression")
				}

				if !testLiteralExpr(t, infix.Left, iTest.left) {
					return
				}

				if infix.Operator != iTest.operator {
					t.Fatalf("infix.Operator is not %s. got=%s", iTest.operator, infix.Operator)
				}

				if !testLiteralExpr(t, infix.Right, iTest.right) {
					return
				}
			}
		})

		t.Run("operator precedence testing", func(t *testing.T) {
			type precTest struct {
				input    string
				expected string
			}

			precedenceTests := []precTest{
				{
					"-a * b",
					"((-a) * b)",
				},
				{
					"!-a",
					"(!(-a))",
				},
				{
					"a + b + c",
					"((a + b) + c)",
				},
				{
					"a + b - c",
					"((a + b) - c)",
				},
				{
					"a * b * c",
					"((a * b) * c)",
				},
				{
					"a * b / c",
					"((a * b) / c)",
				},
				{
					"a + b / c",
					"(a + (b / c))",
				},
				{
					"a + b * c + d / e - f",
					"(((a + (b * c)) + (d / e)) - f)",
				},
				{
					"3 + 4; -5 * 5",
					"(3 + 4)((-5) * 5)",
				},
				{
					"5 > 4 == 3 < 4",
					"((5 > 4) == (3 < 4))",
				},
				{
					"5 < 4 != 3 > 4",
					"((5 < 4) != (3 > 4))",
				},
				{
					"3 + 4 * 5 == 3 * 1 + 4 * 5",
					"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
				},
				{
					"3 + 4 * 5 == 3 * 1 + 4 * 5",
					"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
				},
				{
					"true",
					"true",
				},
				{
					"false",
					"false",
				},
				{
					"3 < 4 == false",
					"((3 < 4) == false)",
				},
				{
					"1 + (2 + 3) + 4",
					"((1 + (2 + 3)) + 4)",
				},
				{
					"(5 + 5) * 2",
					"((5 + 5) * 2)",
				},
				{
					"2 / (5 + 5)",
					"(2 / (5 + 5))",
				},
				{
					"-(5 + 5)",
					"(-(5 + 5))",
				},
				{
					"!(true == true)",
					"(!(true == true))",
				},
				{
					"a + add(b * c) + d",
					"((a + add((b * c))) + d)",
				},
				{
					"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
					"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
				},
				{
					"add(a + b + c * d / f + g)",
					"add((((a + b) + ((c * d) / f)) + g))",
				},
			}

			for _, pTest := range precedenceTests {
				l := lexer.New(pTest.input)
				p := parser.New(l)
				program := p.ParseProgram()
				checkParserErrors(t, p)

				if actual := program.String(); actual != pTest.expected {
					t.Errorf("expected=%q, got=%q", pTest.expected, actual)
				}
			}
		})
		t.Run("Boolean literal parsing", func(t *testing.T) {
			type boolTest struct {
				input    string
				expected bool
			}
			tests := []boolTest{
				{input: "true;", expected: true},
				{input: "false;", expected: false},
			}

			for _, test := range tests {
				l := lexer.New(test.input)
				p := parser.New(l)
				program := p.ParseProgram()
				checkParserErrors(t, p)

				if len(program.Statements) != 1 {
					t.Fatalf("program.Statements does not contain %d statement. got=%d", 1, len(program.Statements))
				}

				stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
				if !ok {
					t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
				}

				if !testBoolExpr(t, stmt.Expression, test.expected) {
					return
				}
			}
		})
		t.Run("If expression parsing", func(t *testing.T) {
			input := `if (x > y) { x }`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not containt %d statements. got=%d", 1, len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not *ast.Expression. got=%T", program.Statements[0])
			}

			expr, ok := stmt.Expression.(*ast.IfExpression)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.IfExpression. got=%T", stmt.Expression)
			}

			if !testInfixExpr(t, expr.Condition, "x", ">", "y") {
				return
			}

			if len(expr.Consequence.Statements) != 1 {
				t.Errorf("expr.Consequence.Statements does not contain %d statements. got=%d", 1, len(expr.Consequence.Statements))
			}

			consequence, ok := expr.Consequence.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("expr.Consequence.Statements[0] is not *ast.ExpressionStatement. got=%T", expr.Consequence.Statements[0])
			}

			if !testID(t, consequence.Expression, "x") {
				return
			}

			if expr.Alternative != nil {
				t.Errorf("expr.Alternative is not nil. got=%+v", expr.Alternative)
			}
		})
		t.Run("Test function literal parsing", func(t *testing.T) {
			input := `fn(x, y) { x + y; }`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain %d statements. got=%d",
					1, len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T",
					program.Statements[0])
			}

			function, ok := stmt.Expression.(*ast.FunctionLiteral)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.FunctionLiteral. got=%T",
					stmt.Expression)
			}

			if len(function.Params) != 2 {
				t.Fatalf("function.Params does not contain %d parameters. got=%d",
					2, len(function.Params))
			}

			testLiteralExpr(t, function.Params[0], "x")
			testLiteralExpr(t, function.Params[1], "y")

			if len(function.Body.Statements) != 1 {
				t.Fatalf("function.Body.Statements does not contain %d statements. got=%d",
					1, len(function.Body.Statements))
			}

			bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("function.Body.Statements[0] is not *ast.ExpressionStatement. got=%T",
					function.Body.Statements[0])
			}

			testInfixExpr(t, bodyStmt.Expression, "x", "+", "y")
		})
		t.Run("Test call expression parsing", func(t *testing.T) {
			input := `add(1, 2 * 3, 4 + 5);`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			if len(program.Statements) != 1 {
				t.Fatalf("program.Statements does not contain %d. got=%d",
					1, len(program.Statements))
			}

			stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
			if !ok {
				t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement. got=%T", program.Statements[0])
			}

			expr, ok := stmt.Expression.(*ast.CallExpression)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.CallExpression. got=%T", stmt.Expression)
			}

			if !testID(t, expr.Function, "add") {
				return
			}

			if len(expr.Arguments) != 3 {
				t.Fatalf("wrong length of arguments. got=%d", len(expr.Arguments))
			}

			testLiteralExpr(t, expr.Arguments[0], 1)
			testInfixExpr(t, expr.Arguments[1], 2, "*", 3)
			testInfixExpr(t, expr.Arguments[2], 4, "+", 5)
		})
		t.Run("test string literal expression parsing", func(t *testing.T) {
			input := `"hello world";`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			literal, ok := stmt.Expression.(*ast.StringLiteral)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.StringLiteral. got=%T", stmt.Expression)
			}

			if literal.Value != "hello world" {
				t.Errorf("literal.Value is not %q. got=%q", "hello world", literal.Value)
			}
		})
		t.Run("test array literal parsing", func(t *testing.T) {
			input := "[1, 2 * 3, 3 + 3]"

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			literal, ok := stmt.Expression.(*ast.ArrayLiteral)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.ArrayLiteral. got=%T", stmt.Expression)
			}

			if len(literal.Elements) != 3 {
				t.Fatalf("len(literal.Elements) not 3. got=%d", len(literal.Elements))
			}

			testIntegerLiteral(t, literal.Elements[0], 1)
			testInfixExpr(t, literal.Elements[1], 2, "*", 3)
			testInfixExpr(t, literal.Elements[2], 3, "+", 3)
		})
		t.Run("test index expressions", func(t *testing.T) {
			input := `arr[1 + 1]`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			index, ok := stmt.Expression.(*ast.IndexExpression)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.IndexExpression. got=%T", stmt.Expression)
			}

			if !testID(t, index.Left, "arr") {
				return
			}

			if !testInfixExpr(t, index.Index, 1, "+", 1) {
				return
			}
		})
		t.Run("test hash map literal", func(t *testing.T) {
			input := `{"one": 1, "two": 2}`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			stmt := program.Statements[0].(*ast.ExpressionStatement)
			hm, ok := stmt.Expression.(*ast.HashMapLiteral)
			if !ok {
				t.Fatalf("stmt.Expression is not *ast.HashMapLiteral. got=%T", stmt.Expression)
			}

			if len(hm.Pairs) != 2 {
				t.Errorf("wrong number of key-value pairs not 2. got=%d", len(hm.Pairs))
			}

			expected := map[string]int{
				"one": 1,
				"two": 2,
			}

			for k, v := range hm.Pairs {
				literal, ok := k.(*ast.StringLiteral)
				if !ok {
					t.Errorf("k is not *ast.StringLiteral. got=%T", k)
				}

				testIntegerLiteral(t, v, int64(expected[literal.String()]))
			}
		})
		t.Run("empty hash-map literal", func(t *testing.T) {
			input := "{}"

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			expr := program.Statements[0].(*ast.ExpressionStatement)
			literal, ok := expr.Expression.(*ast.HashMapLiteral)
			if !ok {
				t.Fatalf("expr.Expression is not *ast.HashMapLiteral. got=%T", expr.Expression)
			}

			if len(literal.Pairs) != 0 {
				t.Errorf("wrong number of literal.Pairs. got=%d, expected=%d", len(literal.Pairs), 0)
			}
		})
		t.Run("test hash-map literal parsing with expressions", func(t *testing.T) {
			input := `{"one": 0 + 1, "two": 10 - 8}`

			l := lexer.New(input)
			p := parser.New(l)
			program := p.ParseProgram()
			checkParserErrors(t, p)

			expr := program.Statements[0].(*ast.ExpressionStatement)
			hashMap, ok := expr.Expression.(*ast.HashMapLiteral)
			if !ok {
				t.Fatalf("expr.Expression is not *ast.HashMapLiteral. got=%T", expr.Expression)
			}

			if len(hashMap.Pairs) != 2 {
				t.Errorf("wrong number of literal.Pairs. got=%d, expected=%d", len(hashMap.Pairs), 2)
			}

			pairsTest := map[string]func(ast.Expression){
				"one": func(e ast.Expression) {
					testInfixExpr(t, e, 0, "+", 1)
				},
				"two": func(e ast.Expression) {
					testInfixExpr(t, e, 10, "-", 8)
				},
			}

			for k, v := range hashMap.Pairs {
				literal, ok := k.(*ast.StringLiteral)
				if !ok {
					t.Errorf("k is not *ast.StringLiteral. got=%T", k)
					continue
				}

				testFunc, ok := pairsTest[literal.String()]
				if !ok {
					t.Errorf("no test function for key %q found", literal.String())
					continue
				}

				testFunc(v)
			}
		})
	})
}

func TestDebugging(t *testing.T) {
	t.Run("string function", func(t *testing.T) {
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.Token{Type: token.LET, Literal: "let"},
					Name: &ast.ID{
						Token: token.Token{Type: token.ID, Literal: "a"},
						Value: "a",
					},
					Value: &ast.ID{
						Token: token.Token{Type: token.ID, Literal: "b"},
						Value: "b",
					},
				},
			},
		}

		if program.String() != "let a = b;" {
			t.Errorf("program.String() is not correct. got=%q", program.String())
		}
	})
}

func TestFunctionParamParsing(t *testing.T) {
	type funcParamsParseTest struct {
		input          string
		expectedParams []string
	}

	tests := []funcParamsParseTest{
		{input: "fn() {};", expectedParams: []string{}},
		{input: "fn(x) {};", expectedParams: []string{"x"}},
		{input: "fn(x, y, z) {};", expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		fn := stmt.Expression.(*ast.FunctionLiteral)

		if len(tt.expectedParams) != len(fn.Params) {
			t.Errorf("fn.Params does not contain %d. got=%d",
				len(tt.expectedParams), len(fn.Params))
		}

		for i, id := range fn.Params {
			testLiteralExpr(t, id, tt.expectedParams[i])
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() is not let. got=%s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not %s. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() is not %s. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, errMsg := range errors {
		t.Errorf("parser error: %q", errMsg)
	}

	t.FailNow()
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	integ, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", exp)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func testID(t *testing.T, exp ast.Expression, value string) bool {
	id, ok := exp.(*ast.ID)
	if !ok {
		t.Errorf("exp is not *ast.ID. got=%T", exp)
		return false
	}

	if id.Value != value {
		t.Errorf("id.Value is not %s. got=%s", value, id.Value)
		return false
	}

	if id.TokenLiteral() != value {
		t.Errorf("id.TokenLiteral() is not %s. got=%s", value, id.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpr(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testID(t, exp, v)
	case bool:
		return testBoolExpr(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpr(t *testing.T, exp ast.Expression, left interface{}, op string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not *ast.InfixExpression. got=%T", exp)
		return false
	}

	if !testLiteralExpr(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != op {
		t.Errorf("opExp.Operator is not %s. got=%s", op, opExp.Operator)
		return false
	}

	if !testLiteralExpr(t, opExp.Right, right) {
		return false
	}

	return true
}

func testBoolExpr(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.BooleanExpression)
	if !ok {
		t.Errorf("exp is not an *ast.BooleanExpression")
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value is not %t. got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral() not %t. got=%s", value, bo.TokenLiteral())
		return false
	}

	return true
}
