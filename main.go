package main

import (
	"log/slog"

	day1 "github.com/fish1/advent-of-code-2024/day1"
)

func main() {
	result := day1.Day1p1("input/production.txt")
	slog.Info("day1p1", "result", result)

	result = day1.Day1p2("input/production.txt")
	slog.Info("day1p2", "result", result)
}
