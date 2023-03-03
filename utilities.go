package main

import (
	"bufio"
	"compress/gzip"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

func rnd(min, max int) int {
	r := cryptoRandSecure(int64(max-min)) + int64(min)
	return int(r)
}

// Roll : generates a random like an old dice
func Roll(count, sides int) int {
	var t int
	for i := 0; i < count; i++ {
		r := cryptoRandSecure(int64(sides)) + 1
		t = t + int(r)
	}
	return t
}

// cryptoRandSecure : generate random number from 0 to max
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

// Read in compressed file and return the contents split by <cr>
func readGzipLines(path string) ([]string, error) {

	// referenced file is embeded in the global space
	file, err := efile.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// need a reader to read the gzip file
	gz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer gz.Close()

	// read it line be line into a slice
	scanner := bufio.NewScanner(gz)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
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

// readLines : Read file from "path" spliting into lines on the newline.
func readLinesRaw(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readFileToString(path string) (string, error) {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return "", err
	}
	r := string(b)
	return r, nil
}

func collision(x0, y0, x1, y1, x2, y2, x3, y3 int) bool {

	width1 := x1 - x0
	height1 := y1 - y0
	width2 := x3 - x2
	height2 := y3 - y2

	if x0 < x2+width2 &&
		x0+width1 > x2 &&
		y0 < y2+height2 &&
		y0+height1 > y2 {
		return true
	}
	return false
}
