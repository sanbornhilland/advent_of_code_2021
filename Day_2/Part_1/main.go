package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	horizontalPos := 0
	depth := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		command := line[0]
		moveCount, err := strconv.Atoi(line[1])

		if err != nil {
			fmt.Println("Error reading movement count")
		}

		if command == "forward" {
			horizontalPos += moveCount
		} else if command == "down" {
			depth += moveCount
		} else if command == "up" {
			depth -= moveCount
		} else {
			fmt.Println("Unknown command", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(depth * horizontalPos)
}
