// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package walk

import "sync"

type Condition interface {
	Expression
	Satisfied() bool
}

type MutableCondition struct {
	mu               sync.RWMutex
	satisfied        bool
	changedPublisher EventPublisher
}

func NewMutableCondition() *MutableCondition {
	return new(MutableCondition)
}

func (mc *MutableCondition) Value() any {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	return mc.satisfied
}

func (mc *MutableCondition) Satisfied() bool {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	return mc.satisfied
}

func (mc *MutableCondition) SetSatisfied(satisfied bool) error {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	if satisfied == mc.satisfied {
		return nil
	}

	mc.satisfied = satisfied

	mc.changedPublisher.Publish()

	return nil
}

func (mc *MutableCondition) Changed() *Event {
	return mc.changedPublisher.Event()
}

type DelegateCondition struct {
	satisfied func() bool
	changed   *Event
}

func NewDelegateCondition(satisfied func() bool, changed *Event) *DelegateCondition {
	if changed == nil {
		panic("DelegateCondition: changed event cannot be nil")
	}
	return &DelegateCondition{satisfied, changed}
}

func (dc *DelegateCondition) Value() any {
	return dc.satisfied()
}

func (dc *DelegateCondition) Satisfied() bool {
	return dc.satisfied()
}

func (dc *DelegateCondition) Changed() *Event {
	return dc.changed
}

type compositeCondition struct {
	items               []Condition
	itemsChangedHandles []int
	changedPublisher    EventPublisher
	disposed            bool
}

func (cc *compositeCondition) init(items []Condition) {
	cc.items = append(cc.items, items...)

	for _, item := range items {
		handle := item.Changed().Attach(func() {
			cc.changedPublisher.Publish()
		})
		cc.itemsChangedHandles = append(cc.itemsChangedHandles, handle)
	}
}

func (cc *compositeCondition) satisfied(all bool) bool {
	for _, item := range cc.items {
		if all != item.Satisfied() {
			return !all
		}
	}

	return all
}

func (cc *compositeCondition) Changed() *Event {
	return cc.changedPublisher.Event()
}

func (cc *compositeCondition) Dispose() {
	if cc.disposed {
		return
	}
	cc.disposed = true
	for i, item := range cc.items {
		item.Changed().Detach(cc.itemsChangedHandles[i])
	}
	cc.items = nil
	cc.itemsChangedHandles = nil
}

type allCondition struct {
	compositeCondition
}

func NewAllCondition(items ...Condition) Condition {
	ac := new(allCondition)

	ac.init(items)

	return ac
}

func (ac *allCondition) Value() any {
	return ac.Satisfied()
}

func (ac *allCondition) Satisfied() bool {
	return ac.satisfied(true)
}

type anyCondition struct {
	compositeCondition
}

func NewAnyCondition(items ...Condition) Condition {
	ac := new(anyCondition)

	ac.init(items)

	return ac
}

func (ac *anyCondition) Value() any {
	return ac.Satisfied()
}

func (ac *anyCondition) Satisfied() bool {
	return ac.satisfied(false)
}

type negatedCondition struct {
	other Condition
}

func NewNegatedCondition(other Condition) Condition {
	return &negatedCondition{other}
}

func (nc *negatedCondition) Value() any {
	return nc.Satisfied()
}

func (nc *negatedCondition) Satisfied() bool {
	return !nc.other.Satisfied()
}

func (nc *negatedCondition) Changed() *Event {
	return nc.other.Changed()
}
