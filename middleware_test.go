package events

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

const MiddlewareTestChannel = "channel"

func TestMiddleware_MiddlewareBoolify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareBoolify(func(channel string, value bool) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.True(t, value)
	}))

	bus.Emit(MiddlewareTestChannel, true)
	require.True(t, called)
}

func TestMiddleware_MiddlewareByteify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareByteify(func(channel string, value byte) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, byte('\n'), value)
	}))

	bus.Emit(MiddlewareTestChannel, byte('\n'))
	require.True(t, called)
}

func TestMiddleware_MiddlewareByteSliceify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareByteSliceify(func(channel string, value []byte) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, []byte("foo"), value)
	}))

	bus.Emit(MiddlewareTestChannel, []byte("foo"))
	require.True(t, called)
}

func TestMiddleware_MiddlewareComplex64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareComplex64ify(func(channel string, value complex64) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, complex64(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, complex64(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareComplex128ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareComplex128ify(func(channel string, value complex128) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, complex128(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, complex128(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareErrorify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareErrorify(func(channel string, value error) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, "foo", value.Error())
	}))

	bus.Emit(MiddlewareTestChannel, errors.New("foo"))
	require.True(t, called)
}

func TestMiddleware_MiddlewareFloat32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareFloat32ify(func(channel string, value float32) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, float32(3.14), value)
	}))

	bus.Emit(MiddlewareTestChannel, float32(3.14))
	require.True(t, called)
}

func TestMiddleware_MiddlewareFloat64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareFloat64ify(func(channel string, value float64) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, float64(3.14), value)
	}))

	bus.Emit(MiddlewareTestChannel, float64(3.14))
	require.True(t, called)
}

func TestMiddleware_MiddlewareIntify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareIntify(func(channel string, value int) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, int(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, int(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareInt8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareInt8ify(func(channel string, value int8) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, int8(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, int8(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareInt16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareInt16ify(func(channel string, value int16) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, int16(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, int16(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareInt32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareInt32ify(func(channel string, value int32) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, int32(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, int32(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareInt64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareInt64ify(func(channel string, value int64) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, int64(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, int64(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareRuneify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareRuneify(func(channel string, value rune) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, 'f', value)
	}))

	bus.Emit(MiddlewareTestChannel, 'f')
	require.True(t, called)
}

func TestMiddleware_MiddlewareStringify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareStringify(func(channel string, value string) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, "foo", value)
	}))

	bus.Emit(MiddlewareTestChannel, "foo")
	require.True(t, called)
}

func TestMiddleware_MiddlewareUintify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareUintify(func(channel string, value uint) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, uint(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, uint(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareUint8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareUint8ify(func(channel string, value uint8) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, uint8(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, uint8(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareUint16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareUint16ify(func(channel string, value uint16) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, uint16(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, uint16(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareUint32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareUint32ify(func(channel string, value uint32) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, uint32(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, uint32(1))
	require.True(t, called)
}

func TestMiddleware_MiddlewareUint64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.Before(MiddlewareUint64ify(func(channel string, value uint64) {
		called = true
		require.Equal(t, MiddlewareTestChannel, channel)
		require.Equal(t, uint64(1), value)
	}))

	bus.Emit(MiddlewareTestChannel, uint64(1))
	require.True(t, called)
}
