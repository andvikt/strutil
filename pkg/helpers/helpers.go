package helpers

import (
	"log"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"

	"golang.org/x/exp/constraints"
)

func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var bi strings.Builder

func Rune2S(b rune, s int) (ret string) {
	bi.WriteRune(b)
	ret = bi.String()
	bi.Reset()
	return ret
}

func S2B(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// считаем кол-во совпаших символов с начала строки
func CountPrefix(a []byte, b []byte) int {
	var ret int
	l := Min(len(a), len(b))
	for i := 0; i < l; {
		ra, s := utf8.DecodeRune(a[i:])
		rb, s := utf8.DecodeRune(b[i:])
		i += s
		if ra == rb {
			ret++
		} else {
			return ret
		}
	}
	return ret
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
