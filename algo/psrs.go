package algo

import (
	"sort"
)

func psrs(list *[]uint32) {

}

/*
sort the source (which is a limited sample, then output some amount of 
random samples to the sample destination)
*/
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

  // partition the passed in slice into len(pivot) partitions
  for partitionIndex := range(pivots) {
    index := BinarySearch(sampleStart, pivots[partitionIndex])

    partitions[partitionIndex] = (*sampleStart)[:index]
 
    // don't have to search again. Next pivot will be bigger
    (*sampleStart) = (*sampleStart)[index:]
  }

  // the last partition contains everything else
  partitions[len(pivots)] = (*sampleStart)

  return &partitions
}

