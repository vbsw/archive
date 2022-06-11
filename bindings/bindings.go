/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *      (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package bindings provides bindings for values.
//
// Version 0.4.0.
package bindings

// Boolean is an observable value.
type Boolean interface {
	AddListener(BooleanListener)
	And(Boolean) Boolean
	EqualTo(Boolean) Boolean
	Int() Int
	Not() Boolean
	NotEqualTo(Boolean) Boolean
	Or(Boolean) Boolean
	RemoveListener(BooleanListener)
	Set(bool)
	SetFilter(BooleanFilter)
	String() String
	Value() bool
}

// BooleanFilter provides a function that is called before setting the value of Boolean.
type BooleanFilter interface {
	FilterBoolean(Boolean, bool, bool) bool
}

// BooleanListener is a listener for the observable Boolean. Function BooleanChanged is called
// only when observable value has changed, i.e. new value is not equal to old value.
type BooleanListener interface {
	BooleanChanged(Boolean, bool, bool)
}

// Float64 is an observable value.
type Float64 interface {
	AddListener(Float64Listener)
	Divide(Float64) Float64
	EqualTo(Float64) Boolean
	Float32() Float32
	GreaterThan(Float64) Boolean
	GreaterThanOrEqualTo(Float64) Boolean
	Int() Int
	LessThan(Float64) Boolean
	LessThanOrEqualTo(Float64) Boolean
	Minus(Float64) Float64
	Multiply(Float64) Float64
	NotEqualTo(Float64) Boolean
	Plus(Float64) Float64
	RemoveListener(Float64Listener)
	Set(float64)
	SetFilter(Float64Filter)
	String() String
	Value() float64
}

// Float64Filter provides a function that is called before setting the value of Float64.
type Float64Filter interface {
	FilterFloat64(Float64, float64, float64) float64
}

// Float64Listener is a listener for the observable Float64. Function Float64Changed is called
// only when observable value has changed, i.e. new value is not equal to old value.
type Float64Listener interface {
	Float64Changed(Float64, float64, float64)
}

// Float32 is an observable value.
type Float32 interface {
	AddListener(Float32Listener)
	Divide(Float32) Float32
	EqualTo(Float32) Boolean
	Float64() Float64
	GreaterThan(Float32) Boolean
	GreaterThanOrEqualTo(Float32) Boolean
	Int() Int
	LessThan(Float32) Boolean
	LessThanOrEqualTo(Float32) Boolean
	Minus(Float32) Float32
	Multiply(Float32) Float32
	NotEqualTo(Float32) Boolean
	Plus(Float32) Float32
	RemoveListener(Float32Listener)
	Set(float32)
	SetFilter(Float32Filter)
	String() String
	Value() float32
}

// Float32Filter provides a function that is called before setting the value of Float32.
type Float32Filter interface {
	FilterFloat32(Float32, float32, float32) float32
}

// Float32Listener is a listener for the observable Float32. Function Float32Changed is called
// only when observable value has changed, i.e. new value is not equal to old value.
type Float32Listener interface {
	Float32Changed(Float32, float32, float32)
}

// Int is an observable value.
type Int interface {
	AddListener(IntListener)
	Divide(Int) Int
	EqualTo(Int) Boolean
	Float64() Float64
	Float32() Float32
	GreaterThan(Int) Boolean
	GreaterThanOrEqualTo(Int) Boolean
	LessThan(Int) Boolean
	LessThanOrEqualTo(Int) Boolean
	Minus(Int) Int
	Multiply(Int) Int
	NotEqualTo(Int) Boolean
	Plus(Int) Int
	RemoveListener(IntListener)
	Set(int)
	SetFilter(IntFilter)
	String() String
	Value() int
}

// IntFilter provides a function that is called before setting the value of Int.
type IntFilter interface {
	FilterInt(Int, int, int) int
}

// IntListener is a listener for the observable Int. Function IntChanged is called
// only when observable value has changed, i.e. new value is not equal to old value.
type IntListener interface {
	IntChanged(Int, int, int)
}

// String is an observable value.
type String interface {
	AddListener(StringListener)
	Append(String) String
	Boolean() Boolean
	EqualTo(String) Boolean
	Float64() Float64
	Float32() Float32
	GreaterThan(String) Boolean
	GreaterThanOrEqualTo(String) Boolean
	Int() Int
	LessThan(String) Boolean
	LessThanOrEqualTo(String) Boolean
	LowerCase() String
	NotEqualTo(String) Boolean
	RemoveListener(StringListener)
	Set(string)
	SetFilter(StringFilter)
	Value() string
}

// StringFilter provides a function that is called before setting the value of String.
type StringFilter interface {
	FilterString(String, string, string) string
}

// StringListener is a listener for the observable String. Function StringChanged is called
// only when observable value has changed, i.e. new value is not equal to old value.
type StringListener interface {
	StringChanged(String, string, string)
}

// NewBoolean creates the observable Boolean and returns it.
func NewBoolean(params ...interface{}) Boolean {
	booleanValue := new(tBoolean)
	for _, param := range params {
		if b, ok := param.(bool); ok {
			booleanValue.value = b
		} else if filter, ok := param.(BooleanFilter); ok {
			booleanValue.filter = filter
		}
	}
	return booleanValue
}

// NewFloat64 creates the observable Float64 and returns it.
func NewFloat64(params ...interface{}) Float64 {
	float64Value := new(tFloat64)
	for _, param := range params {
		if filter, ok := param.(Float64Filter); ok {
			float64Value.filter = filter
		} else {
			float64Value.value = toFloat64Ctor(param)
		}
	}
	return float64Value
}

// NewFloat32 creates the observable Float32 and returns it.
func NewFloat32(params ...interface{}) Float32 {
	float32Value := new(tFloat32)
	for _, param := range params {
		if filter, ok := param.(Float32Filter); ok {
			float32Value.filter = filter
		} else {
			float32Value.value = toFloat32Ctor(param)
		}
	}
	return float32Value
}

// NewInt creates the observable Int and returns it.
func NewInt(params ...interface{}) Int {
	intValue := new(tInt)
	for _, param := range params {
		if filter, ok := param.(IntFilter); ok {
			intValue.filter = filter
		} else {
			intValue.value = toIntCtor(param)
		}
	}
	return intValue
}

// NewString creates the observable String and returns it.
func NewString(params ...interface{}) String {
	stringValue := new(tString)
	for _, param := range params {
		if filter, ok := param.(StringFilter); ok {
			stringValue.filter = filter
		}
	}
	return stringValue
}
