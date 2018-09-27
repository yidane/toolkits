package shortID

import (
	"fmt"
	"testing"
)

func TestShortID_Generate(t *testing.T) {
	shortID := New()
	fmt.Println(shortID.Generate())
	t.Error(shortID.Generate())
}

func BenchmarkShortIDGenerate(b *testing.B) {
	shortID := New()

	for i := 0; i < b.N; i++ {
		shortID.Generate()
	}
}
