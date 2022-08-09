package leetcode_go

type ListNode struct {
	Val  int
	Next *ListNode
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(args ...int) (res int) {
	res = args[0]
	for _, arg := range args {
		res = min(res, arg)
	}
	return
}

func max(args ...int) (res int) {
	res = args[0]
	for _, arg := range args {
		res = max(res, arg)
	}
	return
}
