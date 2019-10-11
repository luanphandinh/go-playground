package go_playground

import (
	"fmt"
	"testing"
)

func doConcat() string {
	return "This" + "is" + "simple" + "concat" + "string"
}

func doFormat() string {
	return fmt.Sprintf("%s %s %s %s %s", "This", "is", "simple", "format", "string")
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
