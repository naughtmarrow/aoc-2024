package stars

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func First() int {
	res := 0

	input, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	expr := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	strs := expr.FindAllString(string(input), -1)

	numexp := regexp.MustCompile(`[0-9]{1,3}`)

	for _, s := range strs {
		nums := numexp.FindAllString(s, 2)
		x, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatal(err)
		}

		res += x * y
	}

	return res
}
