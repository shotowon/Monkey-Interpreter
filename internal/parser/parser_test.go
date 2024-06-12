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
		input := `
		let foo = 5;
		let bar = 10;
		let baz = 32131;
		`

		l := lexer.New(input)
		p := parser.New(l)

		program := p.ParseProgram()
		if program == nil {
			t.Fatalf("p.ParseProgram() returned nil")
		}

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements do not contain 3 statements. got=%d", len(program.Statements))
		}

		tests := []struct {
			expectedID string
		}{
			{"foo"},
			{"bar"},
			{"baz"},
		}

		for i, test := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, test.expectedID) {
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
				intValue int64
			}

			prefixTests := []prefixTest{
				{"!15", "!", 15},
				{"-15", "-", 15},
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

				if !testIntegerLiteral(t, exp.Right, pTest.intValue) {
					return
				}
			}
		})
		t.Run("Parsing infix expressions", func(t *testing.T) {
			type infixTest struct {
				input    string
				left     int64
				operator string
				right    int64
			}

			infixTests := []infixTest{
				{"5 + 5;", 5, "+", 5},
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

				if !testIntegerLiteral(t, infix.Left, iTest.left) {
					return
				}

				if infix.Operator != iTest.operator {
					t.Fatalf("infix.Operator is not %s. got=%s", iTest.operator, infix.Operator)
				}

				if !testIntegerLiteral(t, infix.Right, iTest.right) {
					return
				}
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
