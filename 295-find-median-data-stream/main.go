package findmediandatastream

import "container/heap"

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MaxHeap) Less(i, j int) bool {
	return h[j] < h[i]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	root int // root of the max heap

	empty bool // indicates if the heaps are empty

	left  *MaxHeap // max heap
	right *MinHeap // min heap
}

func Constructor() MedianFinder {

	leftHeap := &MaxHeap{}
	rightHeap := &MinHeap{}

	heap.Init(leftHeap)
	heap.Init(rightHeap)

	return MedianFinder{
		empty: true,
		left:  leftHeap,
		right: rightHeap,
	}
}

func (mf *MedianFinder) AddNum(num int) {

	if mf.empty {
		mf.root = num
		mf.empty = false
		return
	}

	if num <= mf.root {
		heap.Push(mf.left, num)
		if mf.left.Len() > mf.right.Len()+1 {
			// left heap is larger, move the largest element to the right heap
			largest := heap.Pop(mf.left).(int)
			heap.Push(mf.right, mf.root)
			mf.root = largest
		}
	} else {
		heap.Push(mf.right, num)
		if mf.right.Len() > mf.left.Len()+1 {
			// right heap is larger, move the smallest element to the left heap
			smallest := heap.Pop(mf.right).(int)
			heap.Push(mf.left, mf.root)
			mf.root = smallest
		}
	}

}

func (mf *MedianFinder) FindMedian() float64 {

	if mf.empty {
		return 0.0
	}

	if mf.left.Len() == 0 && mf.right.Len() == 0 {
		return float64(mf.root)
	}

	// If both heaps are of equal size, the median is the average of the root of both heaps
	if mf.left.Len() == mf.right.Len() {
		return float64(mf.root)
	}

	// If the left heap is larger, the median is the mean between the root and the left heap
	if mf.left.Len() > mf.right.Len() {
		return float64(mf.root+(*mf.left)[0]) / 2.0
	}
	// If the right heap is larger, the median is the mean between the root and the rightheap
	if mf.right.Len() > mf.left.Len() {
		return float64(mf.root+(*mf.right)[0]) / 2.0
	}

	return 0.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
