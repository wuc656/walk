// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"testing"
)

func TestRectangle_Top(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "zero rectangle",
			rect: Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			want: 0,
		},
		{
			name: "positive coordinates",
			rect: Rectangle{X: 10, Y: 20, Width: 100, Height: 50},
			want: 20,
		},
		{
			name: "negative coordinates",
			rect: Rectangle{X: -10, Y: -20, Width: 100, Height: 50},
			want: -20,
		},
		{
			name: "mixed coordinates",
			rect: Rectangle{X: 10, Y: -20, Width: 100, Height: 50},
			want: -20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Top(); got != tt.want {
				t.Errorf("Rectangle.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}
