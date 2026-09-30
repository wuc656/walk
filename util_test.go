// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"reflect"
	"testing"
)

func TestMarginsFrom96DPI(t *testing.T) {
	tests := []struct {
		name     string
		input    Margins
		dpi      int
		expected Margins
	}{
		{
			name:     "96 DPI (Scale 1.0)",
			input:    Margins{HNear: 96, VNear: 48, HFar: 192, VFar: 0},
			dpi:      96,
			expected: Margins{HNear: 96, VNear: 48, HFar: 192, VFar: 0},
		},
		{
			name:     "144 DPI (Scale 1.5)",
			input:    Margins{HNear: 96, VNear: 48, HFar: 192, VFar: 10},
			dpi:      144,
			expected: Margins{HNear: 144, VNear: 72, HFar: 288, VFar: 15},
		},
		{
			name:     "192 DPI (Scale 2.0)",
			input:    Margins{HNear: 10, VNear: 20, HFar: 30, VFar: 40},
			dpi:      192,
			expected: Margins{HNear: 20, VNear: 40, HFar: 60, VFar: 80},
		},
		{
			name:     "120 DPI (Scale 1.25) - Rounding",
			input:    Margins{HNear: 10, VNear: 10, HFar: 10, VFar: 10}, // 10 * 1.25 = 12.5 -> rounds to 13
			dpi:      120,
			expected: Margins{HNear: 13, VNear: 13, HFar: 13, VFar: 13},
		},
		{
			name:     "0 DPI (Scale 0.0)",
			input:    Margins{HNear: 96, VNear: 48, HFar: 192, VFar: 10},
			dpi:      0,
			expected: Margins{HNear: 0, VNear: 0, HFar: 0, VFar: 0},
		},
		{
			name:     "Negative margins",
			input:    Margins{HNear: -10, VNear: -20, HFar: -30, VFar: -40},
			dpi:      144,
			expected: Margins{HNear: -15, VNear: -30, HFar: -45, VFar: -60},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarginsFrom96DPI(tt.input, tt.dpi)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MarginsFrom96DPI(%v, %d) = %v, want %v", tt.input, tt.dpi, got, tt.expected)
			}
		})
	}
}
