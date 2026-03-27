package tests

import (
	"testing"
)

func init() {
	InitTestEnv("sample_test")
}

func TestSumOnePlusOne(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("expected 1 + 1 to equal 2")
	}
}
