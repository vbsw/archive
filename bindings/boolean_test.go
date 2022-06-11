/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "testing"

type testBooleanListener struct {
	called   bool
	oldValue bool
	newValue bool
}

func (booleanListener *testBooleanListener) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	booleanListener.called = true
	booleanListener.oldValue = oldValue
	booleanListener.newValue = newValue
}

func TestBooleanValueAddListener(t *testing.T) {
	booleanValue := new(tBoolean)
	booleanListener := new(testBooleanListener)

	booleanValue.AddListener(booleanListener)
	if len(booleanValue.listeners) != 1 {
		t.Error(len(booleanValue.listeners))
	}

	booleanValue.AddListener(booleanListener)
	if len(booleanValue.listeners) != 1 {
		t.Error(len(booleanValue.listeners))
	}
}

func TestBooleanValueSet(t *testing.T) {
	booleanValue := new(tBoolean)

	if booleanValue.Value() != false {
		t.Error(booleanValue.Value())
	}

	booleanValue.Set(true)
	if booleanValue.Value() != true {
		t.Error(booleanValue.Value())
	}
}

func TestBooleanValueBooleanChanged(t *testing.T) {
	booleanValue := new(tBoolean)
	booleanListener := new(testBooleanListener)

	booleanValue.AddListener(booleanListener)
	booleanValue.Set(false)
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValue.Set(true)
	if booleanListener.called != true {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	booleanValue.Set(false)
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}

func TestBooleanValueAnd(t *testing.T) {
	booleanValueA := new(tBoolean)
	booleanValueB := new(tBoolean)
	booleanValueAnd := booleanValueA.And(booleanValueB)
	booleanListener := new(testBooleanListener)

	booleanValueAnd.AddListener(booleanListener)
	booleanValueA.Set(false)
	booleanValueB.Set(false)
	if booleanValueAnd.Value() != false {
		t.Error(booleanValueAnd.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValueA.Set(true)
	if booleanValueAnd.Value() != false {
		t.Error(booleanValueAnd.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}

	booleanValueB.Set(true)
	if booleanListener.called != true {
		t.Error(booleanListener.called)
	}
	if booleanValueAnd.Value() != true {
		t.Error(booleanValueAnd.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}
}

func TestBooleanValueEqualTo(t *testing.T) {
	booleanValueA := new(tBoolean)
	booleanValueB := new(tBoolean)
	booleanValueEqual := booleanValueA.EqualTo(booleanValueB)
	booleanListener := new(testBooleanListener)

	booleanValueEqual.AddListener(booleanListener)
	booleanValueA.Set(false)
	booleanValueB.Set(false)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValueA.Set(true)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}

	booleanValueB.Set(true)
	if booleanValueEqual.Value() != true {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}
}

func TestBooleanValueNot(t *testing.T) {
	booleanValue := new(tBoolean)
	booleanValueNot := booleanValue.Not()
	booleanListener := new(testBooleanListener)

	booleanValueNot.AddListener(booleanListener)
	booleanValue.Set(false)
	if booleanValueNot.Value() != false {
		t.Error(booleanValueNot.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValue.Set(true)
	if booleanValueNot.Value() != false {
		t.Error(booleanValueNot.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}

	booleanValue.Set(false)
	if booleanValueNot.Value() != true {
		t.Error(booleanValueNot.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}
}

func TestBooleanValueNotEqualTo(t *testing.T) {
	booleanValueA := new(tBoolean)
	booleanValueB := new(tBoolean)
	booleanValueNotEqual := booleanValueA.NotEqualTo(booleanValueB)
	booleanListener := new(testBooleanListener)

	booleanValueNotEqual.AddListener(booleanListener)
	booleanValueA.Set(false)
	booleanValueB.Set(false)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValueA.Set(true)
	if booleanValueNotEqual.Value() != true {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != true {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	booleanValueB.Set(true)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}

func TestBooleanValueOr(t *testing.T) {
	booleanValueA := new(tBoolean)
	booleanValueB := new(tBoolean)
	booleanValueOr := booleanValueA.Or(booleanValueB)
	booleanListener := new(testBooleanListener)

	booleanValueOr.AddListener(booleanListener)
	booleanValueA.Set(false)
	booleanValueB.Set(false)
	if booleanValueOr.Value() != false {
		t.Error(booleanValueOr.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	booleanValueA.Set(true)
	if booleanValueOr.Value() != true {
		t.Error(booleanValueOr.Value())
	}
	if booleanListener.called != true {
		t.Error(booleanListener.called)
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	booleanValueB.Set(true)
	if booleanValueOr.Value() != true {
		t.Error(booleanValueOr.Value())
	}
	// booleanListener was not called, old value stays the same
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	booleanValueA.Set(false)
	booleanValueB.Set(false)
	if booleanValueOr.Value() != false {
		t.Error(booleanValueOr.Value())
	}
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}
