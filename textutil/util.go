package textutil

import (
	"io"
	"os"
	"strings"
)

func CountFreq(f string) (map[string]int, error) {
	var hashMap = make(map[string]int)
	file, err := os.Open(f)
	if err != nil {
		return hashMap, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	textArr := strings.TrimSpace(string(data))

	word := strings.Fields(textArr)
	for _, p := range word {
		hashMap[p]++
	}
	return hashMap, nil

}

func CountChar(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	file.Close()

	return len(data), nil
}

func CountWords(f string) (int, error) {
	file, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	text := strings.TrimSpace(string(data)) // Trim extra spaces
	if text == "" {                         // If file is empty, return count as 0
		return 0, nil
	}

	words := strings.Fields(text) // Fields splits by any whitespace, handling multiple spaces correctly
	return len(words), nil
}
