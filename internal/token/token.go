package token

type TokenType uint

type Token struct {
	Type    TokenType
	Literal string
}

func New(t TokenType, l rune) Token {
	return Token{t, string(l)}
}
