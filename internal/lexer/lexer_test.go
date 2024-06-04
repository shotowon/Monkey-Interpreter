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

	t.Run("Lexing numbers, identifiers and keywords", func(t *testing.T) {
		input := `let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;	
	};

	let result = add(five, ten);`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.LET, "let"},
			{token.ID, "five"},
			{token.EQ, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.ID, "ten"},
			{token.EQ, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.ID, "add"},
			{token.EQ, "="},
			{token.FUNC, "fn"},
			{token.LPAREN, "("},
			{token.ID, "x"},
			{token.COMMA, ","},
			{token.ID, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.ID, "x"},
			{token.PLUS, "+"},
			{token.ID, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},

			{token.LET, "let"},
			{token.ID, "result"},
			{token.EQ, "="},
			{token.ID, "add"},
			{token.LPAREN, "("},
			{token.ID, "five"},
			{token.COMMA, ","},
			{token.ID, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
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
