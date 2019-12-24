package main

import (
	"fmt"
	"sort"
  "container/heap"
)

type kv struct {
  Key   string
  Value int
}
type kvHeap []kv

func (h kvHeap) Len() int           { return len(h) }
func (h kvHeap) Less(i, j int) bool {
  if h[i].Value == h[j].Value {
    return h[i].Key > h[j].Key
  }
  return h[i].Value < h[j].Value
}
func (h kvHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *kvHeap) Push(x interface{}) {
  *h = append(*h, x.(kv))
}

func (h *kvHeap) Pop() interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[0 : n-1]
  return x
}

func main() {
  words := []string{"i", "love", "leetcode", "i", "love", "coding"}
  fmt.Println(topKFrequent(words, 2))
}

func topKFrequent(words []string, k int) []string {
    sort.Strings(words)
    collect := map[string]int{}
    for _, word := range words {
      collect[word]++
    }

    h := &kvHeap{}
    heap.Init(h)
    for key, val := range collect {
      heap.Push(h, kv{key, val})
      if h.Len() > k {
        heap.Pop(h)
      }
    }

    ans := make([]string, k)
    for i := k-1; i >= 0; i-- {
        wc := heap.Pop(h).(kv)
        ans[i] = wc.Key
    }
    return ans
}
