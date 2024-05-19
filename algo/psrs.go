package algo

import (
	"container/heap"
	"sort"
	"sync"
)

type Job struct {
  sample []uint32;
  w uint32;
  p uint32;
  pivots []uint32;
  gatheredSample []uint32;
  partitions [][]uint32;
  wg1 *sync.WaitGroup;
  wg2 *sync.WaitGroup;
  wg3 *sync.WaitGroup;
  wg4 *sync.WaitGroup;
}

func psrs(list []uint32, p uint32, destination []uint32) {
  // this is stupid but GO doesn't seem to have a barrier?
  var wg1 = sync.WaitGroup{};
  var wg2 = sync.WaitGroup{};
  var wg3 = sync.WaitGroup{};
  var wg4 = sync.WaitGroup{};

  n := uint32(len(list));
  w := uint32(n / (p * p));

  leftOver := n % p;
  sampleSize := n / p;

  gatheredSample := make([]uint32, p*p);
  pivots := make([]uint32, p-1);
  partitions := make([][]uint32, p*p);

  wg1.Add(int(p));
  wg2.Add(1);
  wg3.Add(int(p));
  wg4.Add(int(p));

  var mainThreadJob Job;

  for i := range(p) {
    // if we want to take extras
    sampleAdjust := uint32(0);
    if leftOver & i > leftOver {
      sampleAdjust = 1;
    }

    // if we want to adjust the start because we took extras
    startAdjust := uint32(0);
    if leftOver & i > 0 {
      startAdjust = min(i, leftOver)
    }
      
      start := i * sampleSize + startAdjust;
      finish := i * sampleSize + sampleSize + sampleAdjust;
      threadJob := Job{list[start: finish],
        w,
        p, 
        pivots, 
        gatheredSample[i*p:i*p+p],
        partitions[p*i:p*i+p],
        &wg1,
        &wg2,
        &wg3,
        &wg4}
    if i < p-1 {
      go workerJob(threadJob);
    } else {
      mainThreadJob = threadJob 
    }
  }
  phase1(mainThreadJob.sample, 
    mainThreadJob.gatheredSample, 
    mainThreadJob.w, 
    mainThreadJob.p, 
    mainThreadJob.wg1);

  wg1.Wait();

  phase2(&gatheredSample, pivots, p, &wg2);
  wg2.Wait();

  phase3(mainThreadJob.sample, pivots, mainThreadJob.partitions, mainThreadJob.wg3);

  wg3.Wait();

  phase4(destination, partitions...);
}

func workerJob(job Job) { 
  // witing for phase 1 to finish
  phase1(job.sample, job.gatheredSample, job.w, job.p, job.wg1);
  job.wg1.Wait();

  // waiting for main thread to finish phase2
  job.wg2.Wait();

  phase3(job.sample,job.pivots, job.partitions, job.wg3);

  job.wg3.Wait();

}

func phase1(localSample []uint32, sampleDest []uint32, w uint32, p uint32, wg *sync.WaitGroup) {
  defer wg.Done();

  sort.Slice(localSample, func(i, j int) bool {
    return localSample[i] < localSample[j]
  });

  for i:= uint32(0) ; i < p ; i++ {
    sampleDest[int(i)] = localSample[i * w];
  }
}

func phase2(gatheredSample *[]uint32, pivots []uint32, p uint32, wg *sync.WaitGroup) {
  defer wg.Done();

  sort.Slice(*gatheredSample, func(i, j int) bool {
    return (*gatheredSample)[i] < (*gatheredSample)[j];
  });

  for i := range(pivots) {
    pivots[i] = (*gatheredSample)[int(p) + i * int(p)];
  }
}

func phase3(localSample []uint32, pivots []uint32, partitions [][]uint32, wg *sync.WaitGroup) {
  defer wg.Done();
  sampleStart := localSample;

  // partition[i] contains all elements in the local sample less than pivots[i]
  for partitionIndex := range(pivots) {
    index := BinarySearch(sampleStart, pivots[partitionIndex]);

    partitions[partitionIndex] = sampleStart[:index];
 
    // cut down search space, next pivot is bigger.
    sampleStart = sampleStart[index:];
  }

  // the last partition contains everything else
  partitions[len(pivots)] = sampleStart;
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
