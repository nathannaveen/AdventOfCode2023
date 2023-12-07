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
}

func part1() {
	file, err := os.Open("5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cnt := 0

	var seeds []int
	var newSeeds []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		if cnt == 0 {
			s = strings.Split(scanner.Text()[7:], " ")
		}

		if len(s) == 2 {
			copy(seeds, newSeeds)
		} else if len(s) == 3 {
			arr := []int{}
			for _, v := range s {
				n, _ := strconv.Atoi(v)
				arr = append(arr, n)
			}

			for i := 0; i < len(seeds); i++ {
				seed := seeds[i]
				if seed >= arr[1] && seed < arr[1]+arr[2] {
					newSeeds[i] = seed - arr[1] + arr[0]
				}
			}

		} else if scanner.Text() != "" { // these are the start seeds
			for i := 0; i < len(s); i += 2 {
				n, _ := strconv.Atoi(s[i])
				n2, _ := strconv.Atoi(s[i+1])

				for j := n; j < n+n2; j++ {
					seeds = append(seeds, j)
				}
			}
			newSeeds = make([]int, len(seeds))
			copy(newSeeds, seeds)
		}

		cnt++
	}

	copy(seeds, newSeeds)

	fmt.Println(seeds)

	smallest := seeds[0]

	for _, v := range seeds {
		if v < smallest {
			smallest = v
		}
	}

	fmt.Println(smallest)
}
