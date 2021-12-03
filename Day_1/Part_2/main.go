package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(arr []int) int {
	result := 0

	for _, v := range arr {
		result += v
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buff := []int{-1, -1, -1}
	count := -1
	prevSum := -1

	for scanner.Scan() {
		curr, _ := strconv.Atoi(scanner.Text())

		buff = append(buff[1:], curr)
		currSum := sum(buff)

		// lol
		if buff[0] == -1 {
			continue
		}

		if buff[1] == -1 {
			continue
		}

		if buff[2] == -1 {
			continue
		}

		if currSum > prevSum {
			count++
		}

		prevSum = currSum
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(count)
}
