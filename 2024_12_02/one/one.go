package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

const input_file = "../input"

func main() {
	reports := read_file(input_file)

	count_safe := 0
	for i := range reports {
		str_report := strings.Split(reports[i], " ")
		
		int_report := report_to_ints(str_report)

		if check_report(int_report) {
			count_safe += 1
		}

	}
	fmt.Println(count_safe)
}

func report_to_ints(report []string) ([]int) {
	var int_report []int
	for _, level := range report {
		int_level, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}
		int_report = append(int_report, int_level)

	}
	return int_report
}

func check_report(report []int) (bool) {
	direction := ""
	first_level := report[0]
	second_level := report[1]
	if first_level > second_level {
		direction = "down"
	} else {
		direction = "up"
	}

	reading_before := report[0]
	for _, reading := range report[1:] {
		if absDiff(reading_before, reading) > 3 {
			return false
		}
		if direction == "up" && reading_before >= reading {
			return false
		} 
		if direction == "down" && reading_before <= reading {
			return false
		}
		reading_before = reading
	}
	return true
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

func absDiff(x, y int) (int) {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}