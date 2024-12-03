package stars

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func validate(nums []int, length int, deleted bool) int {
	dir := 0
	newNums := make([]int, 0)

	for i := 0; i < length-1; i++ {
		diff := nums[i+1] - nums[i]
		if diff == 0 || diff < -3 || diff > 3 {
			// it's unsafe
			if deleted {
				return 0
			}

			copy(newNums, nums[:i+1])
			newNums = append(newNums, nums[i+2:]...)

			return validate(newNums, len(newNums), true)
		}

		if dir == 0 {
			if diff > 0 {
				dir = 1
			}

			if diff < 0 {
				dir = -1
			}
		}

		if dir == 1 {
			if diff < 0 {
				// it's unsafe
				if deleted {
					return 0
				}

				copy(newNums, nums[:i+1])
				newNums = append(newNums, nums[i+2:]...)

				return validate(newNums, len(newNums), true)
			}
		}

		if dir == -1 {
			if diff > 0 {
				// it's unsafe
				if deleted {
					return 0
				}

				copy(newNums, nums[:i+1])
				newNums = append(newNums, nums[i+2:]...)

				return validate(newNums, len(newNums), true)
			}
		}
	}

	return 1
}

func intConv(line []string) []int {
	var is []int
	for _, x := range line {
		n, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		is = append(is, n)
	}

	return is
}

func Second() int {
	safe := 0

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		nums := intConv(line)

		safe += validate(nums, len(nums), false)
	}

	return safe
}
