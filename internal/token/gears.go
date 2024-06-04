package token

var keywords = map[string]TokenType{
	"fn":  FUNC,
	"let": LET,
}

func LookupID(id string) TokenType {
	if tokenType, ok := keywords[id]; ok {
		return tokenType
	}

	return ID
}
