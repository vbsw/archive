/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "strconv"

type tFloat64 struct {
	listeners []Float64Listener
	value     float64
	filter    Float64Filter
}

type tFloat64AB struct {
	tFloat64
	parentA Float64
	parentB Float64
}

type tFloat64BooleanAB struct {
	tBoolean
	parentA Float64
	parentB Float64
}

type tFloat64Divide struct {
	tFloat64AB
}

type tFloat64Equal struct {
	tFloat64BooleanAB
}

type tFloat64Greater struct {
	tFloat64BooleanAB
}

type tFloat64GreaterOrEqual struct {
	tFloat64BooleanAB
}

type tFloat64Less struct {
	tFloat64BooleanAB
}

type tFloat64LessOrEqual struct {
	tFloat64BooleanAB
}

type tFloat64Minus struct {
	tFloat64AB
}

type tFloat64Multiply struct {
	tFloat64AB
}

type tFloat64NotEqual struct {
	tFloat64BooleanAB
}

type tFloat64Plus struct {
	tFloat64AB
}

func (float64Value *tFloat64) AddListener(listener Float64Listener) {
	if !containsFloat64Listener(float64Value.listeners, listener) {
		float64Value.listeners = append(float64Value.listeners, listener)
	}
}

func (float64Value *tFloat64) Divide(float64ValueB Float64) Float64 {
	float64Divide := new(tFloat64Divide)
	float64Divide.parentA = float64Value
	float64Divide.parentB = float64ValueB
	float64Value.AddListener(float64Divide)
	float64ValueB.AddListener(float64Divide)
	return float64Divide
}

func (float64Value *tFloat64) EqualTo(float64ValueB Float64) Boolean {
	float64Equal := new(tFloat64Equal)
	float64Equal.parentA = float64Value
	float64Equal.parentB = float64ValueB
	float64Value.AddListener(float64Equal)
	float64ValueB.AddListener(float64Equal)
	return float64Equal
}

func (float64Value *tFloat64) Float32() Float32 {
	float32Value := new(tFloat32)
	float64Value.AddListener(float32Value)
	return float32Value
}

func (float64Value *tFloat64) GreaterThan(float64ValueB Float64) Boolean {
	float64Greater := new(tFloat64Greater)
	float64Greater.parentA = float64Value
	float64Greater.parentB = float64ValueB
	float64Value.AddListener(float64Greater)
	float64ValueB.AddListener(float64Greater)
	return float64Greater
}

func (float64Value *tFloat64) GreaterThanOrEqualTo(float64ValueB Float64) Boolean {
	float64GreaterOrEqual := new(tFloat64GreaterOrEqual)
	float64GreaterOrEqual.parentA = float64Value
	float64GreaterOrEqual.parentB = float64ValueB
	float64Value.AddListener(float64GreaterOrEqual)
	float64ValueB.AddListener(float64GreaterOrEqual)
	return float64GreaterOrEqual
}

func (float64Value *tFloat64) Int() Int {
	intValue := new(tInt)
	float64Value.AddListener(intValue)
	return intValue
}

func (float64Value *tFloat64) LessThan(float64ValueB Float64) Boolean {
	float64Less := new(tFloat64Less)
	float64Less.parentA = float64Value
	float64Less.parentB = float64ValueB
	float64Value.AddListener(float64Less)
	float64ValueB.AddListener(float64Less)
	return float64Less
}

func (float64Value *tFloat64) LessThanOrEqualTo(float64ValueB Float64) Boolean {
	float64LessOrEqual := new(tFloat64LessOrEqual)
	float64LessOrEqual.parentA = float64Value
	float64LessOrEqual.parentB = float64ValueB
	float64Value.AddListener(float64LessOrEqual)
	float64ValueB.AddListener(float64LessOrEqual)
	return float64LessOrEqual
}

func (float64Value *tFloat64) Minus(float64ValueB Float64) Float64 {
	float64Minus := new(tFloat64Minus)
	float64Minus.parentA = float64Value
	float64Minus.parentB = float64ValueB
	float64Value.AddListener(float64Minus)
	float64ValueB.AddListener(float64Minus)
	return float64Minus
}

func (float64Value *tFloat64) Multiply(float64ValueB Float64) Float64 {
	float64Multiply := new(tFloat64Multiply)
	float64Multiply.parentA = float64Value
	float64Multiply.parentB = float64ValueB
	float64Value.AddListener(float64Multiply)
	float64ValueB.AddListener(float64Multiply)
	return float64Multiply
}

func (float64Value *tFloat64) NotEqualTo(float64ValueB Float64) Boolean {
	float64ValueNotEqual := new(tFloat64NotEqual)
	float64ValueNotEqual.parentA = float64Value
	float64ValueNotEqual.parentB = float64ValueB
	float64Value.AddListener(float64ValueNotEqual)
	float64ValueB.AddListener(float64ValueNotEqual)
	return float64ValueNotEqual
}

func (float64Value *tFloat64) Plus(float64ValueB Float64) Float64 {
	float64Plus := new(tFloat64Plus)
	float64Plus.parentA = float64Value
	float64Plus.parentB = float64ValueB
	float64Value.AddListener(float64Plus)
	float64ValueB.AddListener(float64Plus)
	return float64Plus
}

func (float64Value *tFloat64) RemoveListener(listener Float64Listener) {
	i := indexFloat64Listener(float64Value.listeners, listener)
	if i >= 0 {
		float64Value.listeners = removeFloat64Listener(float64Value.listeners, i)
	}
}

func (float64Value *tFloat64) Set(newValue float64) {
	oldValue := float64Value.value
	if float64Value.filter != nil {
		newValue = float64Value.filter.FilterFloat64(float64Value, oldValue, newValue)
	}
	if float64Value.value != newValue {
		observable := Float64(float64Value)
		float64Value.value = newValue
		for _, listener := range float64Value.listeners {
			listener.Float64Changed(observable, oldValue, newValue)
		}
	}
}

func (float64Value *tFloat64) SetFilter(filter Float64Filter) {
	float64Value.filter = filter
}

func (float64Value *tFloat64) String() String {
	stringValue := new(tString)
	float64Value.AddListener(stringValue)
	return stringValue
}

func (float64Value *tFloat64) Value() float64 {
	return float64Value.value
}

func (float64Value *tFloat64) Float32Changed(observable Float32, oldValue, newValue float32) {
	float64Value.Set(float64(newValue))
}

func (float64Value *tFloat64) IntChanged(observable Int, oldValue, newValue int) {
	float64Value.Set(float64(newValue))
}

func (float64Value *tFloat64) StringChanged(observable String, oldValue, newValue string) {
	if val, err := strconv.ParseFloat(newValue, 64); err == nil {
		float64Value.Set(val)
	}
}

func (float64Value *tFloat64Divide) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(newValue / float64Value.parentB.Value())
	} else {
		float64Value.Set(float64Value.parentA.Value() / newValue)
	}
}

func (float64Value *tFloat64Equal) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() == newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() == newValue)
	}
}

func (float64Value *tFloat64Greater) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() < newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() > newValue)
	}
}

func (float64Value *tFloat64GreaterOrEqual) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() <= newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() >= newValue)
	}
}

func (float64Value *tFloat64Less) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() > newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() < newValue)
	}
}

func (float64Value *tFloat64LessOrEqual) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() >= newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() <= newValue)
	}
}

func (float64Value *tFloat64Minus) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(newValue - float64Value.parentB.Value())
	} else {
		float64Value.Set(float64Value.parentA.Value() - newValue)
	}
}

func (float64Value *tFloat64Multiply) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() * newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() * newValue)
	}
}

func (float64Value *tFloat64NotEqual) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() != newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() != newValue)
	}
}

func (float64Value *tFloat64Plus) Float64Changed(observable Float64, oldValue, newValue float64) {
	if float64Value.parentA == observable {
		float64Value.Set(float64Value.parentB.Value() + newValue)
	} else {
		float64Value.Set(float64Value.parentA.Value() + newValue)
	}
}
