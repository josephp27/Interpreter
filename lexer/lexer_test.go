package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"../token"
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

	for i, tt := range tests {

		nextToken := actual.NextToken()

		assert.Equal(nextToken.Type, tt.expectedType, "not equal")
		assert.Equal(nextToken.Literal, tt.expectedLiteral, "not equal")
	}
}
