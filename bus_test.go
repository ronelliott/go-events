package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkBus_Emit_One(b *testing.B) {
	bus := NewBus()

	ch := "channel"
	bus.On(ch, func(...interface{}) {})

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func BenchmarkBus_Emit_Ten(b *testing.B) {
	bus := NewBus()

	ch := "channel"
	for i := 0; i < 10; i++ {
		bus.On(ch, func(...interface{}) {})
	}

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func BenchmarkBus_Emit_Twenty(b *testing.B) {
	bus := NewBus()

	ch := "channel"
	for i := 0; i < 20; i++ {
		bus.On(ch, func(...interface{}) {})
	}

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func BenchmarkBus_Emit_Thirty(b *testing.B) {
	bus := NewBus()

	ch := "channel"
	for i := 0; i < 30; i++ {
		bus.On(ch, func(...interface{}) {})
	}

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func BenchmarkBus_Emit_Forty(b *testing.B) {
	bus := NewBus()

	ch := "channel"
	for i := 0; i < 40; i++ {
		bus.On(ch, func(...interface{}) {})
	}

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func TestBus_After(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	calledAfter := false
	calledRegular := false

	bus.After(func(channel string, data ...interface{}) {
		require.True(t, calledRegular)
		calledAfter = true
		require.Equal(t, "foo", channel)
		require.Equal(t, []interface{}{"bar"}, data)
	})

	bus.On("foo", func(data ...interface{}) {
		require.False(t, calledAfter)
		calledRegular = true
	})

	bus.Emit("foo", "bar")

	require.True(t, calledAfter)
	require.True(t, calledRegular)
}

func TestBus_Before(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	calledBefore := false
	calledRegular := false

	bus.Before(func(channel string, data ...interface{}) {
		require.False(t, calledRegular)
		calledBefore = true
		require.Equal(t, "foo", channel)
		require.Equal(t, []interface{}{"bar"}, data)
	})

	bus.On("foo", func(data ...interface{}) {
		require.True(t, calledBefore)
		calledRegular = true
	})

	bus.Emit("foo", "bar")

	require.True(t, calledBefore)
	require.True(t, calledRegular)
}

func TestBus_Emit_Nil(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called1 := false
	called2 := false

	ch1 := "channel-1"
	ch2 := "channel-2"

	bus.On(ch1, func(data ...interface{}) {
		require.Nil(t, data)
		called1 = true
	})

	bus.On(ch2, func(data ...interface{}) {
		require.Nil(t, data)
		called2 = true
	})

	bus.Emit(ch1)

	require.True(t, called1)
	require.False(t, called2)

	called1 = false
	called2 = false

	bus.Emit(ch1)
	bus.Emit(ch2)

	require.True(t, called1)
	require.True(t, called2)

	called1 = false
	called2 = false

	bus.Emit(ch2)

	require.False(t, called1)
	require.True(t, called2)
}

func TestBus_Emit_Values(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	called1 := false
	called2 := false

	ch1 := "channel-1"
	ch2 := "channel-2"

	bus.On(ch1, func(data ...interface{}) {
		require.Equal(t, []interface{}{ch2}, data)
		called1 = true
	})

	bus.On(ch2, func(data ...interface{}) {
		require.Equal(t, []interface{}{ch1}, data)
		called2 = true
	})

	bus.Emit(ch1, ch2)

	require.True(t, called1)
	require.False(t, called2)

	called1 = false
	called2 = false

	bus.Emit(ch1, ch2)
	bus.Emit(ch2, ch1)

	require.True(t, called1)
	require.True(t, called2)

	called1 = false
	called2 = false

	bus.Emit(ch2, ch1)

	require.False(t, called1)
	require.True(t, called2)
}

func TestBus_Remove(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)
	id := bus.On("ch", func(...interface{}) {})
	require.Equal(t, 1, len(bus.(*BusImpl).Subscriptions))
	bus.Remove(id)
	require.Equal(t, 0, len(bus.(*BusImpl).Subscriptions))
}
