package textutil

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func FindMatchingWords(f []string) ([]string, error) {
	file1, _ := os.Open(f[0])
	file2, _ := os.Open(f[1])

	// if err != nil {
	// 	return err
	// }
	defer file1.Close()
	defer file2.Close()

	data1, _ := io.ReadAll(file1)
	textArr1 := strings.TrimSpace(string(data1))
	wordArr1 := strings.Fields(textArr1)

	data2, _ := io.ReadAll(file2)
	textArr2 := strings.TrimSpace(string(data2))
	wordArr2 := strings.Fields(textArr2)

	var l []string
	if len(wordArr1) < len(wordArr2) {
		l = wordArr1
	} else {
		l = wordArr2
	}

	var res []string
	var j int
	for i := range l {
		if wordArr1[i] == wordArr2[j] {
			res = append(res, wordArr1[i])
			j++
		}
	}
	return res, nil
}

func CountFreq(f []string) (map[string]int, error) {
	var hashMap = make(map[string]int)
	file, err := os.Open(f[0])
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

func CountChar(f []string) (int, error) {
	file, err := os.Open(f[0])
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	file.Close()

	return len(data), nil
}

func CountWords(f []string) (int, error) {
	file, err := os.Open(f[0])
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
