package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type inputStruct struct {
	Arr    []int
	Ans    int
	AnsArr []int
	Type   string
}

func main() {
	var inputFile string
	var inputData inputStruct

	flag.StringVar(&inputFile, "json_file_name", "", "Input JSON data file name.")
	flag.Parse()

	if inputFile == "" {
		log.Fatal("Input file is empty.")
	}

	inputJSON, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error reading JSON data from file.")
	}

	err = json.Unmarshal(inputJSON, &inputData)
	if err != nil {
		log.Fatal("Error unmarshalling JSON data.")
	}

	if inputData.Type == "" {
		log.Fatal("Type is not provided in the input data.")
	} else if inputData.Ans == 0 {
		log.Fatal("Answer is not provided in the input data.")
	} else if inputData.Type == "max_array" && inputData.AnsArr == nil {
		log.Fatal("Answer array not provided for type max_array in the input data.")
	}

	message := "Expected Answer"
	output := ""
	res := -1
	resArr := make([]int, 0)
	if inputData.Type == "max" || inputData.Type == "dp" {
		if inputData.Type == "max" {
			res = KadaneAlgo(inputData.Arr)
		} else {
			res = DynamicProgramingAlgo(inputData.Arr)
		}
		if inputData.Ans == res {
			output = "Correct Answer!"
		} else {
			output = "Wrong Answer!"
		}
		output = fmt.Sprintf("%s - %d\n%s - %d", output, res, message, inputData.Ans)
	} else if inputData.Type == "max_array" {
		res, resArr = KadaneAlgoWithSubarry(inputData.Arr)
		fmt.Println(resArr)
		if inputData.Ans == res && slices.Compare(inputData.AnsArr, resArr) == 0 {
			output = "Correct Answer!"
		} else {
			output = "Wrong Answer!"
		}
		output = fmt.Sprintf("%s - %d %v\n%s - %d %v", output, res, resArr, message, inputData.Ans, inputData.AnsArr)
	}

	fmt.Println(output)
}

func KadaneAlgo(arr []int) int {
	maxSum := math.MinInt64
	currSum := 0

	for _, val := range arr {
		currSum += val

		if maxSum < currSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}

	return maxSum
}

func KadaneAlgoWithSubarry(arr []int) (int, []int) {
	maxSum := math.MinInt64
	currSum, start, end, reset := 0, 0, 0, 0

	for i, val := range arr {
		currSum += val

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

	return maxSum, arr[start : end+1]
}

func DynamicProgramingAlgo(arr []int) int {
	dp := make([]int, len(arr))

	dp[0] = arr[0]

	ans := dp[0]

	for i := 1; i < len(arr); i++ {
		dp[i] = int(math.Max(float64(arr[i]), float64(arr[i])+float64(dp[i-1])))

		ans = int(math.Max(float64(ans), float64(dp[i])))
	}

	return ans
}
