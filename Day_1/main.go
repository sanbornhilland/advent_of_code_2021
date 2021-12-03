package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := -1
	prev := -1

	for scanner.Scan() {
		curr, _ := strconv.Atoi(scanner.Text())

		if curr > prev {
			count++
		}

		prev = curr
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(count)
}
