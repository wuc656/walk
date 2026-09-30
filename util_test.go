// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"reflect"
	"testing"
)

func TestMarginsTo96DPI(t *testing.T) {
	tests := []struct {
		name     string
		input    Margins
		dpi      int
		expected Margins
	}{
		{
			name:     "96 DPI - 1:1 scale",
			input:    Margins{96, 96, 96, 96},
			dpi:      96,
			expected: Margins{96, 96, 96, 96},
		},
		{
			name:     "192 DPI - 2:1 scale (half size in 96DPI)",
			input:    Margins{192, 192, 192, 192},
			dpi:      192,
			expected: Margins{96, 96, 96, 96},
		},
		{
			name:     "144 DPI - 1.5:1 scale",
			input:    Margins{144, 144, 144, 144},
			dpi:      144,
			expected: Margins{96, 96, 96, 96},
		},
		{
			name:     "96 DPI - irregular margins",
			input:    Margins{10, 20, 30, 40},
			dpi:      96,
			expected: Margins{10, 20, 30, 40},
		},
		{
			name:     "144 DPI - irregular margins",
			input:    Margins{15, 30, 45, 60},
			dpi:      144,
			expected: Margins{10, 20, 30, 40},
		},
		{
			name:     "0 margins",
			input:    Margins{0, 0, 0, 0},
			dpi:      120,
			expected: Margins{0, 0, 0, 0},
		},
		{
			name:     "Rounding check - rounds up",
			input:    Margins{11, 11, 11, 11},
			dpi:      144,
			// 11 * 96 / 144 = 11 / 1.5 = 7.333 -> 7
			expected: Margins{7, 7, 7, 7},
		},
		{
			name:     "Rounding check - rounds up to next",
			input:    Margins{14, 14, 14, 14},
			dpi:      144,
			// 14 * 96 / 144 = 14 / 1.5 = 9.333 -> 9
			expected: Margins{9, 9, 9, 9},
		},
		{
			name:     "Rounding check - round half up",
			input:    Margins{15, 15, 15, 15},
			dpi:      144,
			// 15 / 1.5 = 10
			expected: Margins{10, 10, 10, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarginsTo96DPI(tt.input, tt.dpi)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MarginsTo96DPI(%v, %d) = %v, want %v", tt.input, tt.dpi, got, tt.expected)
			}
		})
	}
}
