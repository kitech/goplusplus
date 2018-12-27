package gopp

import "sort"

func SortStrings(a []string) []string {
	sort.Strings(a)
	return a
}
