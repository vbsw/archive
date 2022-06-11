/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "testing"

type testStringListener struct {
	called   bool
	oldValue string
	newValue string
}

func (stringListener *testStringListener) StringChanged(observable String, oldValue, newValue string) {
	stringListener.called = true
	stringListener.oldValue = oldValue
	stringListener.newValue = newValue
}

func TestStringValueAddListener(t *testing.T) {
	stringValue := new(tString)
	stringListener := new(testStringListener)

	stringValue.AddListener(stringListener)
	if len(stringValue.listeners) != 1 {
		t.Error(len(stringValue.listeners))
	}

	stringValue.AddListener(stringListener)
	if len(stringValue.listeners) != 1 {
		t.Error(len(stringValue.listeners))
	}
}

func TestStringValueSet(t *testing.T) {
	stringValue := new(tString)
	testString := "asdf"

	if stringValue.Value() != "" {
		t.Error(stringValue.Value())
	}

	stringValue.Set(testString)
	if stringValue.Value() != testString {
		t.Error(stringValue.Value())
	}
}

func TestStringValueStringChanged(t *testing.T) {
	stringValue := new(tString)
	stringListener := new(testStringListener)
	testStringA := "asdf"
	testStringB := "qwer"

	stringValue.AddListener(stringListener)
	stringValue.Set("")
	if stringListener.called != false {
		t.Error(stringListener.called)
	}

	stringValue.Set(testStringA)
	if stringListener.called != true {
		t.Error(stringListener.called)
	}
	if stringListener.oldValue != "" {
		t.Error(stringListener.oldValue)
	}
	if stringListener.newValue != testStringA {
		t.Error(stringListener.newValue)
	}

	stringValue.Set(testStringB)
	if stringListener.oldValue != testStringA {
		t.Error(stringListener.oldValue)
	}
	if stringListener.newValue != testStringB {
		t.Error(stringListener.newValue)
	}
}

func TestStringValueAppend(t *testing.T) {
	stringValueA := new(tString)
	stringValueB := new(tString)
	stringValueAppend := stringValueA.Append(stringValueB)
	stringListener := new(testStringListener)
	testStringA := "asdf"
	testStringB := "qwer"
	testStringC := testStringA + testStringB

	stringValueAppend.AddListener(stringListener)
	if stringListener.called != false {
		t.Error(stringListener.called)
	}

	stringValueA.Set(testStringA)
	if stringValueAppend.Value() != testStringA {
		t.Error(stringValueAppend.Value())
	}
	if stringListener.called != true {
		t.Error(stringListener.called)
	}
	if stringListener.newValue != testStringA {
		t.Error(stringListener.newValue)
	}

	stringValueB.Set(testStringB)
	if stringValueAppend.Value() != testStringC {
		t.Error(stringValueAppend.Value())
	}
	if stringListener.oldValue != testStringA {
		t.Error(stringListener.oldValue)
	}
	if stringListener.newValue != testStringC {
		t.Error(stringListener.newValue)
	}
}

func TestStringValueEqualTo(t *testing.T) {
	stringValueA := new(tString)
	stringValueB := new(tString)
	booleanValueEqual := stringValueA.EqualTo(stringValueB)
	booleanListener := new(testBooleanListener)
	testString := "asdf"

	booleanValueEqual.AddListener(booleanListener)
	if booleanValueEqual.Value() != false {
		t.Error(booleanValueEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	stringValueA.Set(testString)
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

	stringValueB.Set(testString)
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

func TestStringValueGreaterThan(t *testing.T) {
	stringValueA := new(tString)
	stringValueB := new(tString)
	booleanValueGreater := stringValueA.GreaterThan(stringValueB)
	booleanListener := new(testBooleanListener)
	testString := "asdf"

	booleanValueGreater.AddListener(booleanListener)
	stringValueA.Set("")
	stringValueB.Set("")
	if booleanValueGreater.Value() != false {
		t.Error(booleanValueGreater.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	stringValueA.Set(testString)
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

	stringValueB.Set(testString)
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

func TestStringValueGreaterThanOrEqualTo(t *testing.T) {
	stringValueA := new(tString)
	stringValueB := new(tString)
	booleanValueGreaterOrEqual := stringValueA.GreaterThanOrEqualTo(stringValueB)
	booleanListener := new(testBooleanListener)
	testStringA := "asdf"
	testStringB := "qwer"

	booleanValueGreaterOrEqual.AddListener(booleanListener)
	stringValueA.Set(testStringA)
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

	stringValueB.Set(testStringA)
	if booleanValueGreaterOrEqual.Value() != true {
		t.Error(booleanValueGreaterOrEqual.Value())
	}
	if booleanListener.oldValue != false {
		t.Error(booleanListener.oldValue)
	}
	if booleanListener.newValue != true {
		t.Error(booleanListener.newValue)
	}

	stringValueB.Set(testStringB)
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

func TestStringValueLowerCase(t *testing.T) {
	stringValueA := new(tString)
	stringValueLowerCase := stringValueA.LowerCase()
	stringListener := new(testStringListener)
	testStringA := "ASDF"
	testStringB := "asdf"

	stringValueLowerCase.AddListener(stringListener)
	if stringListener.called != false {
		t.Error(stringListener.called)
	}

	stringValueA.Set(testStringA)
	if stringValueLowerCase.Value() == testStringA {
		t.Error(stringValueLowerCase.Value())
	}
	if stringValueLowerCase.Value() != testStringB {
		t.Error(stringValueLowerCase.Value())
	}
	if stringListener.called != true {
		t.Error(stringListener.called)
	}
	if stringListener.newValue != testStringB {
		t.Error(stringListener.newValue)
	}
}

func TestStringValueNotEqualTo(t *testing.T) {
	stringValueA := new(tString)
	stringValueB := new(tString)
	booleanValueNotEqual := stringValueA.NotEqualTo(stringValueB)
	booleanListener := new(testBooleanListener)
	testString := "asdf"

	booleanValueNotEqual.AddListener(booleanListener)
	if booleanValueNotEqual.Value() != false {
		t.Error(booleanValueNotEqual.Value())
	}
	if booleanListener.called != false {
		t.Error(booleanListener.called)
	}

	stringValueA.Set(testString)
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

	stringValueB.Set(testString)
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
