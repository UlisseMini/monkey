package lexer

import (
	"monkey/token"
	"testing"
)


func TestNextToken(t *testing.T) {
	// whitespace is insignificant in monkey.
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
	  x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 4;
	2 != 3;
	!(1 == 2);
	if (5 < 10) {
		return true; 
	} else {
		return false;
	}
	`


	tests := []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		// function assignment
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},

		// function params
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},

		// function body
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		// let result = add(five, ten);
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// !-/*5;
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// 5 < 10 > 4;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},

		// 2 != 3;
		{token.INT, "2"},
		{token.NE, "!="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},

		// !(1 == 2);
		{token.NOT, "!"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},

		// if (5 < 10)
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},

		// { return true; }
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// else { return false; }
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}


	l := New(input) // create lexer
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
