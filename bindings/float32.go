/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "strconv"

type tFloat32 struct {
	listeners []Float32Listener
	value     float32
	filter    Float32Filter
}

type tFloat32AB struct {
	tFloat32
	parentA Float32
	parentB Float32
}

type tFloat32BooleanAB struct {
	tBoolean
	parentA Float32
	parentB Float32
}

type tFloat32Divide struct {
	tFloat32AB
}

type tFloat32Equal struct {
	tFloat32BooleanAB
}

type tFloat32Greater struct {
	tFloat32BooleanAB
}

type tFloat32GreaterOrEqual struct {
	tFloat32BooleanAB
}

type tFloat32Less struct {
	tFloat32BooleanAB
}

type tFloat32LessOrEqual struct {
	tFloat32BooleanAB
}

type tFloat32Minus struct {
	tFloat32AB
}

type tFloat32Multiply struct {
	tFloat32AB
}

type tFloat32NotEqual struct {
	tFloat32BooleanAB
}

type tFloat32Plus struct {
	tFloat32AB
}

func (float32Value *tFloat32) AddListener(listener Float32Listener) {
	if !containsFloat32Listener(float32Value.listeners, listener) {
		float32Value.listeners = append(float32Value.listeners, listener)
	}
}

func (float32Value *tFloat32) Divide(float32ValueB Float32) Float32 {
	float32Divide := new(tFloat32Divide)
	float32Divide.parentA = float32Value
	float32Divide.parentB = float32ValueB
	float32Value.AddListener(float32Divide)
	float32ValueB.AddListener(float32Divide)
	return float32Divide
}

func (float32Value *tFloat32) EqualTo(float32ValueB Float32) Boolean {
	float32Equal := new(tFloat32Equal)
	float32Equal.parentA = float32Value
	float32Equal.parentB = float32ValueB
	float32Value.AddListener(float32Equal)
	float32ValueB.AddListener(float32Equal)
	return float32Equal
}

func (float32Value *tFloat32) Float64() Float64 {
	float64Value := new(tFloat64)
	float32Value.AddListener(float64Value)
	return float64Value
}

func (float32Value *tFloat32) GreaterThan(float32ValueB Float32) Boolean {
	float32Greater := new(tFloat32Greater)
	float32Greater.parentA = float32Value
	float32Greater.parentB = float32ValueB
	float32Value.AddListener(float32Greater)
	float32ValueB.AddListener(float32Greater)
	return float32Greater
}

func (float32Value *tFloat32) GreaterThanOrEqualTo(float32ValueB Float32) Boolean {
	float32GreaterOrEqual := new(tFloat32GreaterOrEqual)
	float32GreaterOrEqual.parentA = float32Value
	float32GreaterOrEqual.parentB = float32ValueB
	float32Value.AddListener(float32GreaterOrEqual)
	float32ValueB.AddListener(float32GreaterOrEqual)
	return float32GreaterOrEqual
}

func (float32Value *tFloat32) Int() Int {
	intValue := new(tInt)
	float32Value.AddListener(intValue)
	return intValue
}

func (float32Value *tFloat32) LessThan(float32ValueB Float32) Boolean {
	float32Less := new(tFloat32Less)
	float32Less.parentA = float32Value
	float32Less.parentB = float32ValueB
	float32Value.AddListener(float32Less)
	float32ValueB.AddListener(float32Less)
	return float32Less
}

func (float32Value *tFloat32) LessThanOrEqualTo(float32ValueB Float32) Boolean {
	float32LessOrEqual := new(tFloat32LessOrEqual)
	float32LessOrEqual.parentA = float32Value
	float32LessOrEqual.parentB = float32ValueB
	float32Value.AddListener(float32LessOrEqual)
	float32ValueB.AddListener(float32LessOrEqual)
	return float32LessOrEqual
}

func (float32Value *tFloat32) Minus(float32ValueB Float32) Float32 {
	float32Minus := new(tFloat32Minus)
	float32Minus.parentA = float32Value
	float32Minus.parentB = float32ValueB
	float32Value.AddListener(float32Minus)
	float32ValueB.AddListener(float32Minus)
	return float32Minus
}

func (float32Value *tFloat32) Multiply(float32ValueB Float32) Float32 {
	float32Multiply := new(tFloat32Multiply)
	float32Multiply.parentA = float32Value
	float32Multiply.parentB = float32ValueB
	float32Value.AddListener(float32Multiply)
	float32ValueB.AddListener(float32Multiply)
	return float32Multiply
}

func (float32Value *tFloat32) NotEqualTo(float32ValueB Float32) Boolean {
	float32ValueNotEqual := new(tFloat32NotEqual)
	float32ValueNotEqual.parentA = float32Value
	float32ValueNotEqual.parentB = float32ValueB
	float32Value.AddListener(float32ValueNotEqual)
	float32ValueB.AddListener(float32ValueNotEqual)
	return float32ValueNotEqual
}

func (float32Value *tFloat32) Plus(float32ValueB Float32) Float32 {
	float32Plus := new(tFloat32Plus)
	float32Plus.parentA = float32Value
	float32Plus.parentB = float32ValueB
	float32Value.AddListener(float32Plus)
	float32ValueB.AddListener(float32Plus)
	return float32Plus
}

func (float32Value *tFloat32) RemoveListener(listener Float32Listener) {
	i := indexFloat32Listener(float32Value.listeners, listener)
	if i >= 0 {
		float32Value.listeners = removeFloat32Listener(float32Value.listeners, i)
	}
}

func (float32Value *tFloat32) Set(newValue float32) {
	oldValue := float32Value.value
	if float32Value.filter != nil {
		newValue = float32Value.filter.FilterFloat32(float32Value, oldValue, newValue)
	}
	if float32Value.value != newValue {
		observable := Float32(float32Value)
		float32Value.value = newValue
		for _, listener := range float32Value.listeners {
			listener.Float32Changed(observable, oldValue, newValue)
		}
	}
}

func (float32Value *tFloat32) SetFilter(filter Float32Filter) {
	float32Value.filter = filter
}

func (float32Value *tFloat32) String() String {
	stringValue := new(tString)
	float32Value.AddListener(stringValue)
	return stringValue
}

func (float32Value *tFloat32) Value() float32 {
	return float32Value.value
}

func (float32Value *tFloat32) Float64Changed(observable Float64, oldValue, newValue float64) {
	float32Value.Set(float32(newValue))
}

func (float32Value *tFloat32) IntChanged(observable Int, oldValue, newValue int) {
	float32Value.Set(float32(newValue))
}

func (float32Value *tFloat32) StringChanged(observable String, oldValue, newValue string) {
	if val, err := strconv.ParseFloat(newValue, 32); err == nil {
		float32Value.Set(float32(val))
	}
}

func (float32Value *tFloat32Divide) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(newValue / float32Value.parentB.Value())
	} else {
		float32Value.Set(float32Value.parentA.Value() / newValue)
	}
}

func (float32Value *tFloat32Equal) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() == newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() == newValue)
	}
}

func (float32Value *tFloat32Greater) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() < newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() > newValue)
	}
}

func (float32Value *tFloat32GreaterOrEqual) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() <= newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() >= newValue)
	}
}

func (float32Value *tFloat32Less) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() > newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() < newValue)
	}
}

func (float32Value *tFloat32LessOrEqual) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() >= newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() <= newValue)
	}
}

func (float32Value *tFloat32Minus) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(newValue - float32Value.parentB.Value())
	} else {
		float32Value.Set(float32Value.parentA.Value() - newValue)
	}
}

func (float32Value *tFloat32Multiply) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() * newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() * newValue)
	}
}

func (float32Value *tFloat32NotEqual) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() != newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() != newValue)
	}
}

func (float32Value *tFloat32Plus) Float32Changed(observable Float32, oldValue, newValue float32) {
	if float32Value.parentA == observable {
		float32Value.Set(float32Value.parentB.Value() + newValue)
	} else {
		float32Value.Set(float32Value.parentA.Value() + newValue)
	}
}
