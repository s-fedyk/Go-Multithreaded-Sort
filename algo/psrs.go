package algo

import (
	"container/heap"
	"sort"
	"sync"
)

var wg = sync.WaitGroup{}

func psrs(list []uint32, p int) {
  /*
  leftOver := len(list) % p;
  offset := 0

  for i := range(p) {
    sampleAdjust := 0;
    if i > leftOver {
      sampleAdjust = 1;
    }

    sampleAdjust = i > leftOver ? 1 : 0;

  }
  */
}

func phase1(localSample []uint32, sampleDest []uint32, w uint32, p uint32) {
  
  sort.Slice(localSample, func(i, j int) bool {
    return localSample[i] < localSample[j]
  });

  for i:= uint32(0) ; i < p ; i++ {
    sampleDest[int(i)] = localSample[i * w];
  }
}

func phase2(gatheredSample *[]uint32, p uint32) *[]uint32 {
  pivots := make([]uint32, p-1);

  sort.Slice(*gatheredSample, func(i, j int) bool {
    return (*gatheredSample)[i] < (*gatheredSample)[j]
  });

  for i := range(pivots) {
    pivots[i] = (*gatheredSample)[int(p) + i * int(p)]
  }

  return &pivots;
}

func phase3(localSample *[]uint32, pivots []uint32) *[][]uint32 {
  partitions := make([][]uint32, len(pivots) + 1);
  sampleStart := localSample;

  // partition[i] contains all elements in the local sample less than pivots[i]
  for partitionIndex := range(pivots) {
    index := BinarySearch(sampleStart, pivots[partitionIndex])

    partitions[partitionIndex] = (*sampleStart)[:index]
 
    // cut down search space, next pivot is bigger.
    (*sampleStart) = (*sampleStart)[index:]
  }

  // the last partition contains everything else
  partitions[len(pivots)] = (*sampleStart)

  return &partitions
}

func phase4(destination []uint32, partitions ...[]uint32) {
  mheap := &MergeHeap{};
  array := make([]MergablePartition, len(partitions));

  // heapify
  for i := range(partitions) {
    array[i] = MergablePartitionFromSlice(partitions[i]);
  }

  *mheap = array;
  heap.Init(mheap);
  outputIndex := 0

  //k-way merge
  for len(*mheap) > 0 {
    item := heap.Pop(mheap).(MergablePartition);
    destination[outputIndex] = item.partition[item.index];
    outputIndex += 1;

    // aren't at the end of partition. increment, re-add
    if int(item.index + 1) < len(item.partition) {
      item.index += 1;
      heap.Push(mheap, item);
    }
  }
}
