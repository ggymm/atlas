package uuid

import (
	"testing"
)

func Test_NewUUID(t *testing.T) {
	t.Log(New())
}

func BenchmarkNewUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}
