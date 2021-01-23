package ast

import (
	"monkey/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var buf strings.Builder

	for _, s := range p.Statements {
		buf.WriteString(s.String());
	}

	return buf.String()

}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return "<empty program>"
	}
}

// Expressions are statements, eg f(x,y); or x+y; are valid.i
type ExpressionStatement struct {
	Token token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	if (es.Expression != nil) {
		return es.Expression.String()
	}
	return ""
}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Value  string
}

// Identifiers are expressions! eg. (x + y)
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string { return i.Value }

type LetStatement struct {
	Token token.Token // the token.LET token
	Name *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}
func (l *LetStatement) String() string {
	var buf strings.Builder

	buf.WriteString(l.TokenLiteral() + " ")
	buf.WriteString(l.Name.String())
	buf.WriteString(" = ")

	if l.Value != nil {
		buf.WriteString(l.Value.String())
	}

	buf.WriteString(";")
	return buf.String()
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	var buf strings.Builder
	buf.WriteString(r.TokenLiteral() + " ")

	if r.ReturnValue != nil {
		buf.WriteString(r.ReturnValue.String())
	}

	buf.WriteString(";");

	return buf.String()
}
