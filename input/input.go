/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


package input




type EventType int8
type KeyEventType int8
type MouseEventType int8
type Position2IntEventType int8
type Position3IntEventType int8
type Position2FloatEventType int8
type Position3FloatEventType int8

const private_INITIAL_CAPACITY = 7

const (
	KEY_EVENT EventType = iota
	MOUSE_EVENT
	POSITION_2INT_EVENT
	POSITION_3INT_EVENT
	POSITION_2FLOAT_EVENT
	POSITION_3FLOAT_EVENT
)

const (
	KEY_DOWN KeyEventType = iota
	KEY_PRESSED
	KEY_RELEASED
)

const (
	MOUSE_MOVED MouseEventType = iota
	MOUSE_BUTTON_PRESSED
	MOUSE_BUTTON_RELEASED
	MOUSE_WHEEL_MOVED
)

const (
	POSITION_2INT_NORMAL Position2IntEventType = iota
	POSITION_2INT_SPECIAL1
	POSITION_2INT_SPECIAL2
	POSITION_2INT_SPECIAL3
)

const (
	POSITION_3INT_NORMAL Position3IntEventType = iota
	POSITION_3INT_SPECIAL1
	POSITION_3INT_SPECIAL2
	POSITION_3INT_SPECIAL3
)

const (
	POSITION_2FLOAT_NORMAL Position2FloatEventType = iota
	POSITION_2FLOAT_SPECIAL1
	POSITION_2FLOAT_SPECIAL2
	POSITION_2FLOAT_SPECIAL3
)

const (
	POSITION_3FLOAT_NORMAL Position3FloatEventType = iota
	POSITION_3FLOAT_SPECIAL1
	POSITION_3FLOAT_SPECIAL2
	POSITION_3FLOAT_SPECIAL3
)

var events []EventType = make([]EventType, 0, private_INITIAL_CAPACITY)
var eventIndices []int = make([]int, 0, private_INITIAL_CAPACITY)
var keyEvents []KeyEventType = make([]KeyEventType, 0, private_INITIAL_CAPACITY)
var keyCodes []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var mouseEvents []MouseEventType = make([]MouseEventType, 0, private_INITIAL_CAPACITY)
var mouseValues []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntEvents []Position2IntEventType = make([]Position2IntEventType, 0, private_INITIAL_CAPACITY)
var position2IntSourceIds []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntXs []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntYs []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntEvents []Position3IntEventType = make([]Position3IntEventType, 0, private_INITIAL_CAPACITY)
var position3IntSourceIds []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntXs []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntYs []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntZs []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2FloatEvents []Position2FloatEventType = make([]Position2FloatEventType, 0, private_INITIAL_CAPACITY)
var position2FloatSourceIds []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2FloatXs []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position2FloatYs []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatEvents []Position3FloatEventType = make([]Position3FloatEventType, 0, private_INITIAL_CAPACITY)
var position3FloatSourceIds []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3FloatXs []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatYs []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatZs []float32 = make([]float32, 0, private_INITIAL_CAPACITY)

var events_w []EventType = make([]EventType, 0, private_INITIAL_CAPACITY)
var eventIndices_w []int = make([]int, 0, private_INITIAL_CAPACITY)
var keyEvents_w []KeyEventType = make([]KeyEventType, 0, private_INITIAL_CAPACITY)
var keyCodes_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var mouseEvents_w []MouseEventType = make([]MouseEventType, 0, private_INITIAL_CAPACITY)
var mouseValues_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntEvents_w []Position2IntEventType = make([]Position2IntEventType, 0, private_INITIAL_CAPACITY)
var position2IntSourceIds_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntXs_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2IntYs_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntEvents_w []Position3IntEventType = make([]Position3IntEventType, 0, private_INITIAL_CAPACITY)
var position3IntSourceIds_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntXs_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntYs_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3IntZs_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2FloatEvents_w []Position2FloatEventType = make([]Position2FloatEventType, 0, private_INITIAL_CAPACITY)
var position2FloatSourceIds_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position2FloatXs_w []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position2FloatYs_w []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatEvents_w []Position3FloatEventType = make([]Position3FloatEventType, 0, private_INITIAL_CAPACITY)
var position3FloatSourceIds_w []int32 = make([]int32, 0, private_INITIAL_CAPACITY)
var position3FloatXs_w []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatYs_w []float32 = make([]float32, 0, private_INITIAL_CAPACITY)
var position3FloatZs_w []float32 = make([]float32, 0, private_INITIAL_CAPACITY)

var index int = -1

func SwapReadWriteBuffers() {
	eventsTmp := events[0:0]
	eventIndicesTmp := eventIndices[0:0]
	keyEventsTmp := keyEvents[0:0]
	keyCodesTmp := keyCodes[0:0]
	mouseValuesTmp := mouseValues[0:0]
	mouseEventsTmp := mouseEvents[0:0]
	position2IntEventsTmp := position2IntEvents[0:0]
	position2IntXsTmp := position2IntXs[0:0]
	position2IntYsTmp := position2IntYs[0:0]
	position2IntSourceIdsTmp := position2IntSourceIds[0:0]
	position3IntEventsTmp := position3IntEvents[0:0]
	position3IntXsTmp := position3IntXs[0:0]
	position3IntYsTmp := position3IntYs[0:0]
	position3IntZsTmp := position3IntZs[0:0]
	position3IntSourceIdsTmp := position3IntSourceIds[0:0]
	position2FloatEventsTmp := position2FloatEvents[0:0]
	position2FloatXsTmp := position2FloatXs[0:0]
	position2FloatYsTmp := position2FloatYs[0:0]
	position2FloatSourceIdsTmp := position2FloatSourceIds[0:0]
	position3FloatEventsTmp := position3FloatEvents[0:0]
	position3FloatXsTmp := position3FloatXs[0:0]
	position3FloatYsTmp := position3FloatYs[0:0]
	position3FloatZsTmp := position3FloatZs[0:0]
	position3FloatSourceIdsTmp := position3FloatSourceIds[0:0]
	events = events_w
	eventIndices = eventIndices_w
	keyEvents = keyEvents_w
	keyCodes = keyCodes_w
	mouseValues = mouseValues_w
	mouseEvents = mouseEvents_w
	position2IntEvents = position2IntEvents_w
	position2IntXs = position2IntXs_w
	position2IntYs = position2IntYs_w
	position2IntSourceIds = position2IntSourceIds_w
	position3IntEvents = position3IntEvents_w
	position3IntXs = position3IntXs_w
	position3IntYs = position3IntYs_w
	position3IntZs = position3IntZs_w
	position3IntSourceIds = position3IntSourceIds_w
	position2FloatEvents = position2FloatEvents_w
	position2FloatXs = position2FloatXs_w
	position2FloatYs = position2FloatYs_w
	position2FloatSourceIds = position2FloatSourceIds_w
	position3FloatEvents = position3FloatEvents_w
	position3FloatXs = position3FloatXs_w
	position3FloatYs = position3FloatYs_w
	position3FloatZs = position3FloatZs_w
	position3FloatSourceIds = position3FloatSourceIds_w
	events_w = eventsTmp
	eventIndices_w = eventIndicesTmp
	keyEvents_w = keyEventsTmp
	keyCodes_w = keyCodesTmp
	mouseValues_w = mouseValuesTmp
	mouseEvents_w = mouseEventsTmp
	position2IntEvents_w = position2IntEventsTmp
	position2IntXs_w = position2IntXsTmp
	position2IntYs_w = position2IntYsTmp
	position2IntSourceIds_w = position2IntSourceIdsTmp
	position3IntEvents_w = position3IntEventsTmp
	position3IntXs_w = position3IntXsTmp
	position3IntYs_w = position3IntYsTmp
	position3IntZs_w = position3IntZsTmp
	position3IntSourceIds_w = position3IntSourceIdsTmp
	position2FloatEvents_w = position2FloatEventsTmp
	position2FloatXs_w = position2FloatXsTmp
	position2FloatYs_w = position2FloatYsTmp
	position2FloatSourceIds_w = position2FloatSourceIdsTmp
	position3FloatEvents_w = position3FloatEventsTmp
	position3FloatXs_w = position3FloatXsTmp
	position3FloatYs_w = position3FloatYsTmp
	position3FloatZs_w = position3FloatZsTmp
	position3FloatSourceIds_w = position3FloatSourceIdsTmp
}

func HasNext() bool {
	index = index + 1

	if index < len(events) {
		return true
	} else {
		index = -1
		return false
	}
}

func Event() EventType {
	return events[index]
}

func HasNextKeyEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == KEY_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func HasNextMouseEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == MOUSE_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func HasNextPosition2IntEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == POSITION_2INT_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func HasNextPosition3IntEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == POSITION_3INT_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func HasNextPosition2FloatEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == POSITION_2FLOAT_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func HasNextPosition3FloatEvent() bool {
	index = index + 1

	for index < len(events) {
		if events[index] == POSITION_3FLOAT_EVENT {
			return true
		}
		index = index + 1
	}
	index = -1

	return false
}

func KeyEvent() KeyEventType {
	keyEventIndex := eventIndices[index]
	return keyEvents[keyEventIndex]
}

func KeyCode() int32 {
	keyEventIndex := eventIndices[index]
	return keyCodes[keyEventIndex]
}

func MouseEvent() MouseEventType {
	mouseEventIndex := eventIndices[index]
	return mouseEvents[mouseEventIndex]
}

func MouseValue() int32 {
	mouseEventIndex := eventIndices[index]
	return mouseValues[mouseEventIndex]
}

func Position2IntEvent() Position2IntEventType {
	position2IntEventIndex := eventIndices[index]
	return position2IntEvents[position2IntEventIndex]
}

func Position2IntSourceId() int32 {
	position2IntEventIndex := eventIndices[index]
	return position2IntSourceIds[position2IntEventIndex]
}

func Position2IntX() int32 {
	position2IntEventIndex := eventIndices[index]
	return position2IntXs[position2IntEventIndex]
}

func Position2IntY() int32 {
	position2IntEventIndex := eventIndices[index]
	return position2IntYs[position2IntEventIndex]
}

func Position3IntEvent() Position3IntEventType {
	position3IntEventIndex := eventIndices[index]
	return position3IntEvents[position3IntEventIndex]
}

func Position3IntSourceId() int32 {
	position3IntEventIndex := eventIndices[index]
	return position3IntSourceIds[position3IntEventIndex]
}

func Position3IntX() int32 {
	position3IntEventIndex := eventIndices[index]
	return position3IntXs[position3IntEventIndex]
}

func Position3IntY() int32 {
	position3IntEventIndex := eventIndices[index]
	return position3IntYs[position3IntEventIndex]
}

func Position3IntZ() int32 {
	position3IntEventIndex := eventIndices[index]
	return position3IntZs[position3IntEventIndex]
}

func Position2FloatEvent() Position2FloatEventType {
	position2FloatEventIndex := eventIndices[index]
	return position2FloatEvents[position2FloatEventIndex]
}

func Position2FloatSourceId() int32 {
	position2FloatEventIndex := eventIndices[index]
	return position2FloatSourceIds[position2FloatEventIndex]
}

func Position2FloatX() float32 {
	position2FloatEventIndex := eventIndices[index]
	return position2FloatXs[position2FloatEventIndex]
}

func Position2FloatY() float32 {
	position2FloatEventIndex := eventIndices[index]
	return position2FloatYs[position2FloatEventIndex]
}

func Position3FloatEvent() Position3FloatEventType {
	position3FloatEventIndex := eventIndices[index]
	return position3FloatEvents[position3FloatEventIndex]
}

func Position3FloatSourceId() int32 {
	position3FloatEventIndex := eventIndices[index]
	return position3FloatSourceIds[position3FloatEventIndex]
}

func Position3FloatX() float32 {
	position3FloatEventIndex := eventIndices[index]
	return position3FloatXs[position3FloatEventIndex]
}

func Position3FloatY() float32 {
	position3FloatEventIndex := eventIndices[index]
	return position3FloatYs[position3FloatEventIndex]
}

func Position3FloatZ() float32 {
	position3FloatEventIndex := eventIndices[index]
	return position3FloatZs[position3FloatEventIndex]
}

func AddKeyEvent(keyEvent KeyEventType, keyCode int32) {
	keyEventIndex := len(keyEvents)
	events_w = append(events_w, KEY_EVENT)
	eventIndices_w = append(eventIndices_w, keyEventIndex)
	keyEvents_w = append(keyEvents_w, keyEvent)
	keyCodes_w = append(keyCodes_w, keyCode)
}

func AddMouseEvent(mouseEvent MouseEventType, value int32) {
	mouseEventIndex := len(mouseEvents)
	events_w = append(events_w, MOUSE_EVENT)
	eventIndices_w = append(eventIndices_w, mouseEventIndex)
	mouseEvents_w = append(mouseEvents_w, mouseEvent)
	mouseValues_w = append(mouseValues_w, value)
}

func AddPosition2IntEvent(position2IntEvent Position2IntEventType, sourceId, x, y int32) {
	position2IntEventIndex := len(position2IntEvents)
	events_w = append(events_w, POSITION_2INT_EVENT)
	eventIndices_w = append(eventIndices_w, position2IntEventIndex)
	position2IntEvents_w = append(position2IntEvents_w, position2IntEvent)
	position2IntSourceIds_w = append(position2IntSourceIds_w, sourceId)
	position2IntXs_w = append(position2IntXs_w, x)
	position2IntYs_w = append(position2IntYs_w, y)
}

func AddPosition3IntEvent(position3IntEvent Position3IntEventType, sourceId, x, y, z int32) {
	position3IntEventIndex := len(position3IntEvents)
	events_w = append(events_w, POSITION_3INT_EVENT)
	eventIndices_w = append(eventIndices_w, position3IntEventIndex)
	position3IntEvents_w = append(position3IntEvents_w, position3IntEvent)
	position3IntSourceIds_w = append(position3IntSourceIds_w, sourceId)
	position3IntXs_w = append(position3IntXs_w, x)
	position3IntYs_w = append(position3IntYs_w, y)
	position3IntZs_w = append(position3IntZs_w, y)
}

func AddPosition2FloatEvent(position2FloatEvent Position2FloatEventType, sourceId int32, x, y float32) {
	position2FloatEventIndex := len(position2FloatEvents)
	events_w = append(events_w, POSITION_2FLOAT_EVENT)
	eventIndices_w = append(eventIndices_w, position2FloatEventIndex)
	position2FloatEvents_w = append(position2FloatEvents_w, position2FloatEvent)
	position2FloatSourceIds_w = append(position2FloatSourceIds_w, sourceId)
	position2FloatXs_w = append(position2FloatXs_w, x)
	position2FloatYs_w = append(position2FloatYs_w, y)
}

func AddPosition3FloatEvent(position3FloatEvent Position3FloatEventType, sourceId int32, x, y, z float32) {
	position3FloatEventIndex := len(position3FloatEvents)
	events_w = append(events_w, POSITION_3FLOAT_EVENT)
	eventIndices_w = append(eventIndices_w, position3FloatEventIndex)
	position3FloatEvents_w = append(position3FloatEvents_w, position3FloatEvent)
	position3FloatSourceIds_w = append(position3FloatSourceIds_w, sourceId)
	position3FloatXs_w = append(position3FloatXs_w, x)
	position3FloatYs_w = append(position3FloatYs_w, y)
	position3FloatZs_w = append(position3FloatZs_w, y)
}



