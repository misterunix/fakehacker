package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, _ := readLines("rawpass.txt")
	file, err := os.OpenFile("../password.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range lines {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

// Calculate the sha256 hash of the specified string s
// Using md5 instead of sha256 since security is not an issue and sha256 produced a 64
// byte checksum where md5 returns a 32 byte checksum.
func CalcFileNameHash(s string) string {
	sum := md5.Sum([]byte(s))
	r := fmt.Sprintf("%x", sum)
	return r
}

// readLines : Read file from "path" spliting into lines on the newline.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ss := strings.TrimSpace(scanner.Text())
		ct := CalcFileNameHash(ss)
		c := fmt.Sprintf("%s $1$%s", ss, ct)
		lines = append(lines, c)
	}
	return lines, scanner.Err()
}
