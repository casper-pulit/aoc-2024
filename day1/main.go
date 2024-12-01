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

func UniqueVals(slice []int) []int {
	uniqueMap := make(map[int]bool)
	var uniqueSlice []int

	for _, val := range slice {
		if !uniqueMap[val] { // Check if the value is not already in the map
			uniqueMap[val] = true
			uniqueSlice = append(uniqueSlice, val) // Add to the result slice
		}
	}

	return uniqueSlice
}

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

	dist := float64(0)
	for i := 0; i < len(left); i++ {

		diff := right[i] - left[i]
		dist += math.Abs(float64(diff))
	}
	fmt.Println("Distance between the two lists is:")
	fmt.Println(int(dist))
	// doesn't actually do anything because all values in left list are already unique
	u_left := UniqueVals(left)
	sim_score := 0
	for i := range u_left {
		count := 0
		for ii := range right {
			if u_left[i] == right[ii] {
				count += 1
			}
		}
		sim_score += u_left[i] * count

	}
	fmt.Println("Similarity score between the two lists is:")
	fmt.Println(sim_score)

}
