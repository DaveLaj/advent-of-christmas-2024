package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstList := []int{}
	secondList := []int{}

	// doneScanning := make(chan bool)
	// finalTotalDiff := make(chan int, 2)

	file, err := os.OpenFile("testcase2.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()

		reader := strings.NewReader(text)
		bufReader := bufio.NewReader(reader)

		first, second, err := parseTwoInts(bufReader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		firstList = append(firstList, first)
		secondList = append(secondList, second)

		if err := scanner.Err(); err != nil {
			fmt.Println("Error during scanning:", err)
		}

		if text == "" {
			fmt.Println("Empty line detected")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scanning: %v", err)
	}

	rightMap := make(map[int]int)
	for _, r := range secondList {
		rightMap[r]++
	}

	similarity := 0
	for _, f := range firstList {
		count := rightMap[f]
		similarity += count * f
	}

	fmt.Println("Similarity Score: ", similarity)
}

func parseTwoInts(r *bufio.Reader) (int, int, error) {
	// Read the first integer
	first, err := readNextInt(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read first integer: %w", err)
	}

	// Read the second integer
	second, err := readNextInt(r)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read second integer: %w", err)
	}

	return first, second, nil
}

func readNextInt(r *bufio.Reader) (int, error) {
	var sb strings.Builder

	for {
		// Read one character at a time
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		// Stop at the first space
		if b == ' ' {
			// Skip consecutive spaces
			if sb.Len() == 0 {
				continue
			}
			break
		}

		// Append non-space characters
		sb.WriteByte(b)
	}

	// Convert the accumulated string to an integer
	if sb.Len() == 0 {
		return 0, fmt.Errorf("no integer found")
	}

	return strconv.Atoi(sb.String())
}
