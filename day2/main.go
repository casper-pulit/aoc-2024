package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var start int
var prev int
var desc bool
var safe []bool
var safe_count int
var diff int
var i int

const input_path string = "day2/input.txt"

func countTrue(slice []bool) int {
	count := 0
	for _, value := range slice {
		if value {
			count++
		}
	}
	return count
}

func main() {

	file, err := os.Open(input_path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		// split the strings
		s := strings.Split(sc.Text(), " ")
		fmt.Println(s)
		// time.Sleep(2 * time.Second)
		// iterate over slice
		for i := 0; i < len(s); i++ {
			// fmt.Println(s[i])

			val, err := strconv.Atoi(s[i])

			if err != nil {
				log.Fatal(err)
			}
			// first iter
			if i == 0 {
				// store value as previous
				prev = val
			} else {

				// checks whether ascending descending or neither
				if i == 1 {
					if val < prev {
						desc = true
					} else if val > prev {
						desc = false
					} else {
						fmt.Println("not desc or asc, unsafe, breaking")
						safe = append(safe, false)
						break
					}
				}

				if desc {
					// fmt.Println("descending")
					diff = prev - val
				}

				if !desc {
					// fmt.Println("ascending")
					diff = val - prev
				}

				prev = val

				if diff < 1 || diff > 3 {
					fmt.Println("not ok")
					safe = append(safe, false)
					break
				} else {
					// fmt.Println("looks ok")
				}

				if i == len(s)-1 {
					safe = append(safe, true)
					fmt.Println("SUCCESS")

				}

			}
		}
	}

	safe_count = countTrue(safe)

	fmt.Println("Answer: ", safe_count)
}

// 	The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.
