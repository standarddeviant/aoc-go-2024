package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func parse_input() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := [][]int{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	r := regexp.MustCompile(`(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		m := r.FindAllString(line, -1)
		row := []int{}

		for ix := 0; ix < len(m); ix++ {
			tmpX, errX := strconv.Atoi(m[ix])
			if errX != nil {
				fmt.Println("oh no")
			}
			row = append(row, tmpX)
		}
		nums = append(nums, row)
	}

	return nums
}

func check_part1(a []int) bool {
	p_count := 0
	n_count := 0
	for ix := 0; ix < len(a)-1; ix++ {
		d := a[ix+1] - a[ix]
		a := math.Round(math.Abs(float64(d)))
		if d > 0 {
			p_count += 1
		} else {
			n_count += 1
		}
		if a < 1 || 3 < a {
			return false
		}
	}
	// check all same direction
	if p_count > 0 && n_count > 0 {
		return false
	}
	return true
}

func remove(slice []int, s int) []int {
	out := []int{}
	for ix := 0; ix < len(slice); ix++ {
		if ix != s {
			out = append(out, slice[ix])
		}
	}
	return out
}

func check_part2(a []int) bool {
	if check_part1(a) {
		return true
	}
	// fmt.Println("\n*** a = ", a)
	for skip_ix := 0; skip_ix < len(a); skip_ix++ {
		tmpv := remove(a, skip_ix)
		// fmt.Println(skip_ix, " : ", tmpv)
		if check_part1(tmpv) {
			return true
		}
	}
	return false
}

func main() {
	nums := parse_input()

	sum1 := 0
	for ix := 0; ix < len(nums); ix++ {
		if check_part1(nums[ix]) {
			sum1 += 1
		}
	}
	fmt.Println("day02, part1 = ", sum1)

	sum2 := 0
	for ix := 0; ix < len(nums); ix++ {
		if check_part2(nums[ix]) {
			sum2 += 1
		}
	}
	fmt.Println("day02, part2 = ", sum2)
}
