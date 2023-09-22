package main

type activity struct {
	cnt        int
	account_id string
}

type ActivityHeap []activity

type NewBank struct {
	AccountIds      map[string]float64
	ActivityCounter ActivityHeap
}

// Heap interface implementation 5 method

func (h ActivityHeap) Len() int {
	return len(h)
}

func (h ActivityHeap) Less(i_index int, j_index int) bool {
	activity_i := h[i_index]
	activity_j := h[j_index]

	if activity_i.cnt == activity_j.cnt {
		return activity_i.account_id < activity_j.account_id
	} else {
		return activity_i.cnt < activity_j.cnt
	}
}

func (h ActivityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ActivityHeap) Push(n activity) {
	*h = append(*h, n)
}
func (h *ActivityHeap) Pop() activity {
	old := *h
	l := len(old)
	x := old[l-1]
	*h = old[:l-1]
	return x
}


