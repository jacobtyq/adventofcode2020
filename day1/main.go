package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	content, err := os.Open("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers := []int{}
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	content.Close()

	//Convert String to Int and append to array
	for _, i := range txtlines {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}

	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		remainder := 2020 - numbers[i]

		// Part 1
		find(numbers, remainder)

		// Part 2
		start := i + 1
		end := len(numbers) - 1
		for start < end {
			sum := numbers[start] + numbers[end]
			if sum == remainder {
				product := numbers[i] * numbers[start] * numbers[end]
				fmt.Println("Found numbers that sum up to 2020")
				fmt.Println(numbers[i], numbers[start], numbers[end])
				fmt.Println("Product:", product)
				break
			} else if sum < remainder {
				start++
			} else {
				end--
			}
		}
	}

}

//Find two numbers
func find(a []int, x int) int {
	for i, n := range a {
		if x == n {
			fmt.Println(n)
			return i
		}
	}
	return len(a)
}
