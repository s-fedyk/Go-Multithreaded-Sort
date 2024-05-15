package util

import (
	"testing"
)

func TestGenerateDataset(t *testing.T) {
  resultLength := len(NewDataset(22).List);

  expectedLength := 22;

  if resultLength != expectedLength {
    t.Error("Generating wrong length list");
  }
}
