package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {

	var left []int
	var right []int
	file, err := os.Open("day1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		left_int, err := strconv.Atoi(scanner.Text()[0:5])

		if err != nil {
			log.Fatal(err)
		}

		left = append(left, left_int)

		right_int, err := strconv.Atoi(scanner.Text()[8:])

		right = append(right, right_int)
	}

	sort.Ints(left)
	sort.Ints(right)

	fmt.Println(left[0:20])
	fmt.Println(right[0:20])

	dist := float64(0)
	for i := 0; i < len(left); i++ {

		diff := right[i] - left[i]
		dist += math.Abs(float64(diff))
	}

	fmt.Println(int(dist))
}
