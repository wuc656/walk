// Copyright 2023 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import "testing"

func TestSize_IsZero(t *testing.T) {
	tests := []struct {
		name string
		size Size
		want bool
	}{
		{"zero", Size{0, 0}, true},
		{"width non-zero", Size{1, 0}, false},
		{"height non-zero", Size{0, 1}, false},
		{"both non-zero", Size{1, 1}, false},
		{"negative values", Size{-1, -1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.size.IsZero(); got != tt.want {
				t.Errorf("Size.IsZero() = %v, want %v for %v", got, tt.want, tt.size)
			}
		})
	}
}
