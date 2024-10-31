package fileops

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
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

	if err != nil {
		fmt.Print(output)
	}

	s1, _ := os.Open(filePath1)
	warr, _ := io.ReadAll(s1)

	c := make(chan string)
	go SyncTwoFiles(filePath2, c)
	c <- string(warr)

	return nil
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	if _, err := os.Stat(filepath); err != nil {
		f, err := os.Create(filepath)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func SpawnGoroutine(fileUrl []string, path string) error {
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < len(fileUrl); i++ {
		var s []string = strings.Split(fileUrl[i], "/")
		var str string = s[len(s)-1]
		fmt.Println(fileUrl[i])
		wg.Add(1)
		go func() {
			DownloadFile(path+str, fileUrl[i])
			wg.Done()
		}()

	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println("Time to execute: %s", end)
	return nil
}
