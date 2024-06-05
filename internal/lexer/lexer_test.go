package lexer_test

import (
	"monkey/internal/lexer"
	"monkey/internal/token"
	"testing"
)

type nextTokenExpectedValue struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []nextTokenExpectedValue
	}{
		{
			name:  "Lexing initial set of tokens",
			input: `=+(){},;`,
			expected: []nextTokenExpectedValue{
				{token.EQ, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
			},
		},
		{
			name: "Lexing numbers, identifiers and keywords",
			input: `let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;	
			};

			let result = add(five, ten);`,
			expected: []nextTokenExpectedValue{
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
			},
		},
		{
			name: "Lexing extended operator set",
			input: `let five = 5;
			let ten = 10;
			let add = fn(x, y) {
				x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;
			5 <= 5 >= 5;
			`,
			expected: []nextTokenExpectedValue{
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

				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},

				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.EQ, "="},
				{token.INT, "5"},
				{token.GT, ">"},
				{token.EQ, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "Lexing extended keyword set",
			input: `let five = 5;
			let ten = 10;
			let add = fn(x, y) {
				x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;
			5 <= 5 >= 5;
			if (5 < 10) {
				return true;
			} else {
				return false;
			}`,
			expected: []nextTokenExpectedValue{
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

				{token.BANG, "!"},
				{token.MINUS, "-"},
				{token.SLASH, "/"},
				{token.ASTERISK, "*"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},

				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.GT, ">"},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.EQ, "="},
				{token.INT, "5"},
				{token.GT, ">"},
				{token.EQ, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				/*
					if (5 < 10) {
						return true;
					} else {
						return false;
					}`,
				*/
				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.EOF, ""},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := lexer.New(test.input)

			for i, expected := range test.expected {
				tok := l.NextToken()

				if tok.Type != expected.expectedType {
					t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
						i, expected.expectedType, tok.Type)
				}

				if tok.Literal != expected.expectedLiteral {
					t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
						i, expected.expectedLiteral, tok.Literal)
				}
			}
		})
	}
}
