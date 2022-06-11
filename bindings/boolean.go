/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

type tBoolean struct {
	listeners []BooleanListener
	value     bool
	filter    BooleanFilter
}

type tBooleanAB struct {
	tBoolean
	parentA Boolean
	parentB Boolean
}

type tBooleanAnd struct {
	tBooleanAB
}

type tBooleanEqual struct {
	tBooleanAB
}

type tBooleanNot struct {
	tBoolean
}

type tBooleanNotEqual struct {
	tBooleanAB
}

type tBooleanOr struct {
	tBooleanAB
}

func (booleanValue *tBoolean) AddListener(listener BooleanListener) {
	if !containsBooleanListener(booleanValue.listeners, listener) {
		booleanValue.listeners = append(booleanValue.listeners, listener)
	}
}

func (booleanValue *tBoolean) And(booleanValueB Boolean) Boolean {
	booleanValueAnd := new(tBooleanAnd)
	booleanValueAnd.parentA = booleanValue
	booleanValueAnd.parentB = booleanValueB
	booleanValue.AddListener(booleanValueAnd)
	booleanValueB.AddListener(booleanValueAnd)
	return booleanValueAnd
}

func (booleanValue *tBoolean) EqualTo(booleanValueB Boolean) Boolean {
	booleanValueEqual := new(tBooleanEqual)
	booleanValueEqual.parentA = booleanValue
	booleanValueEqual.parentB = booleanValueB
	booleanValue.AddListener(booleanValueEqual)
	booleanValueB.AddListener(booleanValueEqual)
	return booleanValueEqual
}

func (booleanValue *tBoolean) Int() Int {
	intValue := new(tInt)
	booleanValue.AddListener(intValue)
	return intValue
}

func (booleanValue *tBoolean) Not() Boolean {
	booleanValueNot := new(tBooleanNot)
	booleanValue.AddListener(booleanValueNot)
	return booleanValueNot
}

func (booleanValue *tBoolean) NotEqualTo(booleanValueB Boolean) Boolean {
	booleanValueNotEqual := new(tBooleanNotEqual)
	booleanValueNotEqual.parentA = booleanValue
	booleanValueNotEqual.parentB = booleanValueB
	booleanValue.AddListener(booleanValueNotEqual)
	booleanValueB.AddListener(booleanValueNotEqual)
	return booleanValueNotEqual
}

func (booleanValue *tBoolean) Or(booleanValueB Boolean) Boolean {
	booleanValueOr := new(tBooleanOr)
	booleanValueOr.parentA = booleanValue
	booleanValueOr.parentB = booleanValueB
	booleanValue.AddListener(booleanValueOr)
	booleanValueB.AddListener(booleanValueOr)
	return booleanValueOr
}

func (booleanValue *tBoolean) RemoveListener(listener BooleanListener) {
	i := indexBooleanListener(booleanValue.listeners, listener)
	if i >= 0 {
		booleanValue.listeners = removeBooleanListener(booleanValue.listeners, i)
	}
}

func (booleanValue *tBoolean) Set(newValue bool) {
	oldValue := booleanValue.value
	if booleanValue.filter != nil {
		newValue = booleanValue.filter.FilterBoolean(booleanValue, oldValue, newValue)
	}
	if oldValue != newValue {
		observable := Boolean(booleanValue)
		booleanValue.value = newValue
		for _, listener := range booleanValue.listeners {
			listener.BooleanChanged(observable, oldValue, newValue)
		}
	}
}

func (booleanValue *tBoolean) SetFilter(filter BooleanFilter) {
	booleanValue.filter = filter
}

func (booleanValue *tBoolean) String() String {
	stringValue := new(tString)
	booleanValue.AddListener(stringValue)
	return stringValue
}

func (booleanValue *tBoolean) Value() bool {
	return booleanValue.value
}

func (booleanValue *tBoolean) StringChanged(observable String, oldValue, newValue string) {
	if newValue == "true" || newValue == "1" || newValue == "yes" {
		booleanValue.Set(true)
	} else {
		booleanValue.Set(false)
	}
}

func (booleanValue *tBooleanAnd) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.Set(booleanValue.parentB.Value() && newValue)
	} else {
		booleanValue.Set(booleanValue.parentA.Value() && newValue)
	}
}

func (booleanValue *tBooleanEqual) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.Set(booleanValue.parentB.Value() == newValue)
	} else {
		booleanValue.Set(booleanValue.parentA.Value() == newValue)
	}
}

func (booleanValue *tBooleanNot) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	booleanValue.Set(!newValue)
}

func (booleanValue *tBooleanNotEqual) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.Set(booleanValue.parentB.Value() != newValue)
	} else {
		booleanValue.Set(booleanValue.parentA.Value() != newValue)
	}
}

func (booleanValue *tBooleanOr) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if booleanValue.parentA == observable {
		booleanValue.Set(booleanValue.parentB.Value() || newValue)
	} else {
		booleanValue.Set(booleanValue.parentA.Value() || newValue)
	}
}
