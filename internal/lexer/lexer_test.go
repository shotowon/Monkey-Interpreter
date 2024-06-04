package lexer_test

import (
	"monkey/internal/lexer"
	"monkey/internal/token"
	"testing"
)

func TestNextToken1(t *testing.T) {
	t.Run("Lexing initial set of tokens", func(t *testing.T) {
		input := `=+(){},;`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.EQ, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
		}

		l := lexer.New(input)

		for i, test := range tests {
			tok := l.NextToken()

			if tok.Type != test.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, test.expectedType, tok.Type)
			}

			if tok.Literal != test.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
					i, test.expectedLiteral, tok.Literal)
			}
		}
	})
}
