package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const input_file = "../input"

func main() {
	lines := read_file(input_file)
	answer := 0
	
	var new_lines []string
	line_jump_fix := "do"
	for _, line := range lines {
		do_dont_list := strings.Split(line, "do")
		var do_list string
		for _, element := range do_dont_list {
			if element[:4] == "n't(" {
				line_jump_fix = "dont"
				continue
			} else if element[:2] == "()" {
				line_jump_fix = "do"
				do_list += element
			} else if line_jump_fix == "dont" {
				continue
			} else {
				do_list += element
			}

		}
		new_lines = append(new_lines, do_list)
	}

	for _, line := range new_lines {
		fmt.Println(line)
		split_mul := strings.Split(line, "mul(")
		for _, mul := range split_mul[1:] {
			split_numbers := strings.Split(mul, ")")
			separated_numbers := strings.Split(split_numbers[0], ",")
			if len(separated_numbers) != 2 {
				continue
			}
			number_one, err_one := strconv.Atoi(separated_numbers[0])
			number_two, err_two := strconv.Atoi(separated_numbers[1])
			if err_one != nil {
				log.Println(err_one)
				continue
			}
			if err_two != nil {
				log.Println(err_two)
				continue
			}
			answer += number_one * number_two
		}
	}
	fmt.Println(answer)
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