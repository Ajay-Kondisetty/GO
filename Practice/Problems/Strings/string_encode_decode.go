package main

import (
	"fmt"
	"strconv"
)

func main() {
	var decodedStr, encodedStr string
	_, _ = fmt.Scanf("%s", &decodedStr)
	_, _ = fmt.Scanf("%s", &encodedStr)

	fmt.Println(fmt.Sprintf("Encoded string of given string(%s) is: %s", decodedStr, StringEncode(decodedStr)))
	fmt.Println(fmt.Sprintf("Decoded string of given string(%s) is: %s", encodedStr, StringDecodeNew(encodedStr)))
}

func StringEncode(inputStr string) string {
	// input ---> WIINNNGGIIFFFFFFFYYYY
	// output ---> XK2Q3I2K2M7C4
	if !isStringValid(inputStr, "") {
		return fmt.Sprintf("Invalid String")
	}

	var outStr string
	var count int
	for i := 0; i < len(inputStr); i++ {
		letter := string(inputStr[i])
		val := inputStr[i]

		if isStringValid(letter, "e") {
			count += 1
			if i == len(inputStr)-1 || letter != string(inputStr[i+1]) {
				temp := int(val) + (count % 26)
				if temp > 90 {
					temp = (65 + (temp - 90)) - 1
				}

				if count == 1 {
					outStr += string(rune(temp))
				} else {
					outStr += fmt.Sprintf("%s%d", string(rune(temp)), count)
				}
				count = 0
			}
		} else {
			return fmt.Sprintf("Invalid String")
		}
	}

	return outStr
}

func StringDecode(inputStr string) string {
	// input ---> XK2Q3I2K2M7C4
	// output ---> WIINNNGGIIFFFFFFFYYYY

	if !isStringValid(inputStr, "") {
		return fmt.Sprintf("Invalid String")
	}

	var outStr string
	i := 0
	for i = 0; i < len(inputStr)-1; i++ {
		item := string(inputStr[i])
		nextItem := string(inputStr[i+1])
		if isStringValid(item, "d") && isStringValid(nextItem, "d") {
			if item >= "A" && item <= "Z" {
				temp := 0
				digit := 0
				var err error
				if digit, err = strconv.Atoi(string(inputStr[i+1])); err != nil {
					temp = int(inputStr[i]) - 1
				} else {
					temp = int(inputStr[i]) - (digit % 26)
				}

				if temp < 65 {
					temp = (90 - (65 - temp)) + 1
				}
				if digit == 0 {
					outStr += fmt.Sprintf("%s", string(rune(temp)))
				} else {
					for j := 0; j < digit; j++ {
						outStr += fmt.Sprintf("%s", string(rune(temp)))
					}
				}
			}
		} else {
			return fmt.Sprintf("Invalid String")
		}
	}

	if _, err := strconv.Atoi(string(inputStr[i])); err != nil {
		temp := int(inputStr[i]) - 1
		if temp < 65 {
			temp = (90 - (65 - temp)) + 1
		}
		outStr += fmt.Sprintf("%s", string(rune(temp)))
	}

	return outStr
}

func StringDecodeNew(inputStr string) string {
	// input ---> XK2Q3I2K2M7C4
	// output ---> WIINNNGGIIFFFFFFFYYYY

	if !isStringValid(inputStr, "") {
		return fmt.Sprintf("Invalid String")
	}

	var outStr string
	currentLetter := -1
	startIndex := -1
	for i := 0; i < len(inputStr); i++ {
		if isStringValid(string(inputStr[i]), "d") {
			if _, err := strconv.Atoi(string(inputStr[i])); err != nil {
				if i != len(inputStr)-1 {
					if digit, err := strconv.Atoi(string(inputStr[i+1])); err != nil {
						freq := 0
						if startIndex != -1 {
							freq, _ = strconv.Atoi(inputStr[startIndex:(i + 1)])
						} else {
							currentLetter = int(inputStr[i])
							freq = 1
						}
						outStr += processDecode(currentLetter, freq)
						currentLetter, startIndex = -1, -1
					} else {
						if digit == 0 || digit == 1 {
							return fmt.Sprintf("Invalid String")
						}
						currentLetter = int(inputStr[i])
						startIndex = i + 1
					}
				}
			} else {
				if i == 0 {
					return fmt.Sprintf("Invalid String")
				}

				if i != len(inputStr)-1 {
					if _, err := strconv.Atoi(string(inputStr[i+1])); err != nil {
						freq, _ := strconv.Atoi(inputStr[startIndex:(i + 1)])
						outStr += processDecode(currentLetter, freq)
						currentLetter, startIndex = -1, -1
					}
				}
			}

			if i == len(inputStr)-1 {
				freq := 0
				if startIndex != -1 {
					freq, _ = strconv.Atoi(inputStr[startIndex:(i + 1)])
				} else {
					currentLetter = int(inputStr[i])
					freq = 1
				}
				outStr += processDecode(currentLetter, freq)
			}
		} else {
			return fmt.Sprintf("Invalid String")
		}
	}

	return outStr
}

func processDecode(s, f int) string {
	var outStr string

	temp := s - (f % 26)

	if temp < 65 {
		temp = (90 - (65 - temp)) + 1
	}

	for j := 0; j < f; j++ {
		outStr += fmt.Sprintf("%s", string(rune(temp)))
	}

	return outStr
}

func isStringValid(s, t string) bool {
	if s == "" {
		return false
	}

	if t != "" {
		if s >= "A" && s <= "Z" {
			return true
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				i = -1
			}
			if t == "d" && i >= 0 {
				return true
			}

			return false
		}
	}

	return true
}
