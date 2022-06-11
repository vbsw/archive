/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package bindings

import "strconv"

type tInt struct {
	listeners []IntListener
	value     int
	filter    IntFilter
}

type tIntAB struct {
	tInt
	parentA Int
	parentB Int
}

type tIntBooleanAB struct {
	tBoolean
	parentA Int
	parentB Int
}

type tIntDivide struct {
	tIntAB
}

type tIntEqual struct {
	tIntBooleanAB
}

type tIntGreater struct {
	tIntBooleanAB
}

type tIntGreaterOrEqual struct {
	tIntBooleanAB
}

type tIntLess struct {
	tIntBooleanAB
}

type tIntLessOrEqual struct {
	tIntBooleanAB
}

type tIntMinus struct {
	tIntAB
}

type tIntMultiply struct {
	tIntAB
}

type tIntNotEqual struct {
	tIntBooleanAB
}

type tIntPlus struct {
	tIntAB
}

func (intValue *tInt) AddListener(listener IntListener) {
	if !containsIntListener(intValue.listeners, listener) {
		intValue.listeners = append(intValue.listeners, listener)
	}
}

func (intValue *tInt) Divide(intValueB Int) Int {
	intDivide := new(tIntDivide)
	intDivide.parentA = intValue
	intDivide.parentB = intValueB
	intValue.AddListener(intDivide)
	intValueB.AddListener(intDivide)
	return intDivide
}

func (intValue *tInt) EqualTo(intValueB Int) Boolean {
	intEqual := new(tIntEqual)
	intEqual.parentA = intValue
	intEqual.parentB = intValueB
	intValue.AddListener(intEqual)
	intValueB.AddListener(intEqual)
	return intEqual
}

func (intValue *tInt) Float64() Float64 {
	float64Value := new(tFloat64)
	intValue.AddListener(float64Value)
	return float64Value
}

func (intValue *tInt) Float32() Float32 {
	float32Value := new(tFloat32)
	intValue.AddListener(float32Value)
	return float32Value
}

func (intValue *tInt) GreaterThan(intValueB Int) Boolean {
	intGreater := new(tIntGreater)
	intGreater.parentA = intValue
	intGreater.parentB = intValueB
	intValue.AddListener(intGreater)
	intValueB.AddListener(intGreater)
	return intGreater
}

func (intValue *tInt) GreaterThanOrEqualTo(intValueB Int) Boolean {
	intGreaterOrEqual := new(tIntGreaterOrEqual)
	intGreaterOrEqual.parentA = intValue
	intGreaterOrEqual.parentB = intValueB
	intValue.AddListener(intGreaterOrEqual)
	intValueB.AddListener(intGreaterOrEqual)
	return intGreaterOrEqual
}

func (intValue *tInt) LessThan(intValueB Int) Boolean {
	intLess := new(tIntLess)
	intLess.parentA = intValue
	intLess.parentB = intValueB
	intValue.AddListener(intLess)
	intValueB.AddListener(intLess)
	return intLess
}

func (intValue *tInt) LessThanOrEqualTo(intValueB Int) Boolean {
	intLessOrEqual := new(tIntLessOrEqual)
	intLessOrEqual.parentA = intValue
	intLessOrEqual.parentB = intValueB
	intValue.AddListener(intLessOrEqual)
	intValueB.AddListener(intLessOrEqual)
	return intLessOrEqual
}

func (intValue *tInt) Minus(intValueB Int) Int {
	intMinus := new(tIntMinus)
	intMinus.parentA = intValue
	intMinus.parentB = intValueB
	intValue.AddListener(intMinus)
	intValueB.AddListener(intMinus)
	return intMinus
}

func (intValue *tInt) Multiply(intValueB Int) Int {
	intMultiply := new(tIntMultiply)
	intMultiply.parentA = intValue
	intMultiply.parentB = intValueB
	intValue.AddListener(intMultiply)
	intValueB.AddListener(intMultiply)
	return intMultiply
}

func (intValue *tInt) NotEqualTo(intValueB Int) Boolean {
	intValueNotEqual := new(tIntNotEqual)
	intValueNotEqual.parentA = intValue
	intValueNotEqual.parentB = intValueB
	intValue.AddListener(intValueNotEqual)
	intValueB.AddListener(intValueNotEqual)
	return intValueNotEqual
}

func (intValue *tInt) Plus(intValueB Int) Int {
	intPlus := new(tIntPlus)
	intPlus.parentA = intValue
	intPlus.parentB = intValueB
	intValue.AddListener(intPlus)
	intValueB.AddListener(intPlus)
	return intPlus
}

func (intValue *tInt) RemoveListener(listener IntListener) {
	i := indexIntListener(intValue.listeners, listener)
	if i >= 0 {
		intValue.listeners = removeIntListener(intValue.listeners, i)
	}
}

func (intValue *tInt) Set(newValue int) {
	oldValue := intValue.value
	if intValue.filter != nil {
		newValue = intValue.filter.FilterInt(intValue, oldValue, newValue)
	}
	if intValue.value != newValue {
		observable := Int(intValue)
		intValue.value = newValue
		for _, listener := range intValue.listeners {
			listener.IntChanged(observable, oldValue, newValue)
		}
	}
}

func (intValue *tInt) SetFilter(filter IntFilter) {
	intValue.filter = filter
}

func (intValue *tInt) String() String {
	stringValue := new(tString)
	intValue.AddListener(stringValue)
	return stringValue
}

func (intValue *tInt) Value() int {
	return intValue.value
}

func (intValue *tInt) BooleanChanged(observable Boolean, oldValue, newValue bool) {
	if newValue {
		intValue.Set(1)
	} else {
		intValue.Set(0)
	}
}

func (intValue *tInt) Float64Changed(observable Float64, oldValue, newValue float64) {
	intValue.Set(int(newValue))
}

func (intValue *tInt) Float32Changed(observable Float32, oldValue, newValue float32) {
	intValue.Set(int(newValue))
}

func (intValue *tInt) StringChanged(observable String, oldValue, newValue string) {
	if val, err := strconv.Atoi(newValue); err == nil {
		intValue.Set(val)
	}
}

func (intValue *tIntDivide) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(newValue / intValue.parentB.Value())
	} else {
		intValue.Set(intValue.parentA.Value() / newValue)
	}
}

func (intValue *tIntEqual) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() == newValue)
	} else {
		intValue.Set(intValue.parentA.Value() == newValue)
	}
}

func (intValue *tIntGreater) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() < newValue)
	} else {
		intValue.Set(intValue.parentA.Value() > newValue)
	}
}

func (intValue *tIntGreaterOrEqual) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() <= newValue)
	} else {
		intValue.Set(intValue.parentA.Value() >= newValue)
	}
}

func (intValue *tIntLess) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() > newValue)
	} else {
		intValue.Set(intValue.parentA.Value() < newValue)
	}
}

func (intValue *tIntLessOrEqual) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() >= newValue)
	} else {
		intValue.Set(intValue.parentA.Value() <= newValue)
	}
}

func (intValue *tIntMinus) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(newValue - intValue.parentB.Value())
	} else {
		intValue.Set(intValue.parentA.Value() - newValue)
	}
}

func (intValue *tIntMultiply) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() * newValue)
	} else {
		intValue.Set(intValue.parentA.Value() * newValue)
	}
}

func (intValue *tIntNotEqual) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() != newValue)
	} else {
		intValue.Set(intValue.parentA.Value() != newValue)
	}
}

func (intValue *tIntPlus) IntChanged(observable Int, oldValue, newValue int) {
	if intValue.parentA == observable {
		intValue.Set(intValue.parentB.Value() + newValue)
	} else {
		intValue.Set(intValue.parentA.Value() + newValue)
	}
}
