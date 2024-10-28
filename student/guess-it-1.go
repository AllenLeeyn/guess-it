package main

import (
	"bufio"
	"guess-it-1/pkg/mathFunc"
	"os"
)

func main() {
	var dataSet mathFunc.Data
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dataSet.AddDPoint(scanner.Text())
		dataSet.PrintRange()
	}
}
