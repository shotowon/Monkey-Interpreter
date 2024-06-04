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

	return l
}
