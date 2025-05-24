package utils

import (
	"fmt"
	"strings"
)

func GetIndexFirstChar(str string, char string) int {
	index := strings.Index(str, char)
	return index
}

func IsLower(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func ConvertLowercaseToNumber(str string) (int, error) {
	result := 0

	for i := 0; i < len(str); i++ {
		if !IsLower(str[i]) {
			return -1, fmt.Errorf("Invalid input. Not within a-z")
		}

		amount := int(str[i] - 'a')
		result += amount
	}

	return result, nil
}
