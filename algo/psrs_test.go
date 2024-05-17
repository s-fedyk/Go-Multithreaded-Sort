package algo

import (
	"fmt"
	"slices"
	"testing"
)

func TestPhase1(t *testing.T) {
  resultSampleBlock := []uint32 {0, 16, 27};
  resultLocalBlock := []uint32 {0, 1, 2, 9, 16, 17, 24, 25, 27, 28, 30, 33};

  n := uint32(36);
  p := uint32(3);
  w := n / (p * p) ;

  sampleBlock := make([]uint32, 3, 3);
  localBlock := []uint32 {16, 2, 17, 24, 33, 28, 30, 1, 0, 27, 9, 25};

  phase1(localBlock, sampleBlock, w, p);

  if !slices.Equal(resultLocalBlock, localBlock) {
    t.Error("Phase 1 local block sorting broken");
  }

  if !slices.Equal(resultSampleBlock, sampleBlock) {
    t.Error("Phase 1 sampler broken");
  }
}

func TestPhase2(t *testing.T) {
  p := uint32(3);

  gatheredRegularSample := []uint32 {0, 16, 27, 7, 13, 23, 3, 10, 22};
  expectedPivots := []uint32 {10, 22};

  resultPivots := phase2(&gatheredRegularSample, p)

  if !slices.Equal(expectedPivots, *resultPivots) {
    t.Error("Phase 2 produced bad pivots");
  }
}

func TestPhase3(t *testing.T) {
  localBlock := []uint32 {0, 1, 2, 9, 16, 17, 24, 25, 27, 28, 30, 33};
  pivots := []uint32 {10, 22};

  fmt.Println(*phase3(&localBlock, pivots))
}

func TestBinarySearch(t *testing.T) {
  testlist1 := []uint32 {1, 2, 3, 7, 8, 9};
  testlist2 := []uint32 {3, 4, 5, 6, 7};
  testlist3 := []uint32 {3, 3, 3, 3, 3, 5, 8};

  result := BinarySearch(&testlist1, 6)
  if result != 3 {
    t.Errorf("Search failed #1, wanted 2, got %d", result);
  }

  result = BinarySearch(&testlist2, 2)
  if result != 0 {
    t.Errorf("Search failed #2, wanted 0, got %d", result);
  }

  result = BinarySearch(&testlist2, 8)
  if result != 5 {
    t.Errorf("Search failed #3, wanted 8, got %d", result);
  }

  result = BinarySearch(&testlist2, 5)
  if result != 3 {
    t.Errorf("Search failed #4, wanted 3, got %d", result);
  }

  result = BinarySearch(&testlist3, 3)
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
