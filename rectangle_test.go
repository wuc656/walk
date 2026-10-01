// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
	"github.com/wuc656/win"
)

func TestRectangleFromRECT(t *testing.T) {
	tests := []struct {
		name     string
		input    win.RECT
		expected Rectangle
	}{
		{
			name: "Zero rect",
			input: win.RECT{Left: 0, Top: 0, Right: 0, Bottom: 0},
			expected: Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
		},
		{
			name: "Positive coords",
			input: win.RECT{Left: 10, Top: 20, Right: 50, Bottom: 80},
			expected: Rectangle{X: 10, Y: 20, Width: 40, Height: 60},
		},
		{
			name: "Negative coords",
			input: win.RECT{Left: -50, Top: -80, Right: -10, Bottom: -20},
			expected: Rectangle{X: -50, Y: -80, Width: 40, Height: 60},
		},
		{
			name: "Mixed coords",
			input: win.RECT{Left: -10, Top: -10, Right: 10, Bottom: 10},
			expected: Rectangle{X: -10, Y: -10, Width: 20, Height: 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RectangleFromRECT(tt.input)
			if result != tt.expected {
				t.Errorf("RectangleFromRECT(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
