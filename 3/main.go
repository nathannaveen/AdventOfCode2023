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

type key struct {
	num int
	id  int
}

type pos struct {
	i int
	j int
}

func part1() {
	file, err := os.Open("3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	added := map[int]bool{}
	posToNum := map[pos]key{}
	var row int
	var symbols []pos
	var id int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := "." + scanner.Text() + "."

		num := 0

		for i := 0; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
			} else {
				num2 := num
				var cnt int
				for num > 0 {
					cnt++
					posToNum[pos{row, i - cnt - 1}] = key{num2, id}
					num /= 10
				}
				id++

				if s[i] != '.' {
					symbols = append(symbols, pos{row, i - 1})
				}
			}
		}
		row++
	}

	for _, symbol := range symbols {
		for _, d := range []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, -1}, {1, 1}, {-1, 1}} {
			cell := pos{symbol.i + d.i, symbol.j + d.j}
			if posToNum[cell].num != 0 && !added[posToNum[cell].id] {
				added[posToNum[cell].id] = true
				res += posToNum[cell].num
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}

func part2() {
	file, err := os.Open("3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := 0

	posToNum := map[pos]key{}
	var row int
	var symbols []pos
	var id int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := "." + scanner.Text() + "."

		num := 0

		for i := 0; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				num = num*10 + int(s[i]-'0')
			} else {
				num2 := num
				var cnt int
				for num > 0 {
					cnt++
					posToNum[pos{row, i - cnt - 1}] = key{num2, id}
					num /= 10
				}
				id++

				if s[i] == '*' {
					symbols = append(symbols, pos{row, i - 1})
				}
			}
		}
		row++
	}

	for _, symbol := range symbols {
		partNumbers := map[int]bool{}
		product := 1
		for _, d := range []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, -1}, {1, 1}, {-1, 1}} {
			cell := pos{symbol.i + d.i, symbol.j + d.j}
			if posToNum[cell].num != 0 && partNumbers[posToNum[cell].id] == false {
				partNumbers[posToNum[cell].id] = true
				product *= posToNum[cell].num
			}
		}
		if len(partNumbers) == 2 {
			res += product
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
