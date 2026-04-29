
type pair struct {
    val  int
    freq int
}

type PairHeap []pair

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i].freq < h[j].freq } // min-heap by freq
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x any) {
    *h = append(*h, x.(pair))
}

func (h *PairHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}
func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)

	h := &PairHeap {}

	for _, v := range(nums){
        hash[v]++
    }

    for val, f := range hash {
    heap.Push(h, pair{val: val, freq: f})
    if h.Len() > k {
        heap.Pop(h)
        }
    }
    result := make([]int, k)
    for i := 0; i < k; i++ {
        p := heap.Pop(h).(pair)  // cast 'any' back to pair
        result[i] = p.val
    }
    return result
}
