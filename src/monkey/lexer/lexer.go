package lexer

import (
	"monkey/token"
	"strings"
)

type Lexer struct {
	tokens []token.Token
}

func symbolTokens() []token.Token {
	symbolTokenTypes := []token.TokenType{
		token.PLUS,
		token.MINUS,
		token.BANG,
		token.ASTERISK,
		token.SLASH,

		token.LT,
		token.GT,

		token.EQ,
		token.NOT_EQ,

		token.ASSIGN,

		token.COMMA,
		token.SEMICOLON,

		token.LPAREN,
		token.RPAREN,
		token.LBRACE,
		token.RBRACE,
	}

	ret := []token.Token{}

	for _, tokenType := range symbolTokenTypes {
		ret = append(ret, token.Token{Type: tokenType, Literal: string(tokenType)})
	}

	return ret
}

func keywordTokens() []token.Token {
	return []token.Token{
		{Type: token.RETURN, Literal: "return"},
		{Type: token.FUNCTION, Literal: "function"},
		{Type: token.IF, Literal: "if"},
		{Type: token.ELSE, Literal: "else"},
		{Type: token.LET, Literal: "let"},
		{Type: token.TRUE, Literal: "true"},
		{Type: token.FALSE, Literal: "false"},
	}
}

func New(input string) []token.Token {
	i := 0
	length := len(input)
	tokens := []token.Token{}

	for i < length {
		foundToken := findToken(input[i:])
		if foundToken == nil {
			i += 1
			continue
		}
		if foundToken.Type == token.ILLEGAL {
			return []token.Token{}
		}

		println(foundToken.Literal)
		i += len(foundToken.Literal)
		tokens = append(tokens, *foundToken)
	}

	tokens = append(tokens, token.Token{Type: token.EOF, Literal: ""})

	return tokens
}

func isNumeral(c byte) bool {
	return '0' <= c && c <= '9'
}

func isAlphabet(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isIdentifierHeadCharacter(c byte) bool {
	return isAlphabet(c) || c == '_' || c == '$'
}

func isIdentifierCharacter(c byte) bool {
	return isIdentifierHeadCharacter(c) || isNumeral(c)
}

func isDelimiter(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func findToken(input string) *token.Token {
	// find token in current input.

	if isDelimiter(input[0]) {
		return nil
	}

	// first, find symbol token.
	symbolTokenCandidates := symbolTokens()
	for _, tokenCandidate := range symbolTokenCandidates {
		if equalToken(input, tokenCandidate.Literal) {
			return &token.Token{Type: tokenCandidate.Type, Literal: tokenCandidate.Literal}
		}
	}

	// if not exists, then find symbol number.
	if input[0] == '0' {
		return &token.Token{Type: token.INT, Literal: "0"}
	}

	// In here, input[0] != '0' is true.
	i := 0
	for i < len(input) && isNumeral(input[i]) {
		i++
	}

	// found numeral.
	if i > 0 {
		return &token.Token{Type: token.INT, Literal: input[:i]}
	}

	// Finally, find identifier.
	if isIdentifierHeadCharacter(input[0]) {
		i = 0
		for i < len(input) && isIdentifierCharacter(input[i]) {
			i++
		}

		for _, keyword := range keywordTokens() {
			if strings.Compare(input[:i], keyword.Literal) == 0 {
				return &token.Token{Type: keyword.Type, Literal: keyword.Literal}
			}
		}

		return &token.Token{Type: token.IDENT, Literal: input[:i]}
	}

	return &token.Token{Type: token.ILLEGAL, Literal: ""}
}

func equalToken(input string, tokenLiteral string) bool {
	i := 0
	for i < len(tokenLiteral) && i < len(input) {
		if input[i] != tokenLiteral[i] {
			return false
		}
		i++
	}

	// tokenLiteralとinputを協調して読みだしていたが、inputの読み取りが先に終端に達してしまった
	if i < len(tokenLiteral) {
		return false
	}

	return true
}
