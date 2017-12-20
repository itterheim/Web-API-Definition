package main

import (
	"bufio"
	"bytes"
	"io"
)

// Scanner represents lexical scanner
type Scanner struct {
	reader *bufio.Reader
}

// NewScanner creates new instance of scanner
func NewScanner(reader io.Reader) *Scanner {
	return &Scanner{
		reader: bufio.NewReader(reader),
	}
}

// Read next character
func (scanner *Scanner) Read() rune {
	ch, _, err := scanner.reader.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// Unread last character
func (scanner *Scanner) Unread() {
	scanner.reader.UnreadRune()
}

// Scan returns next token and literal value
func (scanner *Scanner) Scan() Token {
	character := scanner.Read()

	if isWhiteSpace(character) {
		scanner.Unread()
		return scanner.scanWhiteSpace()
	} else if character == '"' {
		scanner.Unread()
		return scanner.scanString()
	} else if character == '{' || character == '[' {
		scanner.Unread()
		return scanner.scanJSON()
	} else if isLetter(character) || character == '-' {
		scanner.Unread()
		return scanner.scanLetter()
	} else if isDigit(character) {
		scanner.Unread()
		return scanner.scanNumber()
	}

	if character == eof {
		return Token{TokenEndOfFile, ""}
	}

	return Token{TokenIllegal, string(character)}
}

func (scanner *Scanner) scanWhiteSpace() Token {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.Read())

	for {
		character := scanner.Read()

		if character == eof {
			break
		} else if !isWhiteSpace(character) {
			scanner.Unread()
			break
		} else {
			buffer.WriteRune(character)
		}
	}

	return Token{TokenWhiteSpace, buffer.String()}
}

func (scanner *Scanner) scanLetter() Token {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.Read())

	for {
		character := scanner.Read()

		if character == eof {
			break
		} else if !isLetter(character) && !isDigit(character) && character != '-' {
			scanner.Unread()
			break
		} else {
			buffer.WriteRune(character)
		}
	}

	str := buffer.String()

	switch str {
	case TokenGET:
		return Token{TokenGET, str}
	case TokenPOST:
		return Token{TokenPOST, str}
	case TokenUPDATE:
		return Token{TokenUPDATE, str}
	case TokenDELETE:
		return Token{TokenDELETE, str}
	case TokenName:
		return Token{TokenName, str}
	case TokenPort:
		return Token{TokenPort, str}
	case TokenResponse:
		return Token{TokenResponse, str}
	case TokenRoutePrefix:
		return Token{TokenRoutePrefix, str}
	}

	return Token{TokenIllegal, str}
}

func (scanner *Scanner) scanString() Token {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.Read())

	for {
		character := scanner.Read()

		if character == eof {
			buffer.WriteRune('"')
			break
		} else if character == '"' {
			buffer.WriteRune(character)
			break
		} else {
			buffer.WriteRune(character)
		}
	}

	str := buffer.String()
	str = str[1 : len(str)-1]

	return Token{TokenString, str}
}

func (scanner *Scanner) scanNumber() Token {
	var buffer bytes.Buffer
	buffer.WriteRune(scanner.Read())

	for {
		character := scanner.Read()

		if character == eof {
			break;
		} else if !isDigit(character) {
			scanner.Unread()
			break;
		} else {
			buffer.WriteRune(character)
		}
	}

	return Token{TokenNumber, buffer.String()}
}

func (scanner *Scanner) scanJSON() Token {
	var buffer bytes.Buffer

	opennedObject := 0
	opennedArray := 0

	for {
		character := scanner.Read()
		if character == eof {
			break
		} else if character == '{' {
			opennedObject++
		} else if character == '}' {
			opennedObject--
		} else if character == '[' {
			opennedArray++
		} else if character == ']' {
			opennedArray--
		} else {

		}
		buffer.WriteRune(character)

		if opennedArray == 0 && opennedObject == 0 {
			return Token{TokenJSON, buffer.String()}
		}
	}

	return Token{TokenIllegal, buffer.String()}
}
