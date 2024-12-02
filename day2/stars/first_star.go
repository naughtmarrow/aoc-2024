package stars

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func First() int {
	safe := 0

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var nprev *int
		dir := 0

		line := strings.Split(scanner.Text(), " ")
		length := len(line)
		for i, x := range line {
			n, err := strconv.Atoi(x)
			if err != nil {
				log.Fatal(err)
			}

			if nprev != nil {
				diff := n - *nprev

				if diff == 0 || diff < -3 || diff > 3 {
					break // it's unsafe so ignore line
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
						break
					}
				}

				if dir == -1 {
					if diff > 0 {
						break
					}
				}
			}

			nprev = &n
			if i == length-1 {
				safe++
			}
		}
	}

	return safe
}
