// Copyright 2025 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_IsZero(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want bool
	}{
		{
			name: "all zero",
			r:    Rectangle{0, 0, 0, 0},
			want: true,
		},
		{
			name: "x not zero",
			r:    Rectangle{1, 0, 0, 0},
			want: false,
		},
		{
			name: "y not zero",
			r:    Rectangle{0, 1, 0, 0},
			want: false,
		},
		{
			name: "width not zero",
			r:    Rectangle{0, 0, 1, 0},
			want: false,
		},
		{
			name: "height not zero",
			r:    Rectangle{0, 0, 0, 1},
			want: false,
		},
		{
			name: "all positive",
			r:    Rectangle{1, 2, 3, 4},
			want: false,
		},
		{
			name: "negative x",
			r:    Rectangle{-1, 0, 0, 0},
			want: false,
		},
		{
			name: "negative fields",
			r:    Rectangle{-1, -2, -3, -4},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsZero(); got != tt.want {
				t.Errorf("Rectangle.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
