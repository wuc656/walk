//go:build windows

package declarative

import "testing"

type builderBaseline struct {
	deferredFuncs []func() error
}

func (b *builderBaseline) Defer(f func() error) {
	b.deferredFuncs = append(b.deferredFuncs, f)
}

type builderProposed struct {
	deferredFuncs []func() error
}

// 8 is what we are proposing and testing
func (b *builderProposed) Defer(f func() error) {
	if b.deferredFuncs == nil {
		b.deferredFuncs = make([]func() error, 0, 8)
	}
	b.deferredFuncs = append(b.deferredFuncs, f)
}

func benchmarkBaselineN(b *testing.B, n int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		builder := &builderBaseline{}
		for j := 0; j < n; j++ {
			builder.Defer(func() error { return nil })
		}
	}
}

func benchmarkProposedN(b *testing.B, n int) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		builder := &builderProposed{}
		for j := 0; j < n; j++ {
			builder.Defer(func() error { return nil })
		}
	}
}

func BenchmarkDeferBaseline_1(b *testing.B)  { benchmarkBaselineN(b, 1) }
func BenchmarkDeferBaseline_2(b *testing.B)  { benchmarkBaselineN(b, 2) }
func BenchmarkDeferBaseline_4(b *testing.B)  { benchmarkBaselineN(b, 4) }
func BenchmarkDeferBaseline_6(b *testing.B)  { benchmarkBaselineN(b, 6) }
func BenchmarkDeferBaseline_8(b *testing.B)  { benchmarkBaselineN(b, 8) }
func BenchmarkDeferBaseline_16(b *testing.B) { benchmarkBaselineN(b, 16) }
func BenchmarkDeferBaseline_32(b *testing.B) { benchmarkBaselineN(b, 32) }
func BenchmarkDeferBaseline_64(b *testing.B) { benchmarkBaselineN(b, 64) }

func BenchmarkDeferProposed_1(b *testing.B)  { benchmarkProposedN(b, 1) }
func BenchmarkDeferProposed_2(b *testing.B)  { benchmarkProposedN(b, 2) }
func BenchmarkDeferProposed_4(b *testing.B)  { benchmarkProposedN(b, 4) }
func BenchmarkDeferProposed_6(b *testing.B)  { benchmarkProposedN(b, 6) }
func BenchmarkDeferProposed_8(b *testing.B)  { benchmarkProposedN(b, 8) }
func BenchmarkDeferProposed_16(b *testing.B) { benchmarkProposedN(b, 16) }
func BenchmarkDeferProposed_32(b *testing.B) { benchmarkProposedN(b, 32) }
func BenchmarkDeferProposed_64(b *testing.B) { benchmarkProposedN(b, 64) }
