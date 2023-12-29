package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func main() {

	url := "assets/aoc02input.txt"

	file, _ := os.Open(url)
	scanner := bufio.NewScanner(file)
    
    sum:= 0
	for scanner.Scan() {
		line := scanner.Text()
        lineresult := processLine(line, 12, 14, 13)
        if(lineresult != -1) {
            sum += lineresult
        }
	}
    fmt.Println("result" , sum)

}

func processLine(s string, red int, blue int, green int) int {
    
    substring := s[5:]
    substrings :=  strings.Split(substring, ":")
    // gameid, err2:= strconv.Atoi(substrings[0])
    // if err2 != nil {
    //     fmt.Println("error", err2)
    // }
    minRed, minBlue, minGreen := 0, 0, 0
    differentgames := strings.Split(substrings[1], ";")
    for _, game := range differentgames {
        colors := strings.Split(game, ",")
        for _, color := range colors {
            colorandnum := strings.Split(color, " ")
            num, err := strconv.Atoi(colorandnum[1])
            if err != nil {
                fmt.Println("error", err)
            }
            if(colorandnum[2] == "red" && num > minRed){
                minRed = num
            }

            if(colorandnum[2] == "blue" && num > minBlue){
                minBlue = num
            }

            if(colorandnum[2] == "green" && num > minGreen){
                minGreen = num
            }
        }
    }
    return minRed*minGreen*minBlue
}
