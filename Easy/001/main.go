// https://leetcode.com/problems/two-sum/description/
// 複数渡されるnumsから2つの数を足して、targetを満たす組み合わせのインデックスを返す
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	numsLine := scanner.Text()
	numsStr := strings.Fields(numsLine)
	nums := make([]int, len(numsStr))
	for i, s := range numsStr {
		nums[i], _ = strconv.Atoi(s)
	}

	scanner.Scan()
	targetLine := scanner.Text()
	target, _ := strconv.Atoi(targetLine)

	// twoSum関数の呼び出しと結果の表示
	result := twoSum(nums, target)
	fmt.Println(result)
}
