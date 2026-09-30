// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestIntFrom96DPI(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		dpi      int
		expected int
	}{
		{"96 DPI (100%)", 100, 96, 100},
		{"120 DPI (125%)", 100, 120, 125},
		{"144 DPI (150%)", 100, 144, 150},
		{"192 DPI (200%)", 100, 192, 200},
		{"Rounding up", 10, 120, 13},
		{"Rounding down", 10, 110, 11},
		{"Negative value", -100, 144, -150},
		{"Zero value", 0, 144, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt := IntFrom96DPI(tt.value, tt.dpi)
			if gotInt != tt.expected {
				t.Errorf("IntFrom96DPI[int](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt, tt.expected)
			}

			gotInt32 := IntFrom96DPI(int32(tt.value), tt.dpi)
			if gotInt32 != int32(tt.expected) {
				t.Errorf("IntFrom96DPI[int32](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt32, tt.expected)
			}

			gotInt64 := IntFrom96DPI(int64(tt.value), tt.dpi)
			if gotInt64 != int64(tt.expected) {
				t.Errorf("IntFrom96DPI[int64](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt64, tt.expected)
			}
		})
	}
}

func TestIntTo96DPI(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		dpi      int
		expected int
	}{
		{"96 DPI (100%)", 100, 96, 100},
		{"120 DPI (125%)", 125, 120, 100},
		{"144 DPI (150%)", 150, 144, 100},
		{"192 DPI (200%)", 200, 192, 100},
		{"Rounding up", 13, 120, 10},
		{"Rounding down", 11, 110, 10},
		{"Negative value", -150, 144, -100},
		{"Zero value", 0, 144, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInt := IntTo96DPI(tt.value, tt.dpi)
			if gotInt != tt.expected {
				t.Errorf("IntTo96DPI[int](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt, tt.expected)
			}

			gotInt32 := IntTo96DPI(int32(tt.value), tt.dpi)
			if gotInt32 != int32(tt.expected) {
				t.Errorf("IntTo96DPI[int32](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt32, tt.expected)
			}

			gotInt64 := IntTo96DPI(int64(tt.value), tt.dpi)
			if gotInt64 != int64(tt.expected) {
				t.Errorf("IntTo96DPI[int64](%d, %d) = %d; want %d", tt.value, tt.dpi, gotInt64, tt.expected)
			}
		})
	}
}
