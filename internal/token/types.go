package token

const (
	INVALID TokenType = iota
	EOF

	ID
	INT

	EQ
	PLUS
	MINUS

	COMMA
	SEMICOLON

	LPAREN // "("
	RPAREN // ")"
	LBRACE // "{"
	RBRACE // "}"

	FUNC
	LET
)
