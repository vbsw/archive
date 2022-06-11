/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "testing"

type testFloat64Listener struct {
	called   bool
	oldValue float64
	newValue float64
}

func (float64Listener *testFloat64Listener) Float64Changed(observable Float64, oldValue, newValue float64) {
	float64Listener.called = true
	float64Listener.oldValue = oldValue
	float64Listener.newValue = newValue
}

func TestFloat64ValueAddListener(t *testing.T) {
	float64Value := new(tFloat64)
	float64Listener := new(testFloat64Listener)

	float64Value.AddListener(float64Listener)
	if len(float64Value.listeners) != 1 {
		t.Error(len(float64Value.listeners))
	}

	float64Value.AddListener(float64Listener)
	if len(float64Value.listeners) != 1 {
		t.Error(len(float64Value.listeners))
	}
}

func TestFloat64ValueSet(t *testing.T) {
	float64Value := new(tFloat64)
	float64Value.value = 0

	if float64Value.Value() != 0 {
		t.Error(float64Value.Value())
	}

	float64Value.Set(10)
	if float64Value.Value() != 10 {
		t.Error(float64Value.Value())
	}
}

func TestFloat64ValueFloat64Changed(t *testing.T) {
	float64Value := new(tFloat64)
	float64Listener := new(testFloat64Listener)
	float64Value.value = 0

	float64Value.AddListener(float64Listener)
	float64Value.Set(0)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64Value.Set(10)
	if float64Listener.called != true {
		t.Error(float64Listener.called)
	}
	if float64Listener.oldValue != 0 {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 10 {
		t.Error(float64Listener.newValue)
	}

	float64Value.Set(50)
	if float64Listener.oldValue != 10 {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 50 {
		t.Error(float64Listener.newValue)
	}
}

func TestFloat64ValueDivide(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	float64ValueA.value = 1
	float64ValueB.value = 1

	float64ValueDivide := float64ValueA.Divide(float64ValueB)
	float64Listener := new(testFloat64Listener)

	float64ValueDivide.AddListener(float64Listener)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64ValueA.Set(10)
	if float64ValueDivide.Value() != 10 {
		t.Error(float64ValueDivide.Value())
	}
	if float64Listener.called != true {
		t.Error(float64Listener.called)
	}
	if float64Listener.newValue != 10 {
		t.Error(float64Listener.newValue)
	}

	float64ValueB.Set(10)
	if float64ValueDivide.Value() != 1 {
		t.Error(float64ValueDivide.Value())
	}
	if float64Listener.oldValue != 10 {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 1 {
		t.Error(float64Listener.newValue)
	}
}

func TestFloat64ValueEqualTo(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	booleanValueEqual := float64ValueA.EqualTo(float64ValueB)
	booleanListener := new(testBooleanListener)
	float64ValueA.value = 0
	float64ValueB.value = 0

	booleanValueEqual.AddListener(booleanListener)
	float64ValueA.Set(0)
	float64ValueB.Set(0)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	float64ValueA.Set(10)
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

	float64ValueB.Set(10)
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

func TestFloat64ValueGreaterThan(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	booleanValueGreater := float64ValueA.GreaterThan(float64ValueB)
	booleanListener := new(testBooleanListener)
	float64ValueA.value = 0
	float64ValueB.value = 0

	booleanValueGreater.AddListener(booleanListener)
	float64ValueA.Set(0)
	float64ValueB.Set(0)
	if booleanValueGreater.Value() != false {
		t.Error(booleanValueGreater.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	float64ValueA.Set(10)
	if booleanValueGreater.Value() != true {
		t.Error(booleanValueGreater.Value())
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

	float64ValueB.Set(10)
	if booleanValueGreater.Value() != false {
		t.Error(booleanValueGreater.Value())
	}
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}

func TestFloat64ValueGreaterThanOrEqualTo(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	booleanValueGreaterOrEqual := float64ValueA.GreaterThanOrEqualTo(float64ValueB)
	booleanListener := new(testBooleanListener)

	booleanValueGreaterOrEqual.AddListener(booleanListener)
	float64ValueA.Set(10)
	if booleanValueGreaterOrEqual.Value() != true {
		t.Error(booleanValueGreaterOrEqual.Value())
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

	float64ValueB.Set(10)
	if booleanValueGreaterOrEqual.Value() != true {
		t.Error(booleanValueGreaterOrEqual.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	float64ValueB.Set(50)
	if booleanValueGreaterOrEqual.Value() != false {
		t.Error(booleanValueGreaterOrEqual.Value())
	}
	if booleanListener.oldValue != true {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != false {
		t.Error(booleanListener.newValue)
	}
}

func TestFloat64ValueMinus(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	float64ValueMinus := float64ValueA.Minus(float64ValueB)
	float64Listener := new(testFloat64Listener)

	float64ValueMinus.AddListener(float64Listener)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64ValueA.Set(10)
	if float64ValueMinus.Value() != 10 {
		t.Error(float64ValueMinus.Value())
	}
	if float64Listener.called != true {
		t.Error(float64Listener.called)
	}
	if float64Listener.newValue != 10 {
		t.Error(float64Listener.newValue)
	}

	float64ValueB.Set(4)
	if float64ValueMinus.Value() != 6 {
		t.Error(float64ValueMinus.Value())
	}
	if float64Listener.oldValue != 10 {
		t.Error(float64Listener.oldValue)
	}
	if float64Listener.newValue != 6 {
		t.Error(float64Listener.newValue)
	}
}

func TestFloat64ValueMultiply(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	float64ValueA.value = 0
	float64ValueB.value = 0

	float64ValueMultiply := float64ValueA.Multiply(float64ValueB)
	float64Listener := new(testFloat64Listener)

	float64ValueMultiply.AddListener(float64Listener)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64ValueA.Set(10)
	if float64Listener.called != false {
		t.Error(float64Listener.called)
	}

	float64ValueB.Set(2)
	if float64Listener.called != true {
		t.Error(float64Listener.called)
	}
	if float64ValueMultiply.Value() != 20 {
		t.Error(float64ValueMultiply.Value())
	}
	if float64Listener.newValue != 20 {
		t.Error(float64Listener.newValue)
	}
}

func TestFloat64ValueNotEqualTo(t *testing.T) {
	float64ValueA := new(tFloat64)
	float64ValueB := new(tFloat64)
	float64ValueA.value = 0
	float64ValueB.value = 0
	booleanValueNotEqual := float64ValueA.NotEqualTo(float64ValueB)
	booleanListener := new(testBooleanListener)

	booleanValueNotEqual.AddListener(booleanListener)
	float64ValueA.Set(0)
	float64ValueB.Set(0)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	float64ValueA.Set(10)
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

	float64ValueB.Set(10)
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
