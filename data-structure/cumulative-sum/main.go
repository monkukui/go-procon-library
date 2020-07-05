package main

import (
  "fmt"
  "os"
  "bufio"
)

// begin
type CumulativeSum struct {
  Array []int64
}

func NewCumulativeSum(v []int64) *CumulativeSum {
  cs := new(CumulativeSum)
  cs.Array = make([]int64, len(v) + 1)
  for i := 1; i < len(cs.Array); i++ {
    cs.Array[i] += cs.Array[i - 1] + v[i - 1]
  }
  return cs;
}

// [l, r) 0-indexed
func (cs CumulativeSum) get(l int, r int) int64 {
  return cs.Array[r] - cs.Array[l]
}
// end

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func nextInt() int {
	in.Scan()
	r := int(0)
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}


// Static Range Sum（https://judge.yosupo.jp/problem/static_range_sum）
func solve() {

  in.Split(bufio.ScanWords)
  n := nextInt()
  q := nextInt()
  arr := make([]int64, n)
  for i := 0; i < n; i++ {
    arr[i] = int64(nextInt())
  }
  cs := NewCumulativeSum(arr)

  for i := 0; i < q; i++ {
    l := nextInt()
    r := nextInt()
    fmt.Fprintf(out, "%d\n", cs.get(l, r))
  }
  out.Flush()
}

func main() {
  solve()
}
