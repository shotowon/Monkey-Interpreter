package token

const (
	INVALID TokenType = iota
	EOF

	ID
	INT

	EQ
	PLUS
	MINUS
	BANG     // "!"
	ASTERISK // "*"
	SLASH    // "/"

	LT // "<"
	GT // ">"

	COMMA
	SEMICOLON

	LPAREN // "("
	RPAREN // ")"
	LBRACE // "{"
	RBRACE // "}"

	FUNC
	LET
)
