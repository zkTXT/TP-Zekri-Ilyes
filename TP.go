package main

import (
	"fmt"
)

func Ft_coin(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Ft_missing(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j][1] > intervals[j+1][1] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}

	end := intervals[0][1]
	count := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

func Ft_profit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func Ft_max_substring(s string) int {
	charIndex := make(map[byte]int)
	maxLength, start := 0, 0

	for i := 0; i < len(s); i++ {
		if index, ok := charIndex[s[i]]; ok && index >= start {
			start = index + 1
		}
		charIndex[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	countT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	left, right := 0, 0
	required := len(countT)
	formed := 0
	windowCounts := make(map[byte]int)
	ans := []int{-1, 0, 0}

	for right < len(s) {
		c := s[right]
		windowCounts[c]++
		if countT[c] > 0 && windowCounts[c] == countT[c] {
			formed++
		}

		for left <= right && formed == required {
			c = s[left]
			if ans[0] == -1 || right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}
			windowCounts[c]--
			if countT[c] > 0 && windowCounts[c] < countT[c] {
				formed--
			}
			left++
		}
		right++
	}

	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

func main() {
	fmt.Println("")
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0

	fmt.Println("")
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // resultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // resultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8

	fmt.Println("")
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // resultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // resultat : 2

	fmt.Println("")
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))    // resultat : 0

	fmt.Println("")
	fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
	fmt.Println(Ft_max_substring("bbbbb"))    // resultat : 1

	fmt.Println("")
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // resultat : ""
}
