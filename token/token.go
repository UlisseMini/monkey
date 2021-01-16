package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 13401

	// Operators
	ASSIGN = "="
	GT = ">"
	LT = "<"
	EQ = "=="
	NE = "!="
	NOT = "!"

	PLUS = "+"
	MINUS = "-"
	MUL = "*"
	DIV = "/"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	FALSE = "FALSE"
	TRUE = "TRUE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"false": FALSE,
	"true": TRUE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if typ, ok := keywords[ident]; ok {
		return typ
	}
	return IDENT
}
