package events

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCallbacks_Boolify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Boolify(func(value bool) {
		called = true
		require.True(t, value)
	}))

	bus.Emit("channel", true)
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Byteify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Byteify(func(value byte) {
		called = true
		require.Equal(t, byte('\n'), value)
	}))

	bus.Emit("channel", byte('\n'))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_ByteSliceify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", ByteSliceify(func(value []byte) {
		called = true
		require.Equal(t, []byte("foo"), value)
	}))

	bus.Emit("channel", []byte("foo"))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Complex64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Complex64ify(func(value complex64) {
		called = true
		require.Equal(t, complex64(1), value)
	}))

	bus.Emit("channel", complex64(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Complex128ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Complex128ify(func(value complex128) {
		called = true
		require.Equal(t, complex128(1), value)
	}))

	bus.Emit("channel", complex128(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Errorify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Errorify(func(value error) {
		called = true
		require.Equal(t, "foo", value.Error())
	}))

	bus.Emit("channel", errors.New("foo"))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Float32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Float32ify(func(value float32) {
		called = true
		require.Equal(t, float32(3.14), value)
	}))

	bus.Emit("channel", float32(3.14))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Float64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Float64ify(func(value float64) {
		called = true
		require.Equal(t, float64(3.14), value)
	}))

	bus.Emit("channel", float64(3.14))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Intify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Intify(func(value int) {
		called = true
		require.Equal(t, int(1), value)
	}))

	bus.Emit("channel", int(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Int8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Int8ify(func(value int8) {
		called = true
		require.Equal(t, int8(1), value)
	}))

	bus.Emit("channel", int8(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Int16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Int16ify(func(value int16) {
		called = true
		require.Equal(t, int16(1), value)
	}))

	bus.Emit("channel", int16(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Int32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Int32ify(func(value int32) {
		called = true
		require.Equal(t, int32(1), value)
	}))

	bus.Emit("channel", int32(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Int64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Int64ify(func(value int64) {
		called = true
		require.Equal(t, int64(1), value)
	}))

	bus.Emit("channel", int64(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Runeify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Runeify(func(value rune) {
		called = true
		require.Equal(t, 'f', value)
	}))

	bus.Emit("channel", 'f')
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Stringify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Stringify(func(value string) {
		called = true
		require.Equal(t, "foo", value)
	}))

	bus.Emit("channel", "foo")
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Uintify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Uintify(func(value uint) {
		called = true
		require.Equal(t, uint(1), value)
	}))

	bus.Emit("channel", uint(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Uint8ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Uint8ify(func(value uint8) {
		called = true
		require.Equal(t, uint8(1), value)
	}))

	bus.Emit("channel", uint8(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Uint16ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Uint16ify(func(value uint16) {
		called = true
		require.Equal(t, uint16(1), value)
	}))

	bus.Emit("channel", uint16(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Uint32ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Uint32ify(func(value uint32) {
		called = true
		require.Equal(t, uint32(1), value)
	}))

	bus.Emit("channel", uint32(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}

func TestCallbacks_Uint64ify(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called := false

	bus.On("channel", Uint64ify(func(value uint64) {
		called = true
		require.Equal(t, uint64(1), value)
	}))

	bus.Emit("channel", uint64(1))
	bus.(*BusImpl).WaitAsync()
	require.True(t, called)
}
