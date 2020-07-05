# 累積和（Cumulative Sum）

## 概要
数列に対し，区間 [l, r) の総和をO(1)で求めることができる．

## 使い方

- `NewCumulativeSum(v []int64) *CumulativeSum`：`[]int64`を引数として，累積和を初期化（O(len(v))）
- `get(l int, r int) int64`：区間 [l, r) の総和を取得する（O(1)）

