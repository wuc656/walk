// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestIntTo96DPI(t *testing.T) {
	tests := []struct {
		name  string
		value int
		dpi   int
		want  int
	}{
		{"96 DPI, zero", 0, 96, 0},
		{"96 DPI, positive", 10, 96, 10},
		{"96 DPI, negative", -10, 96, -10},

		{"192 DPI, positive even", 10, 192, 5},
		{"192 DPI, positive odd", 11, 192, 6},   // 11 * 0.5 = 5.5 -> 6
		{"192 DPI, negative odd", -11, 192, -6}, // -11 * 0.5 = -5.5 -> -6

		{"120 DPI, exact", 15, 120, 12},      // 15 * (96/120) = 15 * 0.8 = 12
		{"120 DPI, round down", 14, 120, 11}, // 14 * 0.8 = 11.2 -> 11
		{"120 DPI, round up", 16, 120, 13},   // 16 * 0.8 = 12.8 -> 13

		{"144 DPI, exact", 15, 144, 10},     // 15 * (96/144) = 15 * 0.666... = 10
		{"144 DPI, round down", 14, 144, 9}, // 14 * 0.666... = 9.333 -> 9
		{"144 DPI, round up", 16, 144, 11},  // 16 * 0.666... = 10.666 -> 11
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntTo96DPI(tt.value, tt.dpi); got != tt.want {
				t.Errorf("IntTo96DPI(%d, %d) = %d, want %d", tt.value, tt.dpi, got, tt.want)
			}
			// Test with int32
			if got := IntTo96DPI(int32(tt.value), tt.dpi); got != int32(tt.want) {
				t.Errorf("IntTo96DPI(int32(%d), %d) = %d, want %d", tt.value, tt.dpi, got, tt.want)
			}
		})
	}
}

func TestIntFrom96DPI(t *testing.T) {
	tests := []struct {
		name  string
		value int
		dpi   int
		want  int
	}{
		{"96 DPI, zero", 0, 96, 0},
		{"96 DPI, positive", 10, 96, 10},
		{"96 DPI, negative", -10, 96, -10},

		{"192 DPI, positive", 10, 192, 20},
		{"192 DPI, negative", -11, 192, -22},

		{"120 DPI, exact", 12, 120, 15},           // 12 * 1.25 = 15
		{"120 DPI, round up", 11, 120, 14},        // 11 * 1.25 = 13.75 -> 14
		{"120 DPI, round up half", 10, 120, 13},   // 10 * 1.25 = 12.5 -> 13
		{"120 DPI, negative half", -10, 120, -13}, // -10 * 1.25 = -12.5 -> -13

		{"144 DPI, exact", 10, 144, 15},           // 10 * 1.5 = 15
		{"144 DPI, round up half", 9, 144, 14},    // 9 * 1.5 = 13.5 -> 14
		{"144 DPI, negative half", -9, 144, -14},  // -9 * 1.5 = -13.5 -> -14
		{"144 DPI, round up half 2", 11, 144, 17}, // 11 * 1.5 = 16.5 -> 17
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntFrom96DPI(tt.value, tt.dpi); got != tt.want {
				t.Errorf("IntFrom96DPI(%d, %d) = %d, want %d", tt.value, tt.dpi, got, tt.want)
			}
			// Test with int32
			if got := IntFrom96DPI(int32(tt.value), tt.dpi); got != int32(tt.want) {
				t.Errorf("IntFrom96DPI(int32(%d), %d) = %d, want %d", tt.value, tt.dpi, got, tt.want)
			}
		})
	}
}
