/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import (
	"strconv"
	"strings"
)

type tString struct {
	listeners []StringListener
	value     string
	filter    StringFilter
}

type tStringAB struct {
	tString
	parentA String
	parentB String
}

type tStringAppend struct {
	tStringAB
}

type tStringBooleanAB struct {
	tBoolean
	parentA String
	parentB String
}

type tStringEqual struct {
	tStringBooleanAB
}

type tStringGreater struct {
	tStringBooleanAB
}

type tStringGreaterOrEqual struct {
	tStringBooleanAB
}

type tStringLess struct {
	tStringBooleanAB
}

type tStringLessOrEqual struct {
	tStringBooleanAB
}

type tStringLowerCase struct {
	tString
}

type tStringNotEqual struct {
	tStringBooleanAB
}

func (stringValue *tString) AddListener(listener StringListener) {
	if !containsStringListener(stringValue.listeners, listener) {
		stringValue.listeners = append(stringValue.listeners, listener)
	}
}

func (stringValue *tString) Append(stringValueB String) String {
	stringAppend := new(tStringAppend)
	stringAppend.parentA = stringValue
	stringAppend.parentB = stringValueB
	stringValue.AddListener(stringAppend)
	stringValueB.AddListener(stringAppend)
	return stringAppend
}

func (stringValue *tString) Boolean() Boolean {
	booleanValue := new(tBoolean)
	stringValue.AddListener(booleanValue)
	return booleanValue
}

func (stringValue *tString) EqualTo(stringValueB String) Boolean {
	stringEqual := new(tStringEqual)
	stringEqual.parentA = stringValue
	stringEqual.parentB = stringValueB
	stringValue.AddListener(stringEqual)
	stringValueB.AddListener(stringEqual)
	return stringEqual
}

func (stringValue *tString) Float64() Float64 {
	float64Value := new(tFloat64)
	stringValue.AddListener(float64Value)
	return float64Value
}

func (stringValue *tString) Float32() Float32 {
	float32Value := new(tFloat32)
	stringValue.AddListener(float32Value)
	return float32Value
}

func (stringValue *tString) GreaterThan(stringValueB String) Boolean {
	stringGreater := new(tStringGreater)
	stringGreater.parentA = stringValue
	stringGreater.parentB = stringValueB
	stringValue.AddListener(stringGreater)
	stringValueB.AddListener(stringGreater)
	return stringGreater
}

func (stringValue *tString) GreaterThanOrEqualTo(stringValueB String) Boolean {
	stringGreaterOrEqual := new(tStringGreaterOrEqual)
	stringGreaterOrEqual.parentA = stringValue
	stringGreaterOrEqual.parentB = stringValueB
	stringValue.AddListener(stringGreaterOrEqual)
	stringValueB.AddListener(stringGreaterOrEqual)
	return stringGreaterOrEqual
}

func (stringValue *tString) Int() Int {
	intValue := new(tInt)
	stringValue.AddListener(intValue)
	return intValue
}

func (stringValue *tString) LessThan(stringValueB String) Boolean {
	stringLess := new(tStringLess)
	stringLess.parentA = stringValue
	stringLess.parentB = stringValueB
	stringValue.AddListener(stringLess)
	stringValueB.AddListener(stringLess)
	return stringLess
}

func (stringValue *tString) LessThanOrEqualTo(stringValueB String) Boolean {
	stringLessOrEqual := new(tStringLessOrEqual)
	stringLessOrEqual.parentA = stringValue
	stringLessOrEqual.parentB = stringValueB
	stringValue.AddListener(stringLessOrEqual)
	stringValueB.AddListener(stringLessOrEqual)
	return stringLessOrEqual
}

func (stringValue *tString) LowerCase() String {
	stringLowerCase := new(tStringLowerCase)
	stringValue.AddListener(stringLowerCase)
	return stringLowerCase
}

func (stringValue *tString) NotEqualTo(stringValueB String) Boolean {
	stringValueNotEqual := new(tStringNotEqual)
	stringValueNotEqual.parentA = stringValue
	stringValueNotEqual.parentB = stringValueB
	stringValue.AddListener(stringValueNotEqual)
	stringValueB.AddListener(stringValueNotEqual)
	return stringValueNotEqual
}

func (stringValue *tString) RemoveListener(listener StringListener) {
	i := indexStringListener(stringValue.listeners, listener)
	if i >= 0 {
		stringValue.listeners = removeStringListener(stringValue.listeners, i)
	}
}

func (stringValue *tString) Set(newValue string) {
	oldValue := stringValue.value
	if stringValue.filter != nil {
		newValue = stringValue.filter.FilterString(stringValue, oldValue, newValue)
	}
	if stringValue.value != newValue {
		observable := String(stringValue)
		stringValue.value = newValue
		for _, listener := range stringValue.listeners {
			listener.StringChanged(observable, oldValue, newValue)
		}
	}
}

func (stringValue *tString) SetFilter(filter StringFilter) {
	stringValue.filter = filter
}

func (stringValue *tString) Value() string {
	return stringValue.value
}

func (stringValue *tString) Float64Changed(observable Float64, oldValue, newValue float64) {
	stringValue.Set(strconv.FormatFloat(newValue, 'f', -1, 64))
}

func (stringValue *tString) Float32Changed(observable Float32, oldValue, newValue float32) {
	stringValue.Set(strconv.FormatFloat(float64(newValue), 'f', -1, 32))
}

func (stringValue *tString) IntChanged(observable Int, oldValue, newValue int) {
	stringValue.Set(strconv.Itoa(newValue))
}

func (stringValue *tString) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if newValue {
		stringValue.Set("true")
	} else {
		stringValue.Set("false")
	}
}

func (stringValue *tStringAppend) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(newValue + stringValue.parentB.Value())
	} else {
		stringValue.Set(stringValue.parentA.Value() + newValue)
	}
}

func (stringValue *tStringEqual) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() == newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() == newValue)
	}
}

func (stringValue *tStringGreater) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() < newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() > newValue)
	}
}

func (stringValue *tStringGreaterOrEqual) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() <= newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() >= newValue)
	}
}

func (stringValue *tStringLess) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() > newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() < newValue)
	}
}

func (stringValue *tStringLessOrEqual) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() >= newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() <= newValue)
	}
}

func (stringValue *tStringLowerCase) StringChanged(observable String, oldValue, newValue string) {
	stringValue.Set(strings.ToLower(newValue))
}

func (stringValue *tStringNotEqual) StringChanged(observable String, oldValue, newValue string) {
	if stringValue.parentA == observable {
		stringValue.Set(stringValue.parentB.Value() != newValue)
	} else {
		stringValue.Set(stringValue.parentA.Value() != newValue)
	}
}
