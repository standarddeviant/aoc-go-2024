package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func parse_input() ([]int, []int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	left := []int{}
	right := []int{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	r := regexp.MustCompile(`(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindAllString(line, -1)

		if 2 == len(m) {
			tmpL, errL := strconv.Atoi(m[0])
			if errL != nil {
				fmt.Println("oh no")
			}
			left = append(left, tmpL)

			tmpR, errR := strconv.Atoi(m[1])
			if errR != nil {
				fmt.Println("oh no")
			}
			right = append(right, tmpR)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func main() {
	left, right := parse_input()
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	sum1 := 0.0
	for ix := 0; ix < len(left); ix++ {
		diff := float64(left[ix] - right[ix])
		sum1 += math.Abs(diff)
	}
	fmt.Println("day01, part1 = ", int(sum1))

	n_times := make(map[int]int)
	for ix := 0; ix < len(left); ix++ {
		// val := right[ix]
		key := right[ix]
		n, contains_key := n_times[key]
		if contains_key {
			n_times[key] = n + 1
			continue
		}
		n_times[key] = 1
	}

	sum2 := 0.0
	for ix := 0; ix < len(left); ix++ {
		val := left[ix]
		n, contains_key := n_times[val]
		if contains_key {
			sum2 += float64(val * n)
		}
	}
	fmt.Println("day01, part2 = ", int(sum2))
}
