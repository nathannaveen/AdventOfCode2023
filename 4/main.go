package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text()[strings.Index(scanner.Text(), ":")+2:], " ")

		encountered := false // if we have encountered the "|"
		m := map[string]bool{}
		n := 0.5

		for _, v := range s {
			if v == "" {
				continue
			}
			if v == "|" {
				encountered = true
			} else if !encountered {
				m[v] = true
			} else if m[v] {
				res += int(math.Ceil(n))
				n *= 2
			}
		}
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	copies := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := strings.TrimSpace(scanner.Text()[4:])
		id, _ := strconv.Atoi(x[:strings.Index(x, ":")])
		s := strings.Split(x[strings.Index(x, ":")+2:], " ")

		encountered := false // if we have encountered the "|"
		m := map[string]bool{}
		numberOfCardsWon := 0

		for _, v := range s {
			if v == "" {
				continue
			}
			if v == "|" {
				encountered = true
			} else if !encountered {
				m[v] = true
			} else if m[v] {
				numberOfCardsWon++
			}
		}

		copies[id]++

		res += copies[id]

		for i := 0; i < numberOfCardsWon; i++ {
			copies[id+i+1] += copies[id]
		}
	}

	fmt.Println(res)
}
