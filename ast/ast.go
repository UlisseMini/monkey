package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string

}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Root node of every AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return "<empty program>"
	}
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value  string
}

// Identifiers are expressions! eg. (x + y)
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string{
	return l.Token.Literal
}


