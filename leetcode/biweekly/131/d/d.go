package main

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math/bits"
	"slices"
)

// https://space.bilibili.com/206214
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

type uf []int

func (f uf) find(x int) int {
	if f[x] != x {
		f[x] = f.find(f[x])
	}
	return f[x]
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			pos = append(pos, q[1])
		}
	}
	m++

	left := make(uf, m+1)
	right := make(uf, m+1)
	for i := range left {
		left[i] = i
		right[i] = i
	}
	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		p, q := pos[i-1], pos[i]
		t.update(q, q-p)
		for j := p + 1; j < q; j++ {
			left[j] = p
			right[j] = q
		}
	}
	for j := pos[len(pos)-1] + 1; j < m; j++ {
		left[j] = pos[len(pos)-1]
		right[j] = m
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre := left.find(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			left[x] = x - 1
			right[x] = x + 1
			nxt := right.find(x)   // x 右侧最近障碍物的位置
			t.update(nxt, nxt-pre) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre), x-pre)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}

//

type seg []int

// 把 i 处的值改成 val
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o] = val
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 查询 [0,R] 中的最大值
func (t seg) query(o, l, r, R int) int {
	if r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, R)
	}
	return max(t[o<<1], t.query(o<<1|1, m+1, r, R))
}

func getResults2(queries [][]int) (ans []bool) {
	m := 0
	for _, q := range queries {
		if q[0] == 1 {
			m = max(m, q[1])
		}
	}
	m++

	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{}) // 哨兵
	set.Put(m, struct{}{})
	t := make(seg, 2<<bits.Len(uint(m)))

	for _, q := range queries {
		x := q[1]
		pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			set.Put(x, struct{}{})
			t.update(1, 0, m, x, x-pre.Key)       // 更新 d[x] = x - pre
			t.update(1, 0, m, nxt.Key, nxt.Key-x) // 更新 d[nxt] = nxt - x
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.query(1, 0, m, pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	return
}
