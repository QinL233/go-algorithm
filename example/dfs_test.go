package example

import (
	"fmt"
	"testing"
)

/*
*
全排序：
输入[1,2,3,4]
输出所有不重复的排序可能
*/
func TestDfs(t *testing.T) {
	input := []int{1}
	result := make([][]int, 0)
	uniMap := make(map[string]bool, 0)
	for i := 0; i < len(input); i++ {
		for _, v := range dfs([]int{input[i]}, subSlice(input, i)) {
			s := toString(v)
			if !uniMap[s] {
				result = append(result, v)
				uniMap[s] = true
			}
		}
	}
	fmt.Println(result)
	// table := make(map[string]bool, 0)
}

/*
*
分析
假设[1]的时候，拿到[2、3、4]进行不同组合组合
假设[2]的时候，拿到[1、3、4]进行不同组合组合
假设[3]的时候，拿到[1、2、4]进行不同组合组合

假设[1、2]的时候，拿到[3、4]进行不同组合组合
假设[1、3]的时候，拿到[2、4]进行不同组合组合
假设[1、4]的时候，拿到[2、3]进行不同组合组合

假设[1、2、3]的时候，拿到[4]仅有一个组合，则返回
假设[1、2、4]的时候，拿到[3]仅有一个组合，则返回
结社
*/
func dfs(left []int, right []int) (out [][]int) {
	if len(right) < 1 {
		out = append(out, left)
		return
	}
	if len(right) == 1 {
		out = append(out, append(left, right[0]))
		return
	}
	for i, v := range right {
		r := dfs(append(left, v), subSlice(right, i))
		for _, v := range r {
			out = append(out, v)
		}
	}
	return
}

func subSlice(input []int, i int) []int {
	newArray := make([]int, 0)
	for j, v := range input {
		if i != j {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

/*
去重
*/
func toString(input []int) string {
	s := ""
	for _, v := range input {
		s += fmt.Sprintf("%d", v)
	}
	return s
}
