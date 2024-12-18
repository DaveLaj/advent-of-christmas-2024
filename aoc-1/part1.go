package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// ch := make(chan int, 5)
	firstList := []int{}
	secondList := []int{}

	// doneScanning := make(chan bool)
	// finalTotalDiff := make(chan int, 2)

	file, err := os.OpenFile("testcase1.txt", os.O_RDONLY, 0644)
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
	} else {
		fmt.Println("Reached EOF")
	}

	sort.Slice(firstList, func(i, j int) bool {
		return firstList[i] < firstList[j]
	})

	sort.Slice(secondList, func(i, j int) bool {
		return secondList[i] < secondList[j]
	})

	var totalDiff int
	for i := 0; i < len(firstList); i++ {
		diff := firstList[i] - secondList[i]

		if diff < 0 {
			diff = -diff
		}

		totalDiff += diff
	}

	fmt.Println("Total diff:", totalDiff)
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
