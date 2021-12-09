package dataStruct

import "fmt"

type MinHeap struct {
	list []int
}

func (h *MinHeap) Push(v int) {
	//맨뒤에 새로운 것 추가
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] < h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.list)
}

func (h *MinHeap) Count() int {
	return len(h.list)
}

func (h *MinHeap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	//맨 위의 값을 가져옴
	top := h.list[0]
	//맨 뒤의 값을 가져옴
	last := h.list[len(h.list)-1]
	//리스트에서 맨 뒤의 것을 제외.
	h.list = h.list[:len(h.list)-1]

	if len(h.list) == 0 {
		return top
	}

	//맨 뒤의 것을 맨 위로 바꿈
	h.list[0] = last
	idx := 0

	//현재 인덱스가 리스트의 길이 보다 크지 않는 경우
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2+1

		if leftIdx >= len(h.list){
			break
		}

		if h.list[leftIdx] < h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2+2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] < h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] > h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
 	}
	return top
}