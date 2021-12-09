package dataStruct

import "fmt"

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	//맨뒤에 새로운 것 추가
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		//추가된 노드가 부모노드보다 값이 크다면, 서로 자리 교체
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Count() int {
	return len(h.list)
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	//맨 위의 값을 가져옴
	top := h.list[0]
	//맨 뒤의 값을 가져옴
	last := h.list[len(h.list)-1]
	//리스트에서 맨 뒤의 것을 제외.
	h.list = h.list[:len(h.list)-1]

	//맨 뒤의 것을 맨 위로 바꿈
	h.list[0] = last
	idx := 0

	//현재 인덱스가 리스트의 길이 보다 크지 않는 경우
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2+1
		//왼쪽 자식 노드가 없음.
		if leftIdx >= len(h.list){
			break
		}
		//왼쪽이랑 현재값이랑 비교해서 왼쪽이 크다
		if h.list[leftIdx] >h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2+2
		//오른쪽 자식 노드가 존재
		if rightIdx < len(h.list) {
			//오른쪽 인덱스가 큰경 우
			if h.list[rightIdx] > h.list[idx] {
				//현재 오른쪽 자식노드가 왼쪽 보다 큰 경우

				//스왑인덱스가 0보다 크거나 같은 경우에, 스왑인덱스가 있는 지 확인하고,
				//스왑인덱스 값이 나의 값보다 작은 경우에 바꾼다.
				//스왑인덱스가 0보다 작은 경우에는 왼쪽 값이 현재 값보다 크지 않다는 의미이므로 바로 오른쪽 값과 바꿈.
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
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