package go_tuning_examples

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initTraverseTestArray()
	initBadStructArray()
	initGoodStructArray()

	exitVal := m.Run()

	os.Exit(exitVal)
}
