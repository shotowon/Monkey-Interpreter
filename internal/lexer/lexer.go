package lexer

import (
	"monkey/internal/token"
	"unicode"
)

type Lexer struct {
	input   []rune
	pos     int
	readPos int
	ch      rune
}

func New(input string) *Lexer {
	l := &Lexer{
		input: []rune(input),
	}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case '=':
		t = token.New(token.EQ, l.ch)
	case ';':
		t = token.New(token.SEMICOLON, l.ch)
	case '(':
		t = token.New(token.LPAREN, l.ch)
	case ')':
		t = token.New(token.RPAREN, l.ch)
	case '{':
		t = token.New(token.LBRACE, l.ch)
	case '}':
		t = token.New(token.RBRACE, l.ch)
	case ',':
		t = token.New(token.COMMA, l.ch)
	case '+':
		t = token.New(token.PLUS, l.ch)
	case '-':
		t = token.New(token.MINUS, l.ch)
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	}

	l.readChar()
	return t
}

func (l *Lexer) readID() string {
	pos := l.pos

	if unicode.IsLetter(l.ch) {
		l.readChar()
	}

	for unicode.IsNumber(l.ch) || unicode.IsLetter(l.ch) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
}
