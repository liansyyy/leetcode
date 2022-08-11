package leetcode_go

import (
	"math"
	"math/rand"
	"sort"
)

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			for j < k && j > i+1 && nums[j] == nums[j-1] {
				j++
			}
			for j < k && k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
			}
			if j >= k {
				break
			}
			if sum := nums[i] + nums[j] + nums[k]; sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			}
			//for j < k && j-1 > i && nums[j] == nums[j-1] {
			//	j++
			//}
			//for j < k && k+1 < len(nums) && nums[k] == nums[k+1] {
			//	k--
			//}
		}
	}
	return
}

func sortArray(nums []int) []int {
	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		randomIndex := rand.Intn(right-left+1) + left
		nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

		pivot, pivotIndex := nums[left], left
		l, r := left, right
		for l < r {
			for l < r && nums[r] >= pivot {
				r--
			}
			for l < r && nums[l] <= pivot {
				l++
			}
			nums[l], nums[r] = nums[r], nums[l]
		}
		nums[l], nums[pivotIndex] = nums[pivotIndex], nums[l]
		quickSort(nums, left, l-1)
		quickSort(nums, l+1, right)
	}
	//quickSort(nums, 0, len(nums)-1)

	var heapSort func(nums []int)
	heapSort = func(nums []int) {
		var buildMaxHeap func(nums []int, end int)
		var maxHeapify func(i, end int)

		buildMaxHeap = func(nums []int, end int) {
			for i := end / 2; i >= 0; i-- {
				maxHeapify(i, end)
			}
		}
		maxHeapify = func(i, end int) {
			if i >= end {
				return
			}
			for j := i; j <= end; {
				leftChild, rightChild := 2*j+1, 2*j+2
				if leftChild <= end && nums[j] < nums[leftChild] {
					j = leftChild
				}
				if rightChild <= end && nums[j] < nums[rightChild] {
					j = rightChild
				}
				if i == j {
					break
				}
				nums[i], nums[j] = nums[j], nums[i]
				i = j
			}
		}

		buildMaxHeap(nums, len(nums)-1)
		for i := len(nums) - 1; i >= 0; i-- {
			nums[0], nums[i] = nums[i], nums[0]
			maxHeapify(0, i-1)
		}
	}
	heapSort(nums)

	var mergeSort func(nums []int, left, right int)
	var merge func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := (left + right) >> 1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, right)
	}
	merge = func(nums []int, left, right int) {
		temp := make([]int, right-left+1)
		i, j, k := left, (left+right)>>1+1, 0
		for i <= (left+right)>>1 && j <= right {
			if nums[i] < nums[j] {
				temp[k] = nums[i]
				i++
			} else {
				temp[k] = nums[j]
				j++
			}
			k++
		}
		for i <= (left+right)>>1 {
			temp[k] = nums[i]
			k++
			i++
		}
		for j <= right {
			temp[k] = nums[j]
			k++
			j++
		}
		//for i := 0; i < len(temp); i++ {
		//	nums[left+i] = temp[i]
		//}
		copy(nums[left:right+1], temp)
	}

	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func maxSubArray(nums []int) (res int) {
	res = nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		res = max(res, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return
}

func maxSubArray(nums []int) (res int) {
	var foo func(nums []int, left, right int) int
	foo = func(nums []int, left, right int) int {
		if left > right {
			return math.MinInt32
		}
		mid := (left + right) >> 1
		leftSum, rightSum := foo(nums, left, mid-1), foo(nums, mid+1, right)
		leftSumCrossMid, rightSumCrossMid := 0, 0

		for i, temp := mid-1, 0; i >= left; i-- {
			temp += nums[i]
			leftSumCrossMid = max(leftSumCrossMid, temp)
		}
		for i, temp := mid+1, 0; i <= right; i++ {
			temp += nums[i]
			rightSumCrossMid = max(rightSumCrossMid, temp)
		}
		return max(leftSum, rightSum, leftSumCrossMid+rightSumCrossMid+nums[mid])
	}

	return foo(nums, 0, len(nums)-1)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, num := range nums {
		if v, ok := hash[target-num]; ok {
			return []int{v, i}
		} else {
			hash[num] = i
		}
	}
	return nil
}
