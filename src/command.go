package main

import "fmt"

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