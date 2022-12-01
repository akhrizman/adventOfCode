package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var elfCalorieMap map[int]int

func main() {
	file, err := os.Open("2022/Day01/input")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Searching for the most Caloric Elf")
		defer file.Close()
		setElfCalorieMap(file)
		PrintTotalOfMostCaloric(1)
		PrintTotalOfMostCaloric(3)
	}
}

// setElfCalorieMap - Read input file line by line and populate global elfCalorieMap
func setElfCalorieMap(file *os.File) {
	scanner := bufio.NewScanner(file)
	elfCalorieMap = map[int]int{}

	total, elf := 0, 1
	for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			elfCalorieMap[elf] = total
			elf++
			total = 0
		} else {
			total += calories
		}
	}
}

// getSortedCalorieTotals - Sort elfCalorieMap values in descending order
func getSortedCalorieTotals() []int {
	calorieTotals := make([]int, 0, len(elfCalorieMap))
	for _, calorieTotal := range elfCalorieMap {
		calorieTotals = append(calorieTotals, calorieTotal)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calorieTotals)))
	return calorieTotals
}

// PrintTotalOfMostCaloric - sum the calorieTotals of the first N elves in the sorted list
func PrintTotalOfMostCaloric(elves int) {
	sortedCalorieTotals := getSortedCalorieTotals()
	result := 0
	for _, calorieTotal := range sortedCalorieTotals[0:elves] {
		result += calorieTotal
	}
	fmt.Printf("Total of Highest %d: %d\n", elves, result)
}
