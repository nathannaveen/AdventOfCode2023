package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var m = make(map[key]key) // key{J, i diff, j diff} -> key{i diff, j diff}

func main() {
	m[key{"J", 0, 1}] = key{i: -1, j: 0}
	m[key{"J", 1, 0}] = key{i: 0, j: -1}
	m[key{"L", 0, -1}] = key{i: -1, j: 0}
	m[key{"L", 1, 0}] = key{i: 0, j: 1}
	m[key{"7", 0, 1}] = key{i: 1, j: 0}
	m[key{"7", -1, 0}] = key{i: 0, j: -1}
	m[key{"F", 0, -1}] = key{i: 1, j: 0}
	m[key{"F", -1, 0}] = key{i: 0, j: 1}
	m[key{"|", 1, 0}] = key{i: 1, j: 0}
	m[key{"|", -1, 0}] = key{i: -1, j: 0}
	m[key{"-", 0, 1}] = key{i: 0, j: 1}
	m[key{"-", 0, -1}] = key{i: 0, j: -1}

	part1()
	part2()
}

func part1() {
	file, err := os.Open("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		arr = append(arr, strings.Split(s, ""))
	}

	posI, posJ := 0, 0
	cnt := 1

	for i := 0; i < len(arr); i++ {
		shouldBreak := false
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == "S" {
				posI = i
				posJ = j
				shouldBreak = true
				break
			}
		}
		if shouldBreak {
			break
		}
	}

	nextI, nextJ := 0, 1 // hardcoded start direction

	for {
		curI, curJ := posI+nextI, posJ+nextJ
		cnt++
		nextI = m[key{arr[curI][curJ], curI - posI, curJ - posJ}].i
		nextJ = m[key{arr[curI][curJ], curI - posI, curJ - posJ}].j
		posI, posJ = curI, curJ

		if arr[posI][posJ] == "S" {
			break
		}
	}

	fmt.Printf("the farthest pipe is %v pipes from the start\n\n", cnt/2) // the farthest point is the length of the path divided by 2

}

func part2() {
	file, err := os.Open("10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr [][]string
	l := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := "." + scanner.Text() + "." // add padding to avoid out of bounds errors
		l = len(s)
		arr = append(arr, strings.Split(s, ""))
	}

	var s []string
	for i := 0; i < l; i++ {
		s = append(s, ".") // add padding to avoid out of bounds errors
	}

	arr = append([][]string{s}, append(arr, s)...)

	enlarged := make([][]string, len(arr)*2)

	for i := 0; i < len(arr)*2; i++ {
		enlarged[i] = make([]string, len(arr[0])*2)
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			enlarged[i*2][j*2] = arr[i][j]
			switch arr[i][j] {
			case "7":
				enlarged[i*2][j*2-1] = "-"
				enlarged[i*2+1][j*2] = "|"
			case "J":
				enlarged[i*2][j*2-1] = "-"
				enlarged[i*2-1][j*2] = "|"
			case "L":
				enlarged[i*2][j*2+1] = "-"
				enlarged[i*2-1][j*2] = "|"
			case "F":
				enlarged[i*2][j*2+1] = "-"
				enlarged[i*2+1][j*2] = "|"
			case "|":
				enlarged[i*2+1][j*2] = "|"
				enlarged[i*2-1][j*2] = "|"
			case "-":
				enlarged[i*2][j*2+1] = "-"
				enlarged[i*2][j*2-1] = "-"
			}
		}
	}

	posI, posJ := 0, 0
	loop := map[pos]bool{}

	for i := 0; i < len(enlarged); i++ {
		shouldBreak := false
		for j := 0; j < len(enlarged[i]); j++ {
			if enlarged[i][j] == "S" {
				posI = i
				posJ = j
				shouldBreak = true
				loop[pos{i: i, j: j}] = true
				break
			}
		}
		if shouldBreak {
			break
		}
	}

	nextI, nextJ := 0, 1 // hardcoded start direction

	for {
		curI, curJ := posI+nextI, posJ+nextJ
		loop[pos{i: curI, j: curJ}] = true

		nextI = m[key{enlarged[curI][curJ], curI - posI, curJ - posJ}].i
		nextJ = m[key{enlarged[curI][curJ], curI - posI, curJ - posJ}].j
		posI, posJ = curI, curJ

		if enlarged[posI][posJ] == "S" {
			break
		}
	}

	// go around the perimeter of the enlarged board
	for i := 0; i < len(enlarged); i++ {
		for _, j := range []int{0, len(enlarged[0]) - 1} {
			if enlarged[i][j] != "X" {
				noEnclosed(enlarged, i, j, loop)
			}
		}
	}
	for j := 1; j < len(enlarged[0])-1; j++ {
		for _, i := range []int{0, len(enlarged) - 1} {
			if enlarged[i][j] != "X" {
				noEnclosed(enlarged, i, j, loop)
			}
		}
	}

	/*
		if you want to see the enlarged board properly, use a small test case, for example:

		..........
		.S------7.
		.|F----7|.
		.||OOOO||.
		.||OOOO||.
		.|L-7F-J|.
		.|II||II|.
		.L--JL--J.
		..........
	*/

	PrintBoard(enlarged)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			arr[i][j] = enlarged[i*2][j*2]
		}
	}

	res := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			s2 := arr[i][j]
			if loop[pos{i * 2, j * 2}] || s2 == "X" {
				continue
			}
			res++
		}
	}

	PrintBoard(arr)

	fmt.Println(res)

}

func noEnclosed(arr [][]string, i, j int, loop map[pos]bool) {
	if loop[pos{i: i, j: j}] || arr[i][j] == "X" {
		return
	}

	arr[i][j] = "X"

	for _, v := range [][]int{{i, j + 1}, {i, j - 1}, {i + 1, j}, {i - 1, j}} {
		if inBounds(arr, v[0], v[1]) {
			noEnclosed(arr, v[0], v[1], loop)
		}
	}
}

func inBounds(arr [][]string, i, j int) bool {
	return i >= 0 && i < len(arr) && j >= 0 && j < len(arr[0])
}

func PrintBoard(arr [][]string) {
	for _, v := range arr {
		for _, v2 := range v {
			x := v2
			if x == "" {
				x = "."
			}
			fmt.Print("\"" + x + "\" ")

		}
		fmt.Println()
	}
	fmt.Println()
}

type key struct {
	s    string
	i, j int
}

type pos struct {
	i, j int
}
