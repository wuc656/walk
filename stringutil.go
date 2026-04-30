// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"sync"
	"syscall"
	"unicode/utf16"
)

// StringToUTF16Cache provides a cached version of StringToUTF16Ptr
// to avoid repeated allocations for frequently used strings.
// This is useful for static strings like window class names and property names.
type StringToUTF16Cache struct {
	cache sync.Map
}

var defaultCache = &StringToUTF16Cache{}

// Get returns a cached UTF16Ptr for the given string.
// The result is valid for the lifetime of the program.
// Ideal for static strings like class names, window themes, etc.
func (c *StringToUTF16Cache) Get(s string) *uint16 {
	if ptr, exists := c.cache.Load(s); exists {
		return ptr.(*uint16)
	}

	ptr, _ := syscall.UTF16PtrFromString(s)

	if cachedPtr, loaded := c.cache.LoadOrStore(s, ptr); loaded {
		return cachedPtr.(*uint16)
	}

	return ptr
}

// CachedStringToUTF16Ptr is a convenient function using the default cache.
// Use this for static strings like:
// - Window class names
// - Property names (SetWindowTheme, GetThemeAppProperties, etc.)
// - Standard registry paths
//
// Example:
//
//	winapi.SetWindowTheme(hwnd, CachedStringToUTF16Ptr("Explorer"), nil)
func CachedStringToUTF16Ptr(s string) *uint16 {
	return defaultCache.Get(s)
}

func appendStringToUTF16(dst []uint16, s string) []uint16 {
	dst = dst[:0]
	for _, r := range s {
		if r < 0x10000 {
			dst = append(dst, uint16(r))
			continue
		}
		r1, r2 := utf16.EncodeRune(r)
		dst = append(dst, uint16(r1), uint16(r2))
	}
	return append(dst, 0)
}
