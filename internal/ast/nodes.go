package ast

import "monkey/internal/token"

type LetStatement struct {
	Token token.Token
	Name  *ID
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ID struct {
	Token token.Token
	Value string
}

func (i *ID) expressionNode()      {}
func (i *ID) TokenLiteral() string { return i.Token.Literal }
