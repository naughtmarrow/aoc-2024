package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func firstStar() int {
	input, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	var nums1 []int
	var nums2 []int
	totalDist := 0

	scan := bufio.NewScanner(input)
	for scan.Scan() {
		strarr := strings.Split(scan.Text(), "   ")

		num1, err := strconv.Atoi(strarr[0])
		if err != nil {
			fmt.Println(err)
		}
		nums1 = append(nums1, num1)

		num2, err := strconv.Atoi(strarr[1])
		if err != nil {
			fmt.Println(err)
		}
		nums2 = append(nums2, num2)
	}

	slices.Sort(nums1)
	slices.Sort(nums2)

	for i, x := range nums1 {
		dist := x - nums2[i]
		if dist < 0 {
			dist = -dist
		}

		totalDist += dist
	}

	return totalDist
}

func secondStar() int {
	score := 0

	input, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	var nums1 []int
	var nums2 []int

	scan := bufio.NewScanner(input)
	for scan.Scan() {
		strarr := strings.Split(scan.Text(), "   ")

		num1, err := strconv.Atoi(strarr[0])
		if err != nil {
			fmt.Println(err)
		}
		nums1 = append(nums1, num1)

		num2, err := strconv.Atoi(strarr[1])
		if err != nil {
			fmt.Println(err)
		}
		nums2 = append(nums2, num2)
	}

	cache := make(map[int]int)
	for _, x := range nums1 {
		sim, ok := cache[x]
		if ok {
			score += x * sim
			cache[x] = sim
			continue
		}

		for _, y := range nums2 {
			if x == y {
				sim++
			}
		}

		score += x * sim
		cache[x] = sim
	}

	return score
}

func main() {
	fmt.Println("The aggregate distance is: ", firstStar())
	fmt.Println("The similarity score is: ", secondStar())
}
