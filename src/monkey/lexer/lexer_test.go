package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;\n123=====  ab0031_cd;_123;ifif;if;returnreturn;return;"

	expects := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.INT, "123"},
		{token.EQ, "=="},
		{token.EQ, "=="},
		{token.ASSIGN, "="},
		{token.IDENT, "ab0031_cd"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "_123"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ifif"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "returnreturn"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// l := New
	tokens := New(input)
	i := 0

	if len(tokens) != len(expects) {
		t.Errorf("len(tokens) is not %d. got=%d", len(expects), len(tokens))
		return
	}

	for i < len(expects) {
		expected := expects[i]
		if tokens[i].Literal != expected.Literal {
			t.Errorf("tokens[%d].Literal is not %s. got %s", i, expected.Literal, tokens[i].Literal)
			return
		}

		if tokens[i].Type != expected.Type {
			t.Errorf("tokens[%d].Type is not %s. got %s", i, expected.Type, tokens[i].Type)
			return
		}
		i++
	}
}
