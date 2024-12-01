package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"sort"
)

const input_file = "../input"

func main() {
	lines := read_file(input_file)
	left_values, right_values := data_to_workable_format(lines)

	final_value := count_repetition(left_values, right_values)
	fmt.Println(final_value)
}

func count_repetition(left_values, right_values []int) (int) {
	final_value := 0
	for i := range left_values {
		frequency := 0
		for y := range right_values {
			if left_values[i] == right_values[y] {
				frequency += 1
			}
		}
		final_value += left_values[i] * frequency
	}
	return final_value
}

func data_to_workable_format(lines []string) ([]int, []int) {
	var left_values []int
	var right_values []int
	for index := range lines {
		entry := strings.Split(lines[index], "   ")
		left_entry, left_err := strconv.Atoi(entry[0])
		right_entry, right_err := strconv.Atoi(entry[1])
		if left_err != nil || right_err != nil {
			log.Fatal("Conversion to integer")
		}
		left_values = append(left_values, left_entry)
		right_values = append(right_values, right_entry)
	}
	sort.Ints(left_values[:])
	sort.Ints(right_values[:])
	return left_values, right_values
}

func read_file(file_name string) ([]string) {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
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