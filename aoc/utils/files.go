package utils

import (
	"os"
	"strings"
)

func ReadContent(fileName string) string {
	content, _ := os.ReadFile(fileName)
	return string(content)
}

func ReadLines(fileName string) []string {
	content, _ := os.ReadFile(fileName)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func ReadListOfStrings(fileName string) []string {
	content, _ := os.ReadFile(fileName)
	return strings.Split(strings.TrimSpace(string(content)), " ")
}
