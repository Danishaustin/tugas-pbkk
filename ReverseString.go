package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reverseText string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan String: ")
	text, _ := reader.ReadString('\n')

	split := strings.Fields(text)

	for _, v := range split {
		reverseText += reverseString(v) + " "
	}

	fmt.Println("Reverse String:", reverseText)
}

func reverseString(s string) string {
	text := []rune(s)
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		text[i], text[j] = text[j], text[i]
	}

	return string(text)
}
