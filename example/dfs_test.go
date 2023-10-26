package example

import (
	"fmt"
	"testing"
)

/**
全排序：
输入[1,2,3,4]
输出所有不重复的排序可能
*/
func TestDfs(t *testing.T) {
	input := []int{1, 2, 3, 4}
	out := struct {
		result [][]int
		table  map[string]bool
	}{make([][]int, 0), make(map[string]bool, 0)}
	for i := 0; i < len(input); i++ {
		dfs([]int{input[i]}, subSlice(input, i))
	}
	fmt.Println(out.result)
}

/**
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
func dfs(left []int, right []int) []int {
	if len(right) == 1 {
		return append(left, right[0])
	}
	for i, v := range right {
		dfs(append(left, v), subSlice(right, i))
	}
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
去重复
*/
func record(input []int) {
	s := ""
	for _, v := range input {
		s += fmt.Sprintf("%d", v)
	}
	if !table[s] {
		result = append(result, input)
		table[s] = true
	}
}
