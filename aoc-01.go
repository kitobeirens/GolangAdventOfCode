package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    url := "assets/aoc01input.txt"

	file, _ := os.Open(url)
	//w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		text := []rune(scanner.Text())
		var first int
		var last int

		for i:= 0 ; i < len(text) ; i++ {
			if v, err := strconv.Atoi(string(text[i])); err == nil {
				first = v
				break
			}

		}
		for i:= 0 ; i < len(text) ; i++ {
			if v, err := strconv.Atoi(string(text[len(text) - 1 -i])); err == nil {
				last = v
				break
			}

		}
		subTotal:= (first * 10) + last
		total = total + subTotal
	}
	fmt.Println(total)
	}
