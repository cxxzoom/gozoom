package main

import (
	"bytes"
	"fmt"
)

// 这是个存放int集合的算法
// 根据暂时是64位的无符号
// 代码讲解：
//  相当于是 words的每个key都对应了一个桶，桶里存放的数据64个二进制位
// 打个比方： words[0] 是 存放的0~63，然后 1 << (uint(x % 64)) 是为了保证0~63每个都独占一个二进制的位
// 然后 通过 & 来判断是否存在

const plat = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/plat, uint(x%plat)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/plat, uint(x%plat)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll batch add
func (s *IntSet) AddAll(xx ...int) {
	for _, v := range xx {
		if s.Has(v) {
			continue
		}

		s.Add(v)
	}
}

// UnionWith A u B
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tWord := range t.words {
		if i < len(s.words) {
			s.words[i] |= tWord
		} else {
			s.words = append(s.words, tWord)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < plat; j++ {
			// 重点是下面这句，怎么存的就怎么判断在不在
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", plat*i+j)
			}

		}

	}
	buf.WriteByte('}')
	return buf.String()
}

// Len 和String类似
func (s *IntSet) Len() int {
	var lens int
	// 和string类似，都是循环然后判断
	for word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < plat; j++ {
			if word&(1<<(uint(j))) != 0 {
				lens++
			}
		}
	}
	return lens
}

// Remove 先用has判断是否存在，不存在直接返回
// 如果存在再考虑删除的事
func (s *IntSet) Remove(x int) {
	if !s.Has(x) {
		return
	}
	// 判断是第几个桶，然后remove
	bucket, bit := x/plat, uint(x%plat)
	s.words[bucket] ^= 1 << bit
}

// Clear all elements
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy return s.words.copy
func (s *IntSet) Copy() *IntSet {
	f := &IntSet{make([]uint, len(s.words))}
	copy(f.words, s.words)
	return f
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if word == 0 {
			continue
		}

		if s.words[i] == 0 {
			s.words[i] = word
			continue
		}

		s.words[i] &= word
	}
}

// DifferentWith 差集
func (s *IntSet) DifferentWith(t *IntSet) {
	// 相同为0
	for i, word := range t.words {

		if word == 0 {
			continue
		}

		if s.words[i] != 0 {
			s.words[i] ^= word
		}
	}
}

// SymmetricDifference 并差集
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, word := range t.words {
		if word == 0 {
			continue
		}

		if i >= len(s.words) {
			s.words = append(s.words, word)
			continue
		}

		if s.words[i] != 0 {
			s.words[i] ^= word
		} else {
			s.words[i] |= word
		}

	}
}

func (s *IntSet) Elems() []int {
	var x []int
	if len(s.words) == 0 {
		return x
	}

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < plat; j++ {
			// 重点是下面这句，怎么存的就怎么判断在不在
			if word&(1<<uint(j)) != 0 {
				x = append(x, plat*i+j)
			}

		}
	}

	return x
}

func main() {
	s := IntSet{make([]uint, 0)}
	fmt.Println(s.Has(1))

	s.Add(2)
	fmt.Println(s)
	s.Add(1)
	fmt.Println(s)
	fmt.Println(s.Has(2))
	fmt.Println(s.Has(1))
	// fmt.Println(s.Has(3))

	f := IntSet{make([]uint, 0)}
	f.Add(1)
	f.Add(3)
	s.UnionWith(&f)
	fmt.Println(s)
	fmt.Println(s.String())
	s.Remove(1)
	fmt.Println(s.String())
	t := s.Copy()
	s.Clear()
	fmt.Println(s.String())
	fmt.Println(t.String())
	s.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s.String())
	s.Add(10)
	fmt.Println(s.String())
	fmt.Println(s.Has(10))

	fmt.Println(t)
	t.IntersectWith(&s)
	fmt.Println(t.String())
	fmt.Println(s.String())
	// s.IntersectWith(t)
	// s.DifferentWith(t)
	t.AddAll(11, 12, 66, 99)
	fmt.Println("====================================")
	fmt.Println(s.String())
	fmt.Println(t.String())
	fmt.Println("====================================")
	s.Add(64)
	s.SymmetricDifference(t)
	fmt.Println(s.String())
	fmt.Println(s.Elems())
}
