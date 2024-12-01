package adventofcode2024

import (
	"bufio"
	"embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input/*
var input embed.FS

func Day1p1(file string) int {
	input := getInput(file)
	list := loadList(input)
	return sumDifferences(list)
}

func Day1p2(file string) int {
	input := getInput(file)
	list := loadList(input)
	return calculateScore(list)
}

func getInput(filename string) string {
	data, err := input.ReadFile(filename)
	if err != nil {
		panic("no input")
	}
	return string(data)
}

type list struct {
	Length int
	Left   []int
	Right  []int
}

func (list *list) NumberOfTimesInRight(number int) int {
	count := 0
	for i := 0; i < list.Length; i += 1 {
		if list.Right[i] == number {
			count += 1
		}
	}
	return count
}

func loadList(input string) list {
	result := list{
		Length: 0,
		Left:   []int{},
		Right:  []int{},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		left, err := strconv.Atoi(split[0])
		if err != nil {
			panic("failed")
		}
		right, err := strconv.Atoi(split[1])
		if err != nil {
			panic("failed 2")
		}

		result.Left = append(result.Left, left)
		result.Right = append(result.Right, right)
		result.Length += 1
	}

	sort.Slice(result.Left, func(i, j int) bool {
		return result.Left[i] < result.Left[j]
	})

	sort.Slice(result.Right, func(i, j int) bool {
		return result.Right[i] < result.Right[j]
	})

	return result
}

func sumDifferences(list list) int {
	result := 0
	for i := 0; i < list.Length; i += 1 {
		difference := list.Left[i] - list.Right[i]
		if difference <= 0 {
			difference *= -1
		}
		result += difference
	}
	return result
}

func calculateScore(list list) int {
	result := 0
	for i := 0; i < list.Length; i += 1 {
		left := list.Left[i]
		result += left * list.NumberOfTimesInRight(left)
	}
	return result
}
