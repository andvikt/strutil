package helpers

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRune2S(t *testing.T) {
	s := "hello"
	r, sz := utf8.DecodeRuneInString(s)
	s = Rune2S(r, sz)
	fmt.Println(s)
}

func BenchmarkRune2S(b *testing.B) {
	s := "hello"
	r, sz := utf8.DecodeRuneInString(s)
	for i := 0; i < b.N; i++ {
		Rune2S(r, sz)
	}
}

func TestCountPref(t *testing.T) {
	a := []byte("нор")
	b := []byte("нариман")

	cnt := CountPrefix(a, b)
	fmt.Println(cnt)
}

func BenchmarkXxx(b *testing.B) {
	a := []byte("андрей")
	bb := []byte("антон")
	for i := 0; i < b.N; i++ {
		CountPrefix(a, bb)
	}
}
