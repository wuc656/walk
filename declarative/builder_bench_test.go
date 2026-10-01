//go:build windows

package declarative

import "testing"

func benchmarkDeferN(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		builder := NewBuilder(nil)
		for j := 0; j < n; j++ {
			builder.Defer(func() error { return nil })
		}
	}
}

func BenchmarkDefer_1(b *testing.B)  { benchmarkDeferN(b, 1) }
func BenchmarkDefer_4(b *testing.B)  { benchmarkDeferN(b, 4) }
func BenchmarkDefer_8(b *testing.B)  { benchmarkDeferN(b, 8) }
func BenchmarkDefer_16(b *testing.B) { benchmarkDeferN(b, 16) }
func BenchmarkDefer_32(b *testing.B) { benchmarkDeferN(b, 32) }
func BenchmarkDefer_64(b *testing.B) { benchmarkDeferN(b, 64) }

func BenchmarkDefer_Realistic(b *testing.B) {
    // Simulate a typical declarative UI building workload
    // Usually it has 4-10 defers per builder instance for actions, menus, data binders.
    for i := 0; i < b.N; i++ {
		builder := NewBuilder(nil)
		// E.g. menu actions, layout setups
		for j := 0; j < 6; j++ {
			builder.Defer(func() error { return nil })
		}
	}
}
