package util

import (
	"math/rand"
	"sort"
)

type Dataset struct {
  List []uint32
}

/*
  specifically for testing only
*/
func NewDataset(n uint64) *Dataset {
  newData := Dataset{};

  newData.generateDataset(n)

  return &newData
}

func NewDatasetWithSeed(n uint64, seed int64) *Dataset {
  rand.Seed(seed)
  return NewDataset(n)
}

/*
  check that our list is sorted. 
*/
func (data Dataset) VerifyAlgorithm(sortFunc func(*[]uint32) ) {

  controlTemp := make([]uint32, len(data.List));
  testTemp := make([]uint32, len(data.List));

  copy(controlTemp, data.List);

  sort.Slice(controlTemp, func(i, j int) bool {
    return i < j;
  });
 
  sortFunc(&testTemp);
}

func (data *Dataset) generateDataset(n uint64) {
  data.List = make([]uint32, n)

  for index := range(data.List) {
    data.List[index] = rand.Uint32()
  }
}
