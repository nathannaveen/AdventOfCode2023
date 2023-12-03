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
	file, err := os.Open("2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.TrimLeft(scanner.Text(), "Game ")

		z := strings.Index(s, ":")
		id, _ := strconv.Atoi(s[:z])
		s = s[z+1:]

		sets := strings.Split(s, ";")

		add := true

		for _, set := range sets {
			dice := strings.Split(set, ",")
			add = true
			for _, d := range dice {
				d = d[1:]
				spaceIndex := strings.Index(d, " ")
				cnt, _ := strconv.Atoi(d[:spaceIndex])
				color := d[spaceIndex+1:]

				if cnt > cubes[color] {
					add = false
					break
				}
			}
			if !add {
				break
			}
		}

		if add {
			res += id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.TrimLeft(scanner.Text(), "Game ")
		s = s[strings.Index(s, ":")+1:]

		sets := strings.Split(s, ";")

		add := true

		cubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			dice := strings.Split(set, ",")
			for _, d := range dice {
				d = d[1:]
				spaceIndex := strings.Index(d, " ")
				cnt, _ := strconv.Atoi(d[:spaceIndex])
				color := d[spaceIndex+1:]

				if cnt > cubes[color] {
					cubes[color] = cnt
				}
			}
		}

		if add {
			res += cubes["red"] * cubes["green"] * cubes["blue"]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
