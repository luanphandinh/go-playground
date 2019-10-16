package go_tuning_examples

import (
	"fmt"
	"testing"
)

func doConcat() string {
	return "This" + "is" + "simple" + "concat" + "string" + "benchmark" + "testing"
}

func doFormat() string {
	return fmt.Sprintf("%s %s %s %s %s %s %s", "This", "is", "simple", "format", "string", "benchmark", "testing")
}

func doFormatMultiple() string {
	p1 := fmt.Sprintf("%s", "This")
	p2 := fmt.Sprintf("%s", "is")
	p3 := fmt.Sprintf("%s", "simple")
	p4 := fmt.Sprintf("%s %s", "format", "string")
	p5 := fmt.Sprintf("%s %s", "benchmark", "testing")

	return p1 + p2 + p3 + p4 + p5
}

func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doConcat()
	}
}

func BenchmarkConcatStringFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFormat()
	}
}

func BenchmarkConcatStringFormatMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFormatMultiple()
	}
}
