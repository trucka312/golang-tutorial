package dynamic

func LongestIncreasingSubsequenceGreedy(nums []int) int {
	longestIncreasingSubsequence := make([]int, 0)

	for _, num := range nums {
		// find the leftmost index in longestIncreasingSubsequence with value >= num
		leftmostIndex := lowerBound(longestIncreasingSubsequence, num)

		if leftmostIndex == len(longestIncreasingSubsequence) {
			longestIncreasingSubsequence = append(longestIncreasingSubsequence, num)
		} else {
			longestIncreasingSubsequence[leftmostIndex] = num
		}
	}

	return len(longestIncreasingSubsequence)
}

// Function to find the leftmost index in arr with value >= val, mimicking the inbuild lower_bound function in C++
// Time Complexity: O(logn)
// Auxiliary Space: O(1)
func lowerBound(arr []int, val int) int {
	searchWindowLeft, searchWindowRight := 0, len(arr)-1

	for searchWindowLeft <= searchWindowRight {
		middle := (searchWindowLeft + searchWindowRight) / 2

		if arr[middle] < val {
			searchWindowLeft = middle + 1
		} else {
			searchWindowRight = middle - 1
		}
	}

	return searchWindowRight + 1
}
