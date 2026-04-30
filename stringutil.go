// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import (
	"sync"
	"syscall"
)

// StringToUTF16Cache provides a cached version of StringToUTF16Ptr
// to avoid repeated allocations for frequently used strings.
// This is useful for static strings like window class names and property names.
type StringToUTF16Cache struct {
	mu    sync.RWMutex
	cache map[string]*uint16
}

var defaultCache = &StringToUTF16Cache{
	cache: make(map[string]*uint16),
}

// Get returns a cached UTF16Ptr for the given string.
// The result is valid for the lifetime of the program.
// Ideal for static strings like class names, window themes, etc.
func (c *StringToUTF16Cache) Get(s string) *uint16 {
	c.mu.RLock()
	if ptr, exists := c.cache[s]; exists {
		c.mu.RUnlock()
		return ptr
	}
	c.mu.RUnlock()

	// String not in cache, create it
	ptr, _ := syscall.UTF16PtrFromString(s)

	c.mu.Lock()
	c.cache[s] = ptr
	c.mu.Unlock()

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
