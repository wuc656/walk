// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package walk

import (
	"errors"
	"fmt"
)

var (
	ErrPropertyReadOnly       = errors.New("read-only property")
	ErrPropertyNotValidatable = errors.New("property not validatable")
)

type Property interface {
	Expression
	ReadOnly() bool
	Get() any
	Set(value any) error
	Source() any
	SetSource(source any) error
	Validatable() bool
	Validator() Validator
	SetValidator(validator Validator) error
}

type property struct {
	get                 func() any
	set                 func(v any) error
	changed             *Event
	source              any
	sourceChangedHandle int
	validator           Validator
}

func NewProperty(get func() any, set func(v any) error, changed *Event) Property {
	return &property{get: get, set: set, changed: changed}
}

func (p *property) ReadOnly() bool {
	return p.set == nil
}

func (p *property) Value() any {
	return p.get()
}

func (p *property) Get() any {
	return p.get()
}

func (p *property) Set(value any) error {
	if p.ReadOnly() {
		return ErrPropertyReadOnly
	}

	if oldValue := p.get(); value == oldValue {
		return nil
	}

	return p.set(value)
}

func (p *property) Changed() *Event {
	return p.changed
}

func (p *property) Source() any {
	return p.source
}

func (p *property) SetSource(source any) error {
	if p.ReadOnly() {
		return ErrPropertyReadOnly
	}

	if source != nil {
		switch source := source.(type) {
		case string:
			// nop

		case Property:
			if err := checkPropertySource(p, source); err != nil {
				return err
			}

			if source != nil {
				p.Set(source.Get())

				p.sourceChangedHandle = source.Changed().Attach(func() {
					p.Set(source.Get())
				})
			}

		case Expression:
			p.Set(source.Value())

			p.sourceChangedHandle = source.Changed().Attach(func() {
				p.Set(source.Value())
			})

		default:
			return newError("invalid source type")
		}
	}

	if oldProp, ok := p.source.(Property); ok {
		oldProp.Changed().Detach(p.sourceChangedHandle)
	}

	p.source = source

	return nil
}

func (p *property) Validatable() bool {
	return true
}

func (p *property) Validator() Validator {
	return p.validator
}

func (p *property) SetValidator(validator Validator) error {
	if p.ReadOnly() {
		return ErrPropertyReadOnly
	}

	p.validator = validator

	return nil
}

type readOnlyProperty struct {
	get     func() any
	changed *Event
}

func NewReadOnlyProperty(get func() any, changed *Event) Property {
	return &readOnlyProperty{get: get, changed: changed}
}

func (*readOnlyProperty) ReadOnly() bool {
	return true
}

func (rop *readOnlyProperty) Value() any {
	return rop.get()
}

func (rop *readOnlyProperty) Get() any {
	return rop.get()
}

func (*readOnlyProperty) Set(value any) error {
	return ErrPropertyReadOnly
}

func (rop *readOnlyProperty) Changed() *Event {
	return rop.changed
}

func (*readOnlyProperty) Source() any {
	return nil
}

func (*readOnlyProperty) SetSource(source any) error {
	return ErrPropertyReadOnly
}

func (*readOnlyProperty) Validatable() bool {
	return false
}

func (*readOnlyProperty) Validator() Validator {
	return nil
}

func (*readOnlyProperty) SetValidator(validator Validator) error {
	return ErrPropertyReadOnly
}

type boolProperty struct {
	get                 func() bool
	set                 func(v bool) error
	changed             *Event
	source              any
	sourceChangedHandle int
}

func NewBoolProperty(get func() bool, set func(b bool) error, changed *Event) Property {
	return &boolProperty{get: get, set: set, changed: changed}
}

func (bp *boolProperty) ReadOnly() bool {
	return bp.set == nil
}

func (bp *boolProperty) Value() any {
	return bp.get()
}

func (bp *boolProperty) Get() any {
	return bp.get()
}

func (bp *boolProperty) Set(value any) error {
	if bp.ReadOnly() {
		return ErrPropertyReadOnly
	}

	/* FIXME: Visible property doesn't like this.
	if oldValue := bp.get(); value == oldValue {
		return nil
	}*/

	return bp.set(value.(bool))
}

func (bp *boolProperty) Changed() *Event {
	return bp.changed
}

func (bp *boolProperty) Source() any {
	return bp.source
}

func (bp *boolProperty) SetSource(source any) error {
	if bp.ReadOnly() {
		return ErrPropertyReadOnly
	}

	if source != nil {
		switch source := source.(type) {
		case string:
			// nop

		case Condition:
			if err := checkPropertySource(bp, source); err != nil {
				return err
			}

			if err := bp.Set(source.Satisfied()); err != nil {
				return err
			}

			bp.sourceChangedHandle = source.Changed().Attach(func() {
				bp.Set(source.Satisfied())
			})

		case Expression:
			if err := checkPropertySource(bp, source); err != nil {
				return err
			}

			if satisfied, ok := source.Value().(bool); ok {
				if err := bp.Set(satisfied); err != nil {
					return err
				}
			}

			bp.sourceChangedHandle = source.Changed().Attach(func() {
				if satisfied, ok := source.Value().(bool); ok {
					bp.Set(satisfied)
				}
			})

		default:
			return newError(fmt.Sprintf(`invalid source: "%s" of type %T`, source, source))
		}
	}

	if oldCond, ok := bp.source.(Condition); ok {
		oldCond.Changed().Detach(bp.sourceChangedHandle)
	}

	bp.source = source

	return nil
}

func (bp *boolProperty) Validatable() bool {
	return false
}

func (*boolProperty) Validator() Validator {
	return nil
}

func (*boolProperty) SetValidator(validator Validator) error {
	return ErrPropertyNotValidatable
}

func (bp *boolProperty) Satisfied() bool {
	return bp.get()
}

type readOnlyBoolProperty struct {
	get     func() bool
	changed *Event
}

func NewReadOnlyBoolProperty(get func() bool, changed *Event) Property {
	return &readOnlyBoolProperty{get: get, changed: changed}
}

func (*readOnlyBoolProperty) ReadOnly() bool {
	return true
}

func (robp *readOnlyBoolProperty) Value() any {
	return robp.get()
}

func (robp *readOnlyBoolProperty) Get() any {
	return robp.get()
}

func (*readOnlyBoolProperty) Set(value any) error {
	return ErrPropertyReadOnly
}

func (robp *readOnlyBoolProperty) Changed() *Event {
	return robp.changed
}

func (*readOnlyBoolProperty) Source() any {
	return nil
}

func (*readOnlyBoolProperty) SetSource(source any) error {
	return ErrPropertyReadOnly
}

func (*readOnlyBoolProperty) Validatable() bool {
	return false
}

func (*readOnlyBoolProperty) Validator() Validator {
	return nil
}

func (*readOnlyBoolProperty) SetValidator(validator Validator) error {
	return ErrPropertyNotValidatable
}

func (robp *readOnlyBoolProperty) Satisfied() bool {
	return robp.get()
}

func checkPropertySource(prop Property, source any) error {
	switch source := source.(type) {
	case Property:
		for cur := source; cur != nil; cur, _ = cur.Source().(Property) {
			if cur == prop {
				return newError("source cycle")
			}
		}
	}

	return nil
}
