package main

import (
	"fmt"
	"os"
	"strconv"
  "multithreaded_sort/util"
)
func myfunc(a *[]uint32)  {
  fmt.Println("ah");
}

func main() {
  fmt.Print(os.Args)

  n, err := strconv.ParseUint(os.Args[1], 10, 64);

  if err != nil {
    fmt.Print(err.Error())
  }

  p, err := strconv.ParseUint(os.Args[2], 10, 64);

  fmt.Print(p, n)
  myData := util.NewDatasetWithSeed(n, 10)

  for i := range(myData.List) {
    fmt.Println(myData.List[i])
  }
}


