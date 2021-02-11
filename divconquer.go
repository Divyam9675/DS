package main

import "fmt"

// Now we need to pass finMaxSubArray function in to maxsubArray function which works as a full array element

func maxsubArray(num []int) int {

	right := len(num) - 1
	return findMaxSubArray(num, 0, right)

}

// this function takes an array and first and last index

func findMaxSubArray(num []int, left int, right int) int {

	// lastly we need to mentioned that we need to divide the array untill we left with last element
	// => So by using condition like left = right

	if left == right {
		return num[left]
	} else {

		// now how to get max sum out of these 3
		//=> By calling each of the function

		mid := (left + right) / 2
		leftMax := findMaxSubArray(num, left, mid)
		rightMax := findMaxSubArray(num, mid+1, right)
		crossMax := maxCrossing(num, left, mid, right)

		// return larger one between these 3

		return max(leftMax, rightMax, crossMax)

	}

}

// Now we call maxcrossing function which works for to findout maximum elements in the mid

func maxCrossing(num []int, left int, right int, mid int) int {

	leftSum := -99999
	rightSum := -99999

	sum := 0

	for i := mid; i >= left; i-- {

		sum += num[i]
		leftSum = max(leftSum, sum)

	}

	sum = 0

	for i := mid + 1; i <= right; i++ {

		sum += num[i]
		rightSum = max(rightSum, sum)

	}

	return leftSum + rightSum

}

// Now we create Max function Here we use variadic function which will creates

func max(values ...int) int {

	// this function gives max value by iterating using for loop

	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func main() {

	problem := []int{-2, -3, 5, 4, -4, -3, 1, 0}
	answer := maxsubArray(problem)

	fmt.Println(answer)

}
