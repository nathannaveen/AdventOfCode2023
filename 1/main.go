package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		start, end := -1, 0
		for _, l := range scanner.Text() {
			if unicode.IsDigit(l) {
				if start == -1 {
					start = int(l - '0')
				}

				end = int(l - '0')
			}
		}

		n := start*10 + end

		sum += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	digits := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
		"1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
		"6": 6, "7": 7, "8": 8, "9": 9, "10": 10,
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		start, end := -1, 0

		for i := 0; i < len(scanner.Text()); i++ {
			for k, v := range digits {
				if i+len(k) <= len(scanner.Text()) && scanner.Text()[i:i+len(k)] == k {
					if start == -1 {
						start = v
					}

					end = v
				}
			}
		}

		n := start*10 + end

		sum += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
