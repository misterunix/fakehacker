package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func Roll(count, sides int) int {
	var t int
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		t = t + int(r)
	}
	return t
}

func cryptoRandSecure(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0
	}
	return nBig.Int64()
}

// Chunks : Split the string into a slice where each string has the max length of chunkSize.
// Each string in the slice is left justfied and padded with spaces.
func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}

	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			//	fmt.Fprintf(os.Stderr, "::%d\n", len(s[currentStart:i]))
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])

	for i, tl := range chunks {
		tt := fmt.Sprintf("%%-%dv", chunkSize)
		ttt := fmt.Sprintf(tt, tl)
		chunks[i] = ttt
	}
	return chunks
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
		lines = append(lines, ss)
	}
	return lines, scanner.Err()
}
