package algo;

type MergeHeap []MergablePartition;

type MergablePartition struct {
  index int64;
  partition []uint32;
}

func (heap MergeHeap) Len() int {
  return len(heap)
}

func (heap MergeHeap) Less(i, j int) bool {
  return heap[i].partition[heap[i].index] < heap[j].partition[heap[j].index];
}

func (heap *MergeHeap) Push(x interface{}) {
  *heap = append(*heap, x.(MergablePartition));
}

func (heap MergeHeap) Swap(i, j int) {
  heap[i], heap[j] = heap[j], heap[i];
}

func MergablePartitionFromSlice(partition []uint32) MergablePartition {
  return MergablePartition{0, partition}
}

func (heap *MergeHeap) Pop() interface{} {
  old := *heap;
  n := len(old);

  x := old[n-1];
  
  *heap = old[:n-1];
  return x;
}

