package fileops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type RequestBody struct {
	Model    string              `json:"model"`
	Messages []map[string]string `json:"messages"`
}

func ModelPrompt(prompt string) (string, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"
	reqBody := RequestBody{
		Model: "llama3-8b-8192",
		Messages: []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status Code:", resp.StatusCode)
	re, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(re) == 0 {
		fmt.Println("Response Body is empty.")
	}

	var in map[string]interface{}
	if err := json.Unmarshal(re, &in); err != nil {
		return "", err
	}

	// Access the "choices" field, which is an array
	choices, ok := in["choices"].([]interface{})
	if !ok {
		return "", err
	}

	// Access the first choice's message
	if len(choices) > 0 {
		choice := choices[0].(map[string]interface{})
		messages, ok := choice["message"].(map[string]interface{})
		if !ok {
			return "", err
		}

		return messages["content"].(string), nil
	}

	return "", nil
}

func ReadCodeFile(file string) string {
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)

	}
	fileContent, err := os.ReadFile(f.Name())
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	absPath, _ := filepath.Abs(f.Name())
	fmt.Println(absPath)
	return "// " + absPath + "\n" + string(fileContent)
}

func IsDirOrNot(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
	if err != nil {
		fmt.Printf("Error accessing %s: %v\n", path, err)
	}

	return info.IsDir()
}
