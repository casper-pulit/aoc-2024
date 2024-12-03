package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var input_length int
var doido []bool
var muls []string
var muls_i [][]int
var dos []int
var do_i [][]int
var dont_i [][]int
var sum int = 0
var sum_cond int = 0

// the name is a trap
func FlattenIndexes(indexes [][]int) []int {
	var flat []int

	for i := range indexes {
		flat = append(flat, indexes[i][0])
	}

	return flat
}

func SliceContains(slice []int, find int) bool {
	for _, v := range slice {
		if find == v {
			return true
		}
	}
	return false
}

func main() {
	// open file
	file, err := os.Open("day3/input")

	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	// for not really required, formatted input to one line
	for sc.Scan() {
		// get length of input for checking dos and donts of mul later
		input_length = len(sc.Text())
		r := regexp.MustCompile("mul\\(\\d+,\\d+\\)")

		// get the actuals muls and indexes of the muls
		muls = r.FindAllString(sc.Text(), 100000)
		muls_i = r.FindAllStringIndex(sc.Text(), 100000)

		r_do := regexp.MustCompile("do\\(\\)")
		r_dont := regexp.MustCompile("don't\\(\\)")

		// get indexes of the dos and donts
		do_i = r_do.FindAllStringIndex(sc.Text(), 100000)
		dont_i = r_dont.FindAllStringIndex(sc.Text(), 100000)
	}
	// start dos off at zero (default condition)
	dos = append(dos, 0)
	// append slice of the start of do only
	dos = append(dos, FlattenIndexes(do_i)...)
	// same for dont's but dont need to default at 0
	donts := FlattenIndexes(dont_i)

	// create slice of bools: doido
	// true true true true true false false false true
	state := true

	// create bool map slice thing to determine whether to do or not to do at each index
	for i := 0; i < input_length; i++ {

		if state == true && SliceContains(donts, i) {
			state = false
		} else if state == false && SliceContains(dos, i) {
			state = true
		}

		doido = append(doido, state)
	}

	// fyi only
	// for i, v := range doido {
	// 	fmt.Println(v, " at: ", i)
	// }

	for i := 0; i < len(muls); i++ {

		r := regexp.MustCompile("[^\\d|^,]")
		clean := r.ReplaceAllString(muls[i], "")
		f_mul, err := strconv.Atoi(strings.Split(clean, ",")[0])

		if err != nil {
			log.Fatal(err)
		}

		l_mul, err := strconv.Atoi(strings.Split(clean, ",")[1])

		if err != nil {
			log.Fatal(err)
		}

		sum += f_mul * l_mul
	}
	fmt.Println("The sum of muls is: ", sum)

	for i := 0; i < len(muls); i++ {

		if doido[muls_i[i][0]] {
			r := regexp.MustCompile("[^\\d|^,]")
			clean := r.ReplaceAllString(muls[i], "")
			f_mul, err := strconv.Atoi(strings.Split(clean, ",")[0])

			if err != nil {
				log.Fatal(err)
			}

			l_mul, err := strconv.Atoi(strings.Split(clean, ",")[1])

			if err != nil {
				log.Fatal(err)
			}

			sum_cond += f_mul * l_mul
		}

	}
	fmt.Println("The sum of muls (conditional) is: ", sum_cond)
}
