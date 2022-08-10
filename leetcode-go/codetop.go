package leetcode_go

import (
	"container/heap"
	"math/rand"
)

func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode = nil, head, nil
	for cur != nil {
		next = cur.Next
		prev, cur, cur.Next = cur, next, prev
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func lengthOfLongestSubstring(s string) (res int) {
	hash := map[rune]int{}
	left := -1
	for i, ch := range s {
		if j, ok := hash[ch]; ok {
			left = max(left, j)
		}
		hash[ch] = i
		res = max(res, i-left)
	}
	return
}

type LRUCache struct {
	head, tail *LRUCacheNode
	hash       map[int]*LRUCacheNode
	length     int
	capacity   int
}

type LRUCacheNode struct {
	prev, next *LRUCacheNode
	key, value int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		head:     &LRUCacheNode{},
		tail:     &LRUCacheNode{},
		hash:     map[int]*LRUCacheNode{},
		length:   0,
		capacity: capacity,
	}
	lruCache.head.next, lruCache.tail.prev = lruCache.tail, lruCache.head
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hash[key]; ok {
		this.moveToFront(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.hash[key]; ok {
		node.value = value
		this.moveToFront(node)
	} else {
		if this.isFull() {
			this.deleteNode(this.tail.prev)
		}
		node := new(LRUCacheNode)
		node.key, node.value = key, value
		this.insertNode(node)
	}
}

func (this *LRUCache) moveToFront(node *LRUCacheNode) {
	this.deleteNode(node)
	this.insertNode(node)
}

func (this *LRUCache) isFull() bool {
	return this.length >= this.capacity
}

func (this *LRUCache) deleteNode(node *LRUCacheNode) {
	delete(this.hash, node.key)
	node.prev.next, node.next.prev = node.next, node.prev
	this.length--
}

func (this *LRUCache) insertNode(node *LRUCacheNode) {
	if this.isFull() {
		this.deleteNode(this.tail.prev)
	}
	this.hash[node.key] = node
	node.prev, node.next = this.head, this.head.next
	this.head.next, this.head.next.prev = node, node
	this.length++
}

func findKthLargest(nums []int, k int) int {
	//return quickSearch(nums, 0, len(nums)-1, k)
	//return heapSortSearch(nums, k)
	nums = heapSort(nums)
	return nums[len(nums)-k]
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

// 决定是大跟堆还是小根堆
// 小根堆
func (m maxHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	pop := (*m)[len(*m)-1]
	*m = (*m)[:m.Len()-1]
	return pop
}

func heapSort(nums []int) []int {
	temp := make([]int, 0)
	hp := maxHeap{}
	for _, num := range nums {
		heap.Push(&hp, num)
	}
	for hp.Len() != 0 {
		pop := heap.Pop(&hp)
		temp = append(temp, pop.(int))
	}
	return temp
}

func heapSortSearch(nums []int, k int) int {
	end := len(nums) - 1
	buildMaxHeap(nums, end)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		maxHeapify(nums, 0, end)
	}
	return nums[0]
}

func maxHeapify(nums []int, i int, end int) {
	if i >= end {
		return
	}
	for j := i; j < end; {
		leftChild, rightChild := 2*j+1, 2*j+2
		if leftChild <= end && nums[j] < nums[leftChild] {
			leftChild, j = j, leftChild
		}
		if rightChild <= end && nums[j] < nums[rightChild] {
			rightChild, j = j, rightChild
		}
		if i == j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func buildMaxHeap(nums []int, end int) {
	for i := end / 2; i >= 0; i-- {
		maxHeapify(nums, i, end)
	}
}

//随机化快排
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	randomIndex := rand.Intn(right-left+1) + left
	nums[left], nums[randomIndex] = nums[randomIndex], nums[left]

	pivot := nums[left]
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
	nums[left], nums[l] = nums[l], nums[left]
	quickSort(nums, left, l-1)
	quickSort(nums, l+1, right)
}

//The three-way radix quicksort algorithm
func quickSort(nums []int, lessThan, greatThan int) {
	if lessThan >= greatThan {
		return
	}
	pivot := nums[lessThan]
	l, g := lessThan, greatThan
	for i := lessThan; i <= greatThan; {
		if nums[i] < pivot {
			nums[lessThan], nums[i] = nums[i], nums[lessThan]
			lessThan++
			i++
		} else if nums[i] == pivot {
			i++
		} else {
			nums[greatThan], nums[i] = nums[i], nums[greatThan]
			greatThan--
		}
	}
	quickSort(nums, l, lessThan-1)
	quickSort(nums, greatThan+1, g)
}

func quickSearch(nums []int, left, right int, k int) int {
	if left > right {
		return -1
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
	nums[pivotIndex], nums[l] = nums[l], nums[pivotIndex]

	if l == len(nums)-k {
		return pivot
	} else if l > len(nums)-k {
		return quickSearch(nums, left, l-1, k)
	} else {
		return quickSearch(nums, l+1, right, k)
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var reverseList func(left, right *ListNode) (*ListNode, *ListNode)
	reverseList = func(left, right *ListNode) (*ListNode, *ListNode) {
		var prev, curr *ListNode = nil, left
		for curr != nil {
			next := curr.Next
			prev, curr.Next = curr, prev
			curr = next
		}
		return right, left
	}

	dummyHead := &ListNode{Next: head}
	curr := dummyHead.Next
	prevTail := dummyHead
	for curr != nil {
		left := curr
		for i := 1; i < k && curr != nil; i++ {
			curr = curr.Next
		}
		if curr == nil {
			break
		}
		right, nextHead := curr, curr.Next
		curr.Next = nil
		left, right = reverseList(left, right)

		prevTail.Next = left
		prevTail = right
		right.Next = nextHead
		curr = nextHead
	}
	return dummyHead.Next
}
