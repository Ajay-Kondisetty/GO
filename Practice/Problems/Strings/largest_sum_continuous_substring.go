package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

type inputStruct struct {
	W string
	N int
	X []string
	B []int
}

func main() {
	var inputFile string
	var inputData inputStruct
	flag.StringVar(&inputFile, "json_file_name", "", "Input JSON file name")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("Input filename should not be empty.")
	}

	inputJSON, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error reading JSON data from file.")
	}

	if err = json.Unmarshal(inputJSON, &inputData); err != nil {
		log.Fatal("Error unmarshalling JSON data.")
	}

	fmt.Println(FindLargestSumContinuousSubstring(inputData.W, inputData.N, inputData.X, inputData.B))
}

func FindLargestSumContinuousSubstring(w string, n int, x []string, b []int) string {
	maxSum := math.MinInt64
	currSum := 0
	start, end, reset := 0, 0, 0

	for i, val := range w {

		replace := slices.Index(x, string(val))
		if replace == -1 {
			replace = int(val)
		} else {
			replace = b[replace]
		}

		currSum += replace

		if maxSum < currSum {
			maxSum = currSum
			start = reset
			end = i
		}

		if currSum < 0 {
			currSum = 0
			reset = i + 1
		}
	}

	return strings.Join(strings.Split(w, "")[start:end+1], "")
}
