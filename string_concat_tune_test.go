package go_tuning_examples

import (
	"fmt"
	"testing"
)

func doConcat(str1 string, str2 string, str3 string) string {
	a := str1
	a += str2
	a += str3

	return a
}

func doByteConcat(str1 string, str2 string, str3 string) string {
	a := make([]byte, 0)
	a = append(a, str1...)
	a = append(a, str2...)
	a = append(a, str3...)

	return string(a)
}

func doByteConcatWithKnownLength(str1 string, str2 string, str3 string) string {
	length := len(str1) + len(str2) + len(str3)
	a := make([]byte, 0, length)
	a = append(a, str1...)
	a = append(a, str2...)
	a = append(a, str3...)

	return string(a)
}

func doFormat(str1 string, str2 string, str3 string) string {
	return fmt.Sprintf("%s %s %s", str1, str2, str3)
}

func doFormatMultiple(str1 string, str2 string, str3 string) string {
	p1 := fmt.Sprintf("%s", str1)
	p2 := fmt.Sprintf("%s", str2)
	p3 := fmt.Sprintf("%s", str3)
	return p1 + p2 + p3
}

func BenchmarkConcatString(b *testing.B) {
	str1 := "This"
	str2 := "is"
	str3 := "simple"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		doConcat(str1, str2, str3)
	}
}

func BenchmarkConcatStringFormat(b *testing.B) {
	str1 := "This"
	str2 := "is"
	str3 := "simple"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		doFormat(str1, str2, str3)
	}
}

func BenchmarkConcatStringFormatMultiple(b *testing.B) {
	str1 := "This"
	str2 := "is"
	str3 := "simple"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		doFormatMultiple(str1, str2, str3)
	}
}

func BenchmarkConcatStringByBytes(b *testing.B) {
	str1 := "This"
	str2 := "is"
	str3 := "simple"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		doByteConcat(str1, str2, str3)
	}
}

func BenchmarkConcatStringByBytesKnownLength(b *testing.B) {
	str1 := "This"
	str2 := "is"
	str3 := "simple"
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		doByteConcatWithKnownLength(str1, str2, str3)
	}
}
