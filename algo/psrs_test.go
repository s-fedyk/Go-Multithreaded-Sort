package algo

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
	"sync"
	"testing"
)

func TestPhase1(t *testing.T) {
  wg := sync.WaitGroup{};

  resultSampleBlock := []uint32 {0, 16, 27};
  resultLocalBlock := []uint32 {0, 1, 2, 9, 16, 17, 24, 25, 27, 28, 30, 33};

  n := uint32(36);
  p := uint32(3);
  w := n / (p * p) ;

  sampleBlock := make([]uint32, 3, 3);
  localBlock := []uint32 {16, 2, 17, 24, 33, 28, 30, 1, 0, 27, 9, 25};

  wg.Add(1);
  phase1(localBlock, sampleBlock, w, p, &wg);
  wg.Wait()


  if !slices.Equal(resultLocalBlock, localBlock) {
    t.Error("Phase 1 local block sorting broken");
  }

  if !slices.Equal(resultSampleBlock, sampleBlock) {
    t.Error("Phase 1 sampler broken");
  }
}

func TestPhase2(t *testing.T) {
  p := uint32(3);

  wg := sync.WaitGroup{};

  gatheredRegularSample := []uint32 {0, 16, 27, 7, 13, 23, 3, 10, 22};
  expectedPivots := []uint32 {10, 22};
  pivots := make([]uint32, p-1);

  wg.Add(1);
  phase2(&gatheredRegularSample, pivots, p, &wg)
  wg.Wait()

  if !slices.Equal(expectedPivots, pivots) {
    t.Error("Phase 2 produced bad pivots");
  }
}

func TestPhase3(t *testing.T) {
  localBlock := []uint32 {0, 1, 2, 9, 16, 17, 24, 25, 27, 28, 30, 33};
  pivots := []uint32 {10, 22};
  partitions := make([][]uint32, len(pivots)+1);
  wg := sync.WaitGroup{};
  wg.Add(1)
  phase3(localBlock, pivots, partitions, &wg);
  wg.Wait();

  expected := [][]uint32 {{0,1,2,9}, {16,17}, {24,25,27,28,30,33}};

  if !reflect.DeepEqual(partitions, expected) {
    t.Error("Partitioning failure");
  }
}

func TestBinarySearch(t *testing.T) {
  testlist1 := []uint32 {1, 2, 3, 7, 8, 9};
  testlist2 := []uint32 {3, 4, 5, 6, 7};
  testlist3 := []uint32 {3, 3, 3, 3, 3, 5, 8};

  result := BinarySearch(testlist1, 6)
  if result != 3 {
    t.Errorf("Search failed #1, wanted 2, got %d", result);
  }

  result = BinarySearch(testlist2, 2)
  if result != 0 {
    t.Errorf("Search failed #2, wanted 0, got %d", result);
  }

  result = BinarySearch(testlist2, 8)
  if result != 5 {
    t.Errorf("Search failed #3, wanted 8, got %d", result);
  }

  result = BinarySearch(testlist2, 5)
  if result != 3 {
    t.Errorf("Search failed #4, wanted 3, got %d", result);
  }

  result = BinarySearch(testlist3, 3)
  if result != 5 {
    t.Errorf("Search failed #5, wanted 5, got %d", result);
  }
}

func TestPhase4(t *testing.T) {
  partition1 := []uint32 {1, 2, 3, 4, 5, 6, 7};
  partition2 := []uint32 {5, 5, 5};
  partition3 := []uint32 {1, 23, 27};

  outputSize := len(partition1) + len(partition2) + len(partition3);
  output := make([]uint32, outputSize);

  phase4(output, partition1, partition2, partition3);

  if !slices.IsSorted(output) || len(output) != outputSize {
    t.Errorf("Phase 4 output list not sorted or of incorrect length")
  }
}

func TestPsrs(t *testing.T) {
  mlist := []uint32 {
    16,2,17,24,33,28,30,1,0,27,9,25,34,23,19,18,11,7,21,13,8,35,12,29,6,3,4,14,22,15,32,10,26,31,20,5,
  };

  expected := make([]uint32, len(mlist));
  copy(expected, mlist);

  sort.Slice(expected, func(i, j int) bool {
    return expected[i] < expected[j]
  });

  dest := make([]uint32, len(mlist));
  fmt.Println(len(mlist));

  psrs(mlist, 3, dest);

  if !slices.Equal(dest, expected) {
    t.Errorf("Psrs output is bad")
  }
}
