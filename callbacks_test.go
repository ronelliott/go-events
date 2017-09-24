package events

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

const CallbacksTestChannel = "channel"

func TestCallbacks_Boolify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Boolify(func(value bool) {
		called = true
		require.True(t, value)
	}))

	bus.Emit(CallbacksTestChannel, true)
	require.True(t, called)
}

func TestCallbacks_Byteify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Byteify(func(value byte) {
		called = true
		require.Equal(t, byte('\n'), value)
	}))

	bus.Emit(CallbacksTestChannel, byte('\n'))
	require.True(t, called)
}

func TestCallbacks_ByteSliceify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, ByteSliceify(func(value []byte) {
		called = true
		require.Equal(t, []byte("foo"), value)
	}))

	bus.Emit(CallbacksTestChannel, []byte("foo"))
	require.True(t, called)
}

func TestCallbacks_Complex64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Complex64ify(func(value complex64) {
		called = true
		require.Equal(t, complex64(1), value)
	}))

	bus.Emit(CallbacksTestChannel, complex64(1))
	require.True(t, called)
}

func TestCallbacks_Complex128ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Complex128ify(func(value complex128) {
		called = true
		require.Equal(t, complex128(1), value)
	}))

	bus.Emit(CallbacksTestChannel, complex128(1))
	require.True(t, called)
}

func TestCallbacks_Errorify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Errorify(func(value error) {
		called = true
		require.Equal(t, "foo", value.Error())
	}))

	bus.Emit(CallbacksTestChannel, errors.New("foo"))
	require.True(t, called)
}

func TestCallbacks_Float32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Float32ify(func(value float32) {
		called = true
		require.Equal(t, float32(3.14), value)
	}))

	bus.Emit(CallbacksTestChannel, float32(3.14))
	require.True(t, called)
}

func TestCallbacks_Float64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Float64ify(func(value float64) {
		called = true
		require.Equal(t, float64(3.14), value)
	}))

	bus.Emit(CallbacksTestChannel, float64(3.14))
	require.True(t, called)
}

func TestCallbacks_Intify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Intify(func(value int) {
		called = true
		require.Equal(t, int(1), value)
	}))

	bus.Emit(CallbacksTestChannel, int(1))
	require.True(t, called)
}

func TestCallbacks_Int8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Int8ify(func(value int8) {
		called = true
		require.Equal(t, int8(1), value)
	}))

	bus.Emit(CallbacksTestChannel, int8(1))
	require.True(t, called)
}

func TestCallbacks_Int16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Int16ify(func(value int16) {
		called = true
		require.Equal(t, int16(1), value)
	}))

	bus.Emit(CallbacksTestChannel, int16(1))
	require.True(t, called)
}

func TestCallbacks_Int32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Int32ify(func(value int32) {
		called = true
		require.Equal(t, int32(1), value)
	}))

	bus.Emit(CallbacksTestChannel, int32(1))
	require.True(t, called)
}

func TestCallbacks_Int64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Int64ify(func(value int64) {
		called = true
		require.Equal(t, int64(1), value)
	}))

	bus.Emit(CallbacksTestChannel, int64(1))
	require.True(t, called)
}

func TestCallbacks_Runeify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Runeify(func(value rune) {
		called = true
		require.Equal(t, 'f', value)
	}))

	bus.Emit(CallbacksTestChannel, 'f')
	require.True(t, called)
}

func TestCallbacks_Stringify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Stringify(func(value string) {
		called = true
		require.Equal(t, "foo", value)
	}))

	bus.Emit(CallbacksTestChannel, "foo")
	require.True(t, called)
}

func TestCallbacks_Uintify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Uintify(func(value uint) {
		called = true
		require.Equal(t, uint(1), value)
	}))

	bus.Emit(CallbacksTestChannel, uint(1))
	require.True(t, called)
}

func TestCallbacks_Uint8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Uint8ify(func(value uint8) {
		called = true
		require.Equal(t, uint8(1), value)
	}))

	bus.Emit(CallbacksTestChannel, uint8(1))
	require.True(t, called)
}

func TestCallbacks_Uint16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Uint16ify(func(value uint16) {
		called = true
		require.Equal(t, uint16(1), value)
	}))

	bus.Emit(CallbacksTestChannel, uint16(1))
	require.True(t, called)
}

func TestCallbacks_Uint32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Uint32ify(func(value uint32) {
		called = true
		require.Equal(t, uint32(1), value)
	}))

	bus.Emit(CallbacksTestChannel, uint32(1))
	require.True(t, called)
}

func TestCallbacks_Uint64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On(CallbacksTestChannel, Uint64ify(func(value uint64) {
		called = true
		require.Equal(t, uint64(1), value)
	}))

	bus.Emit(CallbacksTestChannel, uint64(1))
	require.True(t, called)
}
