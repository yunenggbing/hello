package function

import (
	"errors"
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	//for i, course := range topoSort(prereqs) {
	//	fmt.Printf("%d:\t%s\n", i+1, course)
	//}
	fmt.Println(min(1, 3, 42, 23))
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func min(vars ...int) (int, error) {
	if len(vars) == 0 {
		return 0, errors.New("no variables provided")
	}
	i := vars[0]
	for _, x := range vars {
		if x < i {
			i = x
		}
	}
	return i, nil
}
