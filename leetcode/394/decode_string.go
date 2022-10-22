package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] == ']' {
			stack = append(stack, string(s[i]))
		}

		if s[i] == '[' {
			var chars string
			for char := stack[len(stack)-1]; len(stack) > 1 && char != "]"; char = stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				chars += char
			}
			stack = stack[:len(stack)-1]

			var integer string
			for j := i - 1; j >= 0 && s[j] >= '0' && s[j] <= '9'; j-- {
				integer = fmt.Sprintf("%s%s", string(s[j]), integer)
			}
			integerInt, _ := strconv.Atoi(integer)
			stack = append(stack, strings.Repeat(chars, integerInt))
		}
	}

	var buffer bytes.Buffer
	for i := len(stack) - 1; i >= 0; i-- {
		buffer.WriteString(stack[i])
	}
	return buffer.String()
}
