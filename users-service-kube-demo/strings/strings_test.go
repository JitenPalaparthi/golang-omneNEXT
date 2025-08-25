package strings_test

import (
	"testing"
	"users-service/strings"
)

func BenchmarkGetBigString(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_ = strings.GetBigString(10000)
	}
}

func BenchmarkGetBigStringBuilder(b *testing.B) {
	for i := 1; i < b.N; i++ {
		_ = strings.GetBigStringBuilder(10000)
	}
}
