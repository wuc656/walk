//go:build windows

package declarative

import "testing"

func BenchmarkBuilderDefer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		builder := NewBuilder(nil)
		// A typical number of defers
		for j := 0; j < 6; j++ {
			builder.Defer(func() error { return nil })
		}
	}
}
