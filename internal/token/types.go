package token

const (
	INVALID TokenType = iota
	EOF

	ID
	INT

	// OPERATORS
	ASSIGN // =
	PLUS
	MINUS
	BANG     // !
	ASTERISK // *
	SLASH    // /

	// COMBINED OPERATORS
	EQ     // ==
	NOT_EQ // !=

	LT // <
	GT // >

	COMMA
	SEMICOLON

	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	// KEYWORDS
	FUNC   // fn
	LET    // let
	TRUE   // true
	FALSE  // false
	IF     // if
	ELSE   // else
	RETURN // return

	STRING // string
)
