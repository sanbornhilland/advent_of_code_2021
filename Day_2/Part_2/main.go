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
	aim := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		command := line[0]
		moveCount, err := strconv.Atoi(line[1])

		if err != nil {
			fmt.Println("Error reading movement count")
		}

		if command == "down" {
			aim += moveCount
		} else if command == "up" {
			aim -= moveCount
		} else if command == "forward" {
			horizontalPos += moveCount
			depth += aim * moveCount
		} else {
			fmt.Println("Unknown command", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(depth * horizontalPos)
}
