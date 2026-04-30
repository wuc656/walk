// Copyright 2026 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"reflect"
	"syscall"
	"testing"
)

func TestAppendStringToUTF16MatchesSyscall(t *testing.T) {
	for _, s := range []string{
		"",
		"abc",
		"中文",
		"emoji \U0001f600",
	} {
		got := appendStringToUTF16(nil, s)
		want := syscall.StringToUTF16(s)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("appendStringToUTF16(%q) = %#v, want %#v", s, got, want)
		}
	}
}
