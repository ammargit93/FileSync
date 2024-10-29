package textutil

import (
	"io"
	"os"
	"strings"
)

func ProcessText(f *os.File) []string {
	data, _ := io.ReadAll(f)
	textArr := strings.TrimSpace(string(data))
	wordArr := strings.Fields(textArr)
	return wordArr
}

func ConcurrentProcess2(wordArr1 []string, wordArr2 []string, l []string, c chan []string) {
	var res []string

	for i := len(l) / 2; i < len(l); i++ {
		if wordArr1[i] == wordArr2[i] {
			res = append(res, wordArr1[i])
		}
	}
	c <- res
	close(c)
}

func ConcurrentProcess1(wordArr1 []string, wordArr2 []string, l []string, c chan []string) {
	var res []string

	for i := 0; i < len(l)/2; i++ {
		if wordArr1[i] == wordArr2[i] {
			res = append(res, wordArr1[i])
		}
	}
	c <- res
	close(c)
}

func FindMatchingWords(f []string) ([]string, error) {
	file1, _ := os.Open(f[0])
	file2, _ := os.Open(f[1])

	defer file1.Close()
	defer file2.Close()

	wordArr1 := ProcessText(file1)
	wordArr2 := ProcessText(file2)

	var l []string
	if len(wordArr1) < len(wordArr2) {
		l = wordArr1
	} else {
		l = wordArr2
	}

	var res []string

	c1 := make(chan []string)
	c2 := make(chan []string)

	go ConcurrentProcess1(wordArr1, wordArr2, l, c1)
	go ConcurrentProcess2(wordArr1, wordArr2, l, c2)

	msg1 := <-c1
	msg2 := <-c2
	res = append(msg1, msg2...)
	// res = append(res, SynchronousProcess(wordArr1, wordArr2, l)...)
	return res, nil
}

func CountFreq(f []string) (map[string]int, error) {
	var hashMap = make(map[string]int)
	file, _ := os.Open(f[0])

	defer file.Close()

	word := ProcessText(file)
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
