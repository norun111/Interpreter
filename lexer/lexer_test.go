package lexer

import (
	token2 "goInterpreter/token"
	"testing"
)

func TextNextToken(t *testing.T) {
	input := `let five = 5;

              let ten = 10;

 		      let add = fn(x, y) {
		      	x + y;
			  };
			  let result = add(five, ten)
			 `

	tests := []struct {
		expectedType    token2.TokenType
		expectedLiteral string
	}{
		{token2.LET, "let"},
		{token2.IDENT, "five"},
		{token2.ASSIGN, "="},
		{token2.INT, "5"},
		{token2.ASSIGN, "="},
		{token2.PLUS, "+"},
		{token2.LPAREN, "("},
		{token2.RPAREN, ")"},
		{token2.LBRACE, "{"},
		{token2.RBRACE, "}"},
		{token2.COMMA, ","},
		{token2.SEMICOLON, ";"},
		{token2.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
