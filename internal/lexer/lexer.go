package lexer

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
