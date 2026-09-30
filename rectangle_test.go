// Copyright 2024 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangleRight(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "Zero width, Zero X",
			rect: Rectangle{X: 0, Width: 0},
			want: -1,
		},
		{
			name: "Positive X, Positive Width",
			rect: Rectangle{X: 10, Width: 5},
			want: 14,
		},
		{
			name: "Negative X, Positive Width",
			rect: Rectangle{X: -5, Width: 10},
			want: 4,
		},
		{
			name: "Positive X, Negative Width",
			rect: Rectangle{X: 5, Width: -2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Right(); got != tt.want {
				t.Errorf("Rectangle.Right() = %v, want %v", got, tt.want)
			}
		})
	}
}
