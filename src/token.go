package main

import (
	"fmt"
)

var eof = rune(0)

const (
	TokenIllegal    = "ILLEGAL"
	TokenEndOfFile  = "EOF"
	TokenWhiteSpace = "WHITESPACE"

	TokenString = "STRING"
	TokenJSON   = "JSON"

	TokenGET    = "GET"
	TokenPOST   = "POST"
	TokenUPDATE = "UPDATE"
	TokenDELETE = "DELETE"

	TokenName        = "-name"
	TokenPort        = "-port"
	TokenResponse    = "-response"
	TokenRoutePrefix = "-routePrefix"
)

var (
	RequestTokens = map[string]bool{
		TokenGET:    true,
		TokenPOST:   true,
		TokenUPDATE: true,
		TokenDELETE: true,
	}

	ParameterTokens = map[string]bool{
		TokenName:        true,
		TokenPort:        true,
		TokenResponse:    true,
		TokenRoutePrefix: true,
	}
)

// Token is language element
type Token struct {
	key   string
	value string
}

// String returns token as string
func (token Token) String() string {
	if len(token.value) > 10 {
		return fmt.Sprintf("%v = %.10v...", token.key, token.value)
	}
	return fmt.Sprintf("%v = %v", token.key, token.value)
}

// Command is language command
type Command struct {
	key   string
	value *Token
}

// String returns command as string
func (command Command) String() string {
	value := "-"
	if command.value != nil {
		value = command.value.String()
	}
	return fmt.Sprintf("%v: %v", command.key, value)
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
