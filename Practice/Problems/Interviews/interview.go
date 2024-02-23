// Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.
//
// Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,4,7,5,3,6,8,9]
// Example 2:
// Input: mat = [[1,2],[3,4]]
// Output: [1,2,3,4]

package main

import "fmt"

func main() {
    input := [][]int{
        []int{1, 2},
        []int{3, 4},
        []int{5, 6},
        []int{0, -1},
    }

    output := make([]int, 0)

    i, j := 0, 0

    iLen, jLen := len(input), len(input[0])
    isUp := true

    for i < iLen && j < jLen {
        if isUp == true {
            for i >= 0 && j < jLen {
                output = append(output, input[i][j])
                i--
                j++
            }

            if j < jLen {
                i += 1
            } else {
                i += 2
                j = jLen - 1
            }

            isUp = false
        } else if isUp == false {
            for j >= 0 && i < iLen {
                output = append(output, input[i][j])
                i++
                j--
            }

            if i < iLen {
                j += 1
            } else {
                j += 2
                i = iLen - 1
            }

            isUp = true
        }
    }

    fmt.Println(output)
//         1, 2, 3, 4
//         5, 6, 7, 8
//         9, 10, 11, 12
//         13, 14, 15, 16

// 1, 2, 5, 9, 6, 3, 4, 7, 10, 13, 14, 11, 8, 12, 15, 16

}

