// Copyright 2023 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_Right(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "Happy path",
			rect: Rectangle{X: 10, Width: 20},
			want: 29, // 10 + 20 - 1
		},
		{
			name: "Zero dimensions",
			rect: Rectangle{X: 0, Width: 0},
			want: -1, // 0 + 0 - 1
		},
		{
			name: "Negative coordinates",
			rect: Rectangle{X: -10, Width: 5},
			want: -6, // -10 + 5 - 1
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

func TestRectangle_Bottom(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "Happy path",
			rect: Rectangle{Y: 10, Height: 20},
			want: 29, // 10 + 20 - 1
		},
		{
			name: "Zero dimensions",
			rect: Rectangle{Y: 0, Height: 0},
			want: -1, // 0 + 0 - 1
		},
		{
			name: "Negative coordinates",
			rect: Rectangle{Y: -10, Height: 5},
			want: -6, // -10 + 5 - 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Bottom(); got != tt.want {
				t.Errorf("Rectangle.Bottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
