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

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	for unicode.IsSpace(l.ch) {
		l.readChar()
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = token.New(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			t = token.New(token.BANG, l.ch)
		}
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
	case '[':
		t = token.New(token.LBRACKET, l.ch)
	case ']':
		t = token.New(token.RBRACKET, l.ch)
	case ',':
		t = token.New(token.COMMA, l.ch)
	case '+':
		t = token.New(token.PLUS, l.ch)
	case '-':
		t = token.New(token.MINUS, l.ch)
	case '/':
		t = token.New(token.SLASH, l.ch)
	case '*':
		t = token.New(token.ASTERISK, l.ch)
	case '>':
		t = token.New(token.GT, l.ch)
	case '<':
		t = token.New(token.LT, l.ch)
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	default:
		switch {
		case unicode.IsLetter(l.ch):
			t.Literal = l.readID()
			t.Type = token.LookupID(t.Literal)
		case unicode.IsDigit(l.ch):
			t.Literal = l.readNumber()
			t.Type = token.INT
		default:
			t = token.New(token.INVALID, l.ch)
		}
		return t
	}

	l.readChar()
	return t
}

func (l *Lexer) readID() string {
	pos := l.pos

	if unicode.IsLetter(l.ch) {
		l.readChar()
	}

	for unicode.IsDigit(l.ch) || unicode.IsLetter(l.ch) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
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

func (l *Lexer) readNumber() string {
	pos := l.pos

	for unicode.IsDigit(l.ch) {
		l.readChar()
	}

	return string(l.input[pos:l.pos])
}

func (l *Lexer) readString() string {
	pos := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return string(l.input[pos:l.pos])
}

func (l *Lexer) peekChar() rune {
	if l.readPos >= len(l.input) {
		return 0
	}

	return l.input[l.readPos]
}
