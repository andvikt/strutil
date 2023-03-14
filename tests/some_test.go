package tests

import (
	"testing"

	"github.com/andvikt/strutil/metrics"
)

func BenchmarkDelDist(b *testing.B) {
	bb := []rune("сулийм")
	a := []rune("сулейманов")
	swg := metrics.NewSmithWatermanGotoh()
	for i := 0; i < b.N; i++ {
		swg.Compare(a, bb)
	}
}

// func BenchmarkPool(b *testing.B) {
// 	f := func() {
// 		v1 := metrics.NewDistVec(10)
// 		(*v1) = (*v1)[:0]
// 		metrics.DistPool.Put(v1)
// 	}
// 	for i := 0; i < b.N; i++ {
// 		f()
// 	}
// }
