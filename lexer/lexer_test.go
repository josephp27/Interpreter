package lexer

import (
	"github.com/stretchr/testify/assert"
	"monkeylang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	actual := New(input)

	for _, tt := range tests {

		tok := actual.NextToken()

		assert.Equal(t, tok.Type, tt.expectedType, "not equal")
		assert.Equal(t, tok.Literal, tt.expectedLiteral, "not equal")
	}
}
