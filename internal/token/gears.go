package token

var keywords = map[string]TokenType{
	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupID(id string) TokenType {
	if tokenType, ok := keywords[id]; ok {
		return tokenType
	}

	return ID
}

func (t TokenType) String() string {
	switch t {
	case INVALID:
		return "INVALID"
	case EOF:
		return "EOF"
	case ID:
		return "ID"
	case INT:
		return "INT"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case BANG:
		return "!"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case EQ:
		return "=="
	case NOT_EQ:
		return "!="
	case LT:
		return "<"
	case GT:
		return ">"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNC:
		return "fn"
	case LET:
		return "let"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case RETURN:
		return "return"
	default:
		return "UNKNOWN"
	}
}
