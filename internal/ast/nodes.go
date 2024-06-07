package ast

import "monkey/internal/token"

type LetStatement struct {
	Token token.Token
	Name  *ID
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ID struct {
	Token token.Token
	Value string
}

func (i *ID) expressionNode()      {}
func (i *ID) TokenLiteral() string { return i.Token.Literal }
