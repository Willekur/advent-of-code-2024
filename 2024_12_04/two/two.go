package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const input_file = "../input"
var height int
var width int

func main() {
	lines := read_file(input_file)
	height = len(lines)
	width = len(lines[0])

	counter := 0
	for y_pos, line := range lines {
		for x_pos, char := range line {
			if char == 'A' {
				counter += search_for_cross(lines, y_pos, x_pos)
			}
		}
	}
	fmt.Println(counter)
}

func search_for_cross(lines []string, y_pos int, x_pos int) int {
	counter := 0
	if y_pos-1 >= 0 && x_pos-1 >= 0 && (lines[y_pos-1][x_pos-1] == 'M' || lines[y_pos-1][x_pos-1] == 'S') {
		up_left := lines[y_pos-1][x_pos-1]
		if y_pos+1 < height && x_pos+1 < width && (lines[y_pos+1][x_pos+1] == 'M' || lines[y_pos+1][x_pos+1] == 'S') {
			down_right := lines[y_pos+1][x_pos+1]
			if y_pos+1 < height && x_pos-1 >= 0 && (lines[y_pos+1][x_pos-1] == 'M' || lines[y_pos+1][x_pos-1] == 'S') {
				down_left := lines[y_pos+1][x_pos-1]
				if y_pos-1 >= 0 && x_pos+1 < width && (lines[y_pos-1][x_pos+1] == 'M' || lines[y_pos-1][x_pos+1] == 'S') {
					up_right := lines[y_pos-1][x_pos+1]
					if up_left != down_right && up_right != down_left {
						counter += 1
					}
				}
			}
		}
	}
	return counter
}

func read_file(file_name string) []string {
	file, _ := os.Open(file_name)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}