package token

const (
	INVALID TokenType = iota
	EOF

	ID
	INT

	// OPERATORS
	EQ
	PLUS
	MINUS
	BANG     // !
	ASTERISK // *
	SLASH    // /

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
)
