package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)


func main() {
	var formatted []string
	headers := loadHeaders()
	if len(headers) == 0 {
		fmt.Print("Headers Length: 0 | Format Invalid or Missing")
		os.Exit(1)
	}

	for _, v := range headers {
		s := strings.Split(v, ":")
		sK := strings.TrimSpace(s[0])
		sV := strings.TrimSpace(s[1])
		sFormat := string("\"")
		if strings.Index(s[1], `"`) != -1 {
			sFormat = string("`")
		}

		formatted = append(formatted, fmt.Sprintf("req.Header.Set(\"%s\", %s%s%s)", sK, sFormat, sV, sFormat))
	}

	fmt.Println(strings.Join(formatted, "\n"))
}

func loadHeaders() []string {
	if exists := fileExists("headers.txt"); exists {
		file, err := os.Open("./headers.txt")
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer file.Close()

		var headers []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			headers = append(headers, scanner.Text())
		}

		return headers
	} else {
		fmt.Print("Could Not Open File Headers.txt")
	}

	return []string{}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}