package parser_test

import (
	"monkey/internal/ast"
	"monkey/internal/lexer"
	"monkey/internal/parser"
	"monkey/internal/token"
	"testing"
)

func TestReturnStatement(t *testing.T) {
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
}

func TestLetStatements(t *testing.T) {
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
}

func TestString(t *testing.T) {
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
}

func TestID(t *testing.T) {
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
