package parser

import (
	"fmt"
)

// Function to check if the character at index i is CRLF (\r\n)
func isCRLF(str string, index int) bool {
	// Check if the current and next characters together form CRLF
	if index < len(str)-1 && str[index] == '\r' && str[index+1] == '\n' {
		return true
	}
	return false
}

func DecodeSubCommand(reqStr string) map[string]string {
	spaceCount := 1
	word := ""
	crlfCount := 0
	commandMap := make(map[string]string)
	// PUB CodingChallenge 11\r\nHello John!\r\n
	for index, char := range reqStr {
		ch := fmt.Sprintf("%c", char)
		if ch == " " && crlfCount == 0 {
			if spaceCount == 1 {
				commandMap["command"] = word
			} else if spaceCount == 2 {
				commandMap["topic"] = word
			}
			spaceCount += 1
			word = ""
		} else if isCRLF(reqStr, index) {
			commandMap["subjectId"] = word
			word = ""
			crlfCount += 1
		} else if ch != "\n" && ch != "\r" {
			word += ch
		}
	}
	if len(word) > 0 {
		commandMap["payload"] = word
	}
	return commandMap
}
