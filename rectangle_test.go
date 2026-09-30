// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"reflect"
	"testing"
)

func TestRectangle_Size(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want Size
	}{
		{
			name: "zero size",
			r:    Rectangle{X: 0, Y: 0, Width: 0, Height: 0},
			want: Size{Width: 0, Height: 0},
		},
		{
			name: "positive dimensions",
			r:    Rectangle{X: 10, Y: 20, Width: 100, Height: 200},
			want: Size{Width: 100, Height: 200},
		},
		{
			name: "negative dimensions",
			r:    Rectangle{X: -5, Y: -5, Width: -50, Height: -60},
			want: Size{Width: -50, Height: -60},
		},
		{
			name: "mixed dimensions",
			r:    Rectangle{X: 10, Y: -20, Width: 150, Height: 0},
			want: Size{Width: 150, Height: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Size(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rectangle.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
