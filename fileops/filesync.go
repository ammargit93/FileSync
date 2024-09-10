package fileops

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	// "cliapp/textutil"
)

func SyncTwoFiles(filepath string, c chan string) {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	content := <-c
	close(c)
	_, err = f.Write([]byte(content))

	if err != nil {
		fmt.Println(err)
	}
}

func UpdateFile(filePath1 string, filePath2 string) error {

	_, err := os.Stat(filePath2)
	if os.IsNotExist(err) {
		_, err := os.Create(filePath2)
		if err != nil {
			fmt.Println(err)
		}
	}

	exepath := "C:\\Windows\\system32\\notepad.exe"
	cmd := exec.Command(exepath, filePath1)
	output, _ := cmd.CombinedOutput()

	s1, _ := os.Open(filePath1)
	warr, _ := io.ReadAll(s1)

	c := make(chan string)
	go SyncTwoFiles(filePath2, c)
	c <- string(warr)

	fmt.Print(string(warr))
	fmt.Print(output)

	return nil
}
