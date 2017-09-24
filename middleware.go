package events

// A middleware event callback
type MiddlewareCallback func(string, ...interface{})

// Create a middleware callback adapter that adapts an event message to a single bool
func MiddlewareBoolify(callback func(string, bool)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value bool

		if len(data) > 0 {
			value = data[0].(bool)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single byte
func MiddlewareByteify(callback func(string, byte)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value byte

		if len(data) > 0 {
			value = data[0].(byte)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single byte slice
func MiddlewareByteSliceify(callback func(string, []byte)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value []byte

		if len(data) > 0 {
			value = data[0].([]byte)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single complex64
func MiddlewareComplex64ify(callback func(string, complex64)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value complex64

		if len(data) > 0 {
			value = data[0].(complex64)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single complex128
func MiddlewareComplex128ify(callback func(string, complex128)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value complex128

		if len(data) > 0 {
			value = data[0].(complex128)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single error
func MiddlewareErrorify(callback func(string, error)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value error

		if len(data) > 0 {
			value = data[0].(error)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single float32
func MiddlewareFloat32ify(callback func(string, float32)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value float32

		if len(data) > 0 {
			value = data[0].(float32)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single float64
func MiddlewareFloat64ify(callback func(string, float64)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value float64

		if len(data) > 0 {
			value = data[0].(float64)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single int
func MiddlewareIntify(callback func(string, int)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value int

		if len(data) > 0 {
			value = data[0].(int)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single int8
func MiddlewareInt8ify(callback func(string, int8)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value int8

		if len(data) > 0 {
			value = data[0].(int8)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single int16
func MiddlewareInt16ify(callback func(string, int16)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value int16

		if len(data) > 0 {
			value = data[0].(int16)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single int32
func MiddlewareInt32ify(callback func(string, int32)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value int32

		if len(data) > 0 {
			value = data[0].(int32)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single int64
func MiddlewareInt64ify(callback func(string, int64)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value int64

		if len(data) > 0 {
			value = data[0].(int64)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single rune
func MiddlewareRuneify(callback func(string, rune)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value rune

		if len(data) > 0 {
			value = data[0].(rune)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single string
func MiddlewareStringify(callback func(string, string)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value string

		if len(data) > 0 {
			value = data[0].(string)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single uint
func MiddlewareUintify(callback func(string, uint)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value uint

		if len(data) > 0 {
			value = data[0].(uint)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single uint8
func MiddlewareUint8ify(callback func(string, uint8)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value uint8

		if len(data) > 0 {
			value = data[0].(uint8)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single uint16
func MiddlewareUint16ify(callback func(string, uint16)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value uint16

		if len(data) > 0 {
			value = data[0].(uint16)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single uint32
func MiddlewareUint32ify(callback func(string, uint32)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value uint32

		if len(data) > 0 {
			value = data[0].(uint32)
		}

		callback(channel, value)
	}
}

// Create a middleware callback adapter that adapts an event message to a single uint64
func MiddlewareUint64ify(callback func(string, uint64)) MiddlewareCallback {
	return func(channel string, data ...interface{}) {
		var value uint64

		if len(data) > 0 {
			value = data[0].(uint64)
		}

		callback(channel, value)
	}
}
