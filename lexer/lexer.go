package lexer

import (
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input string

	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		fmt.Println("def")
		if isLetter(l.char) {

			tok.Literal = l.readIdentifier()
			fmt.Printf("litera: %s", tok.Literal)
			tok.Type = token.LookupIdentifier(tok.Literal)
			// do I need tok.Tyoe=token.IDENT
			return tok
		} else if isNumber(l.char) {
			tok.Literal = l.readIdentifier()
			fmt.Printf("litera: %s", tok.Literal)
			tok.Type = token.INT
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tok

}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isWhitespaceIsh(char string) bool {
	whitespaceIshChars := []string{" ", "\t", "\n", "\r"}

	for _, c := range whitespaceIshChars {
		if c == char {
			return true
		}
	}
	return false

}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: t, Literal: string(ch),
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && 9 <= ch
}
