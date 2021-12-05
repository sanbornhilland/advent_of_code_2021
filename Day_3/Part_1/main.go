package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func binaryToDecimal(bin []int) int {
	num := 0

	for i, _ := range bin {
		num += int(math.Pow(2, float64(i))) * bin[len(bin)-1-i]
	}

	return num
}

func getBits(bits []int) (int, int) {
	s := 0

	for _, bit := range bits {
		s += bit
	}

	if s > (len(bits) / 2) {
		return 1, 0
	} else {
		return 0, 1
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

	res := make([][]int, 2)
	res[0] = make([]int, len(digits[0]))
	res[1] = make([]int, len(digits[0]))
	for i, _ := range digits[0] {
		bits := make([]int, len(lines))

		for j, dx := range digits {
			bits[j] = dx[i]
		}

		g, e := getBits(bits)
		res[0][i] = g
		res[1][i] = e
	}

	fmt.Println(binaryToDecimal(res[0]) * binaryToDecimal(res[1]))
}
