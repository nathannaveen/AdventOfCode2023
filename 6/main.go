package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var times []int
	var distances []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if s[:4] == "Time" {
			s = s[5:]
			s = strings.Trim(s, " ")
			arr := strings.Split(s, " ")

			for _, v := range arr {
				if v == "" {
					continue
				}
				n, _ := strconv.Atoi(v)
				times = append(times, n)
			}
		} else {
			s = s[9:]
			s = strings.Trim(s, " ")
			arr := strings.Split(s, " ")

			for _, v := range arr {
				if v == "" {
					continue
				}
				n, _ := strconv.Atoi(v)
				distances = append(distances, n)
			}
		}
	}

	res := 1

	for i := 0; i < len(times); i++ {
		cnt := 0

		for j := 0; j < times[i]; j++ {
			x := j * (times[i] - j)
			if x > distances[i] {
				cnt++
			}
		}

		if cnt != 0 {
			res *= cnt
		}
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var time int
	var distance int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if s[:4] == "Time" {
			s = s[5:]
			s = strings.Trim(s, " ")
			arr := strings.Split(s, " ")

			time, _ = strconv.Atoi(strings.Join(arr, ""))
		} else {
			s = s[9:]
			s = strings.Trim(s, " ")
			arr := strings.Split(s, " ")

			distance, _ = strconv.Atoi(strings.Join(arr, ""))
		}
	}

	res := 0

	for j := 0; j < time; j++ {
		x := j * (time - j)
		if x > distance {
			res++
		}
	}

	fmt.Println(res)
}
