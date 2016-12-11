
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package round


func HalfToEven(value float64) int {
	truncValue := int64(value)

	if value > 0 {
		incrementedTruncValue := float64(truncValue) + 0.5

		if incrementedTruncValue > value {
			return int(truncValue)

		} else if incrementedTruncValue < value {
			return int(truncValue) + 1

		} else if truncValue % 2 == 0 {
			return int(truncValue)

		} else {
			return int(truncValue) + 1
		}

	} else {
		incrementedTruncValue := float64(truncValue) - 0.5

		if incrementedTruncValue < value {
			return int(truncValue)

		} else if incrementedTruncValue > value {
			return int(truncValue) - 1

		} else if truncValue % 2 == 0 {
			return int(truncValue)

		} else {
			return int(truncValue) - 1
		}
	}
}

func HalfToOdd(value float64) int {
	truncValue := int64(value)

	if value > 0 {
		incrementedTruncValue := float64(truncValue) + 0.5

		if incrementedTruncValue > value {
			return int(truncValue)

		} else if incrementedTruncValue < value {
			return int(truncValue) + 1

		} else if truncValue % 2 == 1 {
			return int(truncValue)

		} else {
			return int(truncValue) + 1
		}

	} else {
		incrementedTruncValue := float64(truncValue) - 0.5

		if incrementedTruncValue < value {
			return int(truncValue)

		} else if incrementedTruncValue > value {
			return int(truncValue) - 1

		} else if truncValue % 2 == -1 {
			return int(truncValue)

		} else {
			return int(truncValue) - 1
		}
	}
}

func HalfAwayFromZero(value float64) int {
	return AsInt(value)
}

func Trunc(value float64) int {
	return int(value)
}

func Up(value float64) int {
	truncValue := int64(value)

	if float64(truncValue) != value {
		return int(value + 1.0)

	} else {
		return int(truncValue)
	}
}

func Down(value float64) int {
	truncValue := int64(value)

	if float64(truncValue) != value {
		return int(value - 1.0)

	} else {
		return int(truncValue)
	}
}

func AsInt(value float64) int {
	if value > 0 {
		return int(value + 0.5)

	} else {
		return int(value - 0.5)
	}
}

func AsInt64(value float64) int64 {
	if value > 0 {
		return int64(value + 0.5)

	} else {
		return int64(value - 0.5)
	}
}

func AsInt32(value float64) int32 {
	return int32(AsInt(value))
}

func AsInt16(value float64) int16 {
	return int16(AsInt(value))
}

func AsInt8(value float64) int8 {
	return int8(AsInt(value))
}

func AsUInt(value float64) uint {
	if value > 0 {
		return uint(value + 0.5)

	} else {
		return uint(value - 0.5)
	}
}

func AsUInt64(value float64) uint64 {
	if value > 0 {
		return uint64(value + 0.5)

	} else {
		return uint64(value - 0.5)
	}
}

func AsUInt32(value float64) uint32 {
	return uint32(AsUInt(value))
}

func AsUInt16(value float64) uint16 {
	return uint16(AsUInt(value))
}

func AsUInt8(value float64) uint8 {
	return uint8(AsUInt(value))
}

