package utils

import (
	"os"
	"strings"
)

func ReadLines(fileName string) []string {
	content, _ := os.ReadFile(fileName)
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}
