package algo

import "sort"

func psrs(list *[]uint32) {

}

/*
sort the source (which is a limited sample, then output some amount of 
random samples to the sample destination)
*/
func phase1(src []uint32, sampleDest []uint32, w uint32, p uint32) {
  sort.Slice(src, func(i, j int) bool {
    return i < j
  });

  for i:= uint32(0) ; i < p ; i++ {
    sampleDest[int(i)] = src[i * w];
  }
}
