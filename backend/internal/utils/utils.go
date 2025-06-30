package utils

import (
	"fmt"
	"math"
	"strings"

	"github.com/KainoaGardner/csc/internal/types"
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

	for i := len(str) - 1; i >= 0; i-- {
		if !IsLower(str[i]) {
			return -1, fmt.Errorf("Invalid input. Not within a-z")
		}

		amount := int(str[i]-'a') + 1
		y := len(str) - 1 - i
		result += amount * int(math.Pow(float64(26), float64(y)))
	}

	return result, nil
}

func ConvertNumberToLowercase(x int) (string, error) {
	result := ""
	if !(x >= 0) {
		return "", fmt.Errorf("Invalid input. X lower that 0")
	}
	for x > 0 {
		amount := (x - 1) % 26
		x = (x - 1) / 26
		char := byte(amount) + 'a'
		result = string(char) + result
	}

	return result, nil
}

func AbsoluteValueInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func CheckVec2Equal(x types.Vec2, y types.Vec2) bool {
	return x.X == y.X && x.Y == y.Y
}
