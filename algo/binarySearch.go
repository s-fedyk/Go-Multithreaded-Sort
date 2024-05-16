package algo

/*
Get index of the first item that is greater than the target
*/
func BinarySearch(list *[]uint32, target uint32) int64 {

  low := int64(0)
  high := int64(len(*list)) - 1

  var middle int64;

  for (low < high) {
    middle = (low + high) / 2

    if ((*list)[middle] > target){
      high = middle - 1
    } else {
      low = middle + 1
    }
  }

  // largest item in the list not smaller than taget. Go over
  if (*list)[low] <= target {
    low += 1
  }

  return low
}
