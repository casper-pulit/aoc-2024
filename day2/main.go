package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func readFile(path string) (*bufio.Scanner, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, err
}

func isSafe(s []string) (bool, error) {
	var prev int
	var desc bool

	for i := 0; i < len(s); i++ {
		val, err := strconv.Atoi(s[i])

		if err != nil {
			return false, err
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
					return false, nil
				}
			}

			if desc {
				diff = prev - val
			}

			if !desc {
				diff = val - prev
			}

			prev = val

			if diff < 1 || diff > 3 {
				return false, nil
			}

			if i == len(s)-1 {
				return true, nil
			}

		}
	}

	return true, nil
}

func RemoveIndex(s []string, index int) []string {
	// this must be cursed
	ns := make([]string, 0, len(s)-1)
	if index > 0 {
		ns = append(ns, s[:index]...)
	}
	ns = append(ns, s[index+1:]...)

	return ns
}

func problemDampner(s []string) bool {

	for i := 0; i < len(s); i++ {
		ns := RemoveIndex(s, i)
		is_safe, err := isSafe(ns)

		if err != nil {
			return false
		}

		if is_safe {
			return true
		}
	}
	return false
}

func main() {
	sc, err := readFile(input_path)

	if err != nil {
		log.Fatal(err)
	}

	for sc.Scan() {

		s := strings.Split(sc.Text(), " ")

		sc_safe, err := isSafe(s)

		if err != nil {
			log.Fatal(err)
		}

		if sc_safe {
			safe = append(safe, sc_safe)
		} else {
			safe = append(safe, problemDampner(s))
		}
	}

	fmt.Println("Answer: ", countTrue(safe))

}
