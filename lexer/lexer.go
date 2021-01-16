package lexer

import "monkey/token"


type Lexer struct  {
	input string // the string we are lexing
	position int // our current position
	readPosition int // the reading position, used for peek
	ch byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}


func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.MUL, l.ch)
	case '/':
		tok = newToken(token.DIV, l.ch)
	// TODO: Abstract != and == branches (they are basically the same)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NE, Literal: literal}
		} else {
			tok = newToken(token.NOT, l.ch)
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// TODO: Abstract to l.readTill(pred func(Lexer) bool)
		// with backup so I can avoid early returns.
		if isLetter(l.ch) {
			// This is wrong, because it might be a keyword!
			// tok.Type = token.IDENT
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			// early return because we call readChar once after finishing
			// the token in readIdentifier().
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readInt()
			tok.Type = token.INT
			return tok // same early return
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readInt() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // nullbyte
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()

	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isDigit(ch byte) bool {
	return ch <= '9' && ch >= '0'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch  =='\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
