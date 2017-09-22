package events

type Callback func(...interface{})

// Create a callback adapter that adapts an event message to a single bool
func Boolify(callback func(bool)) Callback {
	return func(data ...interface{}) {
		var value bool

		if len(data) > 0 {
			value = data[0].(bool)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single byte
func Byteify(callback func(byte)) Callback {
	return func(data ...interface{}) {
		var value byte

		if len(data) > 0 {
			value = data[0].(byte)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single byte slice
func ByteSliceify(callback func([]byte)) Callback {
	return func(data ...interface{}) {
		var value []byte

		if len(data) > 0 {
			value = data[0].([]byte)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single complex64
func Complex64ify(callback func(complex64)) Callback {
	return func(data ...interface{}) {
		var value complex64

		if len(data) > 0 {
			value = data[0].(complex64)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single complex128
func Complex128ify(callback func(complex128)) Callback {
	return func(data ...interface{}) {
		var value complex128

		if len(data) > 0 {
			value = data[0].(complex128)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single error
func Errorify(callback func(error)) Callback {
	return func(data ...interface{}) {
		var value error

		if len(data) > 0 {
			value = data[0].(error)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single float32
func Float32ify(callback func(float32)) Callback {
	return func(data ...interface{}) {
		var value float32

		if len(data) > 0 {
			value = data[0].(float32)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single float64
func Float64ify(callback func(float64)) Callback {
	return func(data ...interface{}) {
		var value float64

		if len(data) > 0 {
			value = data[0].(float64)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single int
func Intify(callback func(int)) Callback {
	return func(data ...interface{}) {
		var value int

		if len(data) > 0 {
			value = data[0].(int)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single int8
func Int8ify(callback func(int8)) Callback {
	return func(data ...interface{}) {
		var value int8

		if len(data) > 0 {
			value = data[0].(int8)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single int16
func Int16ify(callback func(int16)) Callback {
	return func(data ...interface{}) {
		var value int16

		if len(data) > 0 {
			value = data[0].(int16)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single int32
func Int32ify(callback func(int32)) Callback {
	return func(data ...interface{}) {
		var value int32

		if len(data) > 0 {
			value = data[0].(int32)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single int64
func Int64ify(callback func(int64)) Callback {
	return func(data ...interface{}) {
		var value int64

		if len(data) > 0 {
			value = data[0].(int64)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single rune
func Runeify(callback func(rune)) Callback {
	return func(data ...interface{}) {
		var value rune

		if len(data) > 0 {
			value = data[0].(rune)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single string
func Stringify(callback func(string)) Callback {
	return func(data ...interface{}) {
		var value string

		if len(data) > 0 {
			value = data[0].(string)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single uint
func Uintify(callback func(uint)) Callback {
	return func(data ...interface{}) {
		var value uint

		if len(data) > 0 {
			value = data[0].(uint)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single uint8
func Uint8ify(callback func(uint8)) Callback {
	return func(data ...interface{}) {
		var value uint8

		if len(data) > 0 {
			value = data[0].(uint8)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single uint16
func Uint16ify(callback func(uint16)) Callback {
	return func(data ...interface{}) {
		var value uint16

		if len(data) > 0 {
			value = data[0].(uint16)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single uint32
func Uint32ify(callback func(uint32)) Callback {
	return func(data ...interface{}) {
		var value uint32

		if len(data) > 0 {
			value = data[0].(uint32)
		}

		callback(value)
	}
}

// Create a callback adapter that adapts an event message to a single uint64
func Uint64ify(callback func(uint64)) Callback {
	return func(data ...interface{}) {
		var value uint64

		if len(data) > 0 {
			value = data[0].(uint64)
		}

		callback(value)
	}
}
