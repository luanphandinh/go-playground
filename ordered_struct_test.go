package go_tuning_examples

import (
	"fmt"
	"testing"
)

const STRUCT_LENGTH = 10000

type badStruct struct {
	ItemId   uint64
	ModelId  uint64
	RuleType uint32
	RuleId   uint64
	Flag     uint8
}

type goodStruct struct {
	Flag     uint8
	RuleType uint32
	RuleId   uint64
	ItemId   uint64
	ModelId  uint64
}

var goodStructArr []goodStruct
var badStructArr []badStruct

func initGoodStructArray() {
	fmt.Println("Init good struct array")
	goodStructArr = make([]goodStruct, STRUCT_LENGTH)
	for i := 0; i < len(goodStructArr); i++ {
		goodStructArr[i] = goodStruct{
			ItemId:   10,
			ModelId:  11,
			RuleId:   12,
			RuleType: 13,
			Flag:     1,
		}
	}
}

func initBadStructArray() {
	fmt.Println("Init bad struct array")
	badStructArr = make([]badStruct, STRUCT_LENGTH)
	for i := 0; i < len(badStructArr); i++ {
		badStructArr[i] = badStruct{
			ItemId:   10,
			ModelId:  11,
			RuleId:   12,
			RuleType: 13,
			Flag:     1,
		}
	}
}

func badStructTraverse() uint64 {
	count := uint64(0)
	for _, badSt := range badStructArr {
		count += badSt.RuleId
	}

	return count
}

func goodStructTraverse() uint64 {
	count := uint64(0)
	for _, goodSt := range goodStructArr {
		count += goodSt.RuleId
	}

	return count
}

func BenchmarkBadStructTraverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		badStructTraverse()
	}
}

func BenchmarkGoodStructTraverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goodStructTraverse()
	}
}
