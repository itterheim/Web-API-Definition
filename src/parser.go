package main

import (
	"errors"
	"io"
)

// Parser represent a parser
type Parser struct {
	scanner     *Scanner
	tokenBuffer struct {
		token  Token
		stored bool
	}
	commandBuffer struct {
		command Command
		stored  bool
	}
}

// NewParser creates new instance of parser
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		scanner: NewScanner(reader),
	}
}

// Scan for next token
func (parser *Parser) Scan() Token {
	if parser.tokenBuffer.stored {
		parser.tokenBuffer.stored = false
		return parser.tokenBuffer.token
	}

	token := parser.scanner.Scan()
	parser.tokenBuffer.token = token

	return token
}

// Unscan returns scanner one step back
func (parser *Parser) Unscan() {
	parser.tokenBuffer.stored = true
}

// ScanIgnoreWhiteSpace is scanning without whitespace
func (parser *Parser) ScanIgnoreWhiteSpace() Token {
	token := parser.Scan()
	if token.key == TokenWhiteSpace {
		token = parser.Scan()
	}

	return token
}

// ScanCommand return prased list of commands
func (parser *Parser) ScanCommand() Command {
	if parser.commandBuffer.stored {
		parser.commandBuffer.stored = false
		return parser.commandBuffer.command
	}

	commandToken := parser.ScanIgnoreWhiteSpace()
	valueToken := parser.ScanIgnoreWhiteSpace()

	if commandToken.key == TokenEndOfFile {
		return Command{TokenEndOfFile, nil}
	}

	command := Command{}
	if ParameterTokens[commandToken.key] || RequestTokens[commandToken.key] {
		command.key = commandToken.key
	} else {
		panic(errors.New("Invalid command " + commandToken.key + ": " + commandToken.value))
	}

	command.value = &valueToken

	parser.commandBuffer.command = command

	return command
}

// UnscanCommand returns command scanner one step back
func (parser *Parser) UnscanCommand() {
	parser.commandBuffer.stored = true
}

// GetApp returns application definiton object
func (parser *Parser) GetApp() App {
	app := App{}
	var endpoint *Endpoint

	command := parser.ScanCommand()
	for command.key != TokenEndOfFile {

		if endpoint == nil && ParameterTokens[command.key] {
			switch command.key {
			case TokenRoutePrefix:
				app.Options.RoutePrefix = command.value.value
			case TokenName:
				app.Name = command.value.value
			case TokenPort:
				app.Options.Port = command.value.value
			}

			command = parser.ScanCommand()
			continue
		}

		if endpoint != nil && RequestTokens[command.key] {
			app.Endpoints = append(app.Endpoints, *endpoint)
			endpoint = nil
		}

		if endpoint == nil {
			endpoint = &Endpoint{
				Method: command.key,
				Route:  command.value.value,
			}
		}

		if ParameterTokens[command.key] {
			switch command.key {
			case TokenName:
				endpoint.Name = command.value.value
			case TokenResponse:
				endpoint.Options.Response = command.value.value
				if command.value.key == TokenJSON && len(endpoint.Options.ContentType) == 0 {
					endpoint.Options.ContentType = "application/json"
				}
			}
		}

		command = parser.ScanCommand()
	}

	if endpoint != nil {
		app.Endpoints = append(app.Endpoints, *endpoint)
	}

	return app
}
