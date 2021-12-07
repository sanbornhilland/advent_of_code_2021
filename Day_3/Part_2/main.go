package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binaryToDecimal(bin []int) int {
	num := 0

	for i, _ := range bin {
		num += bin[i] << (len(bin) - 1 - i)
	}

	return num
}

func split(report [][]int, pos int) ([][]int, [][]int) {
	var ones, zeros [][]int

	for _, line := range report {
		if line[pos] == 1 {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}

	return ones, zeros
}

func findO2Rating(report [][]int, pos int) int {
	ones, zeros := split(report, pos)

	if len(ones) == 1 {
		return binaryToDecimal(ones[0])
	} else if len(ones) >= len(zeros) {
		return findO2Rating(ones, pos+1)
	} else {
		return findO2Rating(zeros, pos+1)
	}
}

func findCO2Scrubber(report [][]int, pos int) int {
	ones, zeros := split(report, pos)

	if len(zeros) == 1 {
		return binaryToDecimal(zeros[0])
	} else if len(zeros) <= len(ones) {
		return findCO2Scrubber(zeros, pos+1)
	} else {
		return findCO2Scrubber(ones, pos+1)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	digits := make([][]int, len(lines))
	for i, line := range lines {
		split := strings.Split(line, "")
		digits[i] = make([]int, len(split))

		for j, char := range split {
			d, _ := strconv.Atoi(char)

			digits[i][j] = d
		}
	}

	fmt.Println(findO2Rating(digits, 0) * findCO2Scrubber(digits, 0))
}
