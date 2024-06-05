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
