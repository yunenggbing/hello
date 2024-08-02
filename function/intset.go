package function

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64+i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet) Len() (i int) {

	if s == nil {
		return i
	}
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				i++
			}
		}
	}
	return i
}

// remove x from the set
func (src *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	// 判断word是否越界
	if word < len(src.words) {
		// 如果没有越界, 则将word的第bit位置为0
		src.words[word] &= ^(1 << bit)
	}
}

// remove all elements from the set
func (src *IntSet) Clear() {
	src.words = nil
}

// return a copy of the set
func (src *IntSet) Copy() *IntSet {
	var des IntSet
	if src == nil {
		return nil
	}

	des.words = append(des.words, src.words...)
	return &des
}
func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())
	fmt.Println(x.Copy())
	fmt.Println("------------")
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())
	fmt.Println("------------")
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

}
