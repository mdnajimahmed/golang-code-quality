package service

import (
	"github.com/mandatorySuicide/golang-code-quality/src/utility"
	"sort"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	elements := utility.MakeRandInt(10000)
	BubbleSort(elements, CompAsc)
}

func BenchmarkMergeSort(b *testing.B) {
	elements := utility.MakeRandInt(10000)
	MergeSort(elements, CompAsc)
}

func BenchmarkGoSort(b *testing.B) {
	elements := utility.MakeRandInt(10000)
	sort.Ints(elements)
}
