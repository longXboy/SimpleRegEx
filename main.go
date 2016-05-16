package main

import (
	"fmt"
	"log"
	"regEx/dfs"
	"regEx/stack"
)

func main() {
	nfa := NFA{}
	nfa.Init("((AB|CD)*ZZZ)")
	fmt.Println(nfa.Test("ABCDCDZZZ"))
	fmt.Println(nfa.Test("ABCDCD"))
}

type NFA struct {
	marked []bool      //标记哪些走过
	pc     []int       //当前匹配到哪的状态
	re     []rune      //正则
	g      dfs.Digraph //有向图
	m      int
}

func (d *NFA) gen(g dfs.Digraph, v ...int) {
	d.pc = make([]int, 0)
	d.marked = make([]bool, g.V)
	for _, data := range v {
		d.dfs(g, data)
	}
}

func (d *NFA) dfs(g dfs.Digraph, v int) {
	for _, data := range g.GetV(v) {
		if data >= d.m {
			d.marked[data] = true
			d.pc = append(d.pc, data)
			continue
		}
		if d.re[data] == rune('(') || d.re[data] == rune('|') || d.re[data] == rune('*') || d.re[data] == rune(')') {
			if !d.marked[data] {
				d.dfs(g, data)
			}
		} else {
			if !d.marked[data] {
				d.marked[data] = true
				d.pc = append(d.pc, data)
			}
		}
	}

}

func (nfa *NFA) Test(str string) bool {
	nfa.gen(nfa.g, 0)
	for _, data := range []rune(str) {
		var match []int
		for _, v := range nfa.pc {
			if v >= nfa.m {
				continue
			}
			if nfa.re[v] == data {
				match = append(match, v)
			}
		}
		nfa.gen(nfa.g, match...)
	}
	for _, data := range nfa.pc {
		if data == nfa.m {
			return true
		}
	}
	return false
}

func (nfa *NFA) Init(reg string) {
	s := stack.Stack{}
	nfa.re = []rune(reg)
	nfa.m = len(nfa.re)
	nfa.g.Init(nfa.m + 1)
	for i := range nfa.re {
		lp := i
		data := nfa.re[i]
		if data == rune('(') || data == rune('|') {
			s.Push(i)
		} else if data == rune(')') {
			orstack := stack.Stack{}
			for {
				or, isfound := s.Pop()
				if isfound {
					if nfa.re[or] == rune('|') {
						nfa.g.AddEdge(or, i)
						orstack.Push(or)
					} else {
						lp = or
						break
					}
				} else {
					log.Fatal("parentheses(%d) not matched!", lp)
				}
			}
			for {
				or, isfound := orstack.Pop()
				if isfound {
					nfa.g.AddEdge(lp, or+1)
				} else {
					break
				}
			}
		}
		if i < (nfa.m-1) && nfa.re[i+1] == rune('*') {
			nfa.g.AddEdge(i+1, lp)
		}
		if data != rune('|') {
			nfa.g.AddEdge(i, i+1)
		}
	}
	fmt.Println(nfa.g.GetAdj())
}
