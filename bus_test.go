package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func doBusEmitBenchmark(b *testing.B, total int) {
	bus := NewBus()

	ch := "channel"
	for i := 0; i < total; i++ {
		bus.On(ch, func(...interface{}) {})
	}

	for n := 0; n < b.N; n++ {
		bus.Emit(ch, nil)
	}
}

func BenchmarkBus_Emit_One(b *testing.B) {
	doBusEmitBenchmark(b, 1)
}

func BenchmarkBus_Emit_Ten(b *testing.B) {
	doBusEmitBenchmark(b, 10)
}

func BenchmarkBus_Emit_Twenty(b *testing.B) {
	doBusEmitBenchmark(b, 20)
}

func BenchmarkBus_Emit_Thirty(b *testing.B) {
	doBusEmitBenchmark(b, 30)
}

func BenchmarkBus_Emit_Forty(b *testing.B) {
	doBusEmitBenchmark(b, 40)
}

func BenchmarkBus_Emit_OneHundred(b *testing.B) {
	doBusEmitBenchmark(b, 100)
}

func TestBus_After(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	calledAfter := false
	calledRegular := false

	err := bus.After(func(channel string, data ...interface{}) {
		require.True(t, calledRegular)
		calledAfter = true
		require.Equal(t, "foo", channel)
		require.Equal(t, []interface{}{"bar"}, data)
	})

	require.Nil(t, err)

	_, err = bus.On("foo", func(data ...interface{}) {
		require.False(t, calledAfter)
		calledRegular = true
	})

	require.Nil(t, err)

	bus.Emit("foo", "bar")

	require.True(t, calledAfter)
	require.True(t, calledRegular)
}

func TestBus_Before(t *testing.T) {
	bus := NewBus()
	require.NotNil(t, bus)

	calledBefore := false
	calledRegular := false

	err := bus.Before(func(channel string, data ...interface{}) {
		require.False(t, calledRegular)
		calledBefore = true
		require.Equal(t, "foo", channel)
		require.Equal(t, []interface{}{"bar"}, data)
	})

	require.Nil(t, err)

	_, err = bus.On("foo", func(data ...interface{}) {
		require.True(t, calledBefore)
		calledRegular = true
	})

	require.Nil(t, err)

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

	_, err := bus.On(ch1, func(data ...interface{}) {
		require.Nil(t, data)
		called1 = true
	})

	require.Nil(t, err)

	_, err = bus.On(ch2, func(data ...interface{}) {
		require.Nil(t, data)
		called2 = true
	})

	require.Nil(t, err)

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

	_, err := bus.On(ch1, func(data ...interface{}) {
		require.Equal(t, []interface{}{ch2}, data)
		called1 = true
	})

	require.Nil(t, err)

	_, err = bus.On(ch2, func(data ...interface{}) {
		require.Equal(t, []interface{}{ch1}, data)
		called2 = true
	})

	require.Nil(t, err)

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
	id, err := bus.On("ch", func(...interface{}) {})
	require.Nil(t, err)
	require.Equal(t, 1, len(bus.(*BusImpl).Subscriptions))
	bus.Remove(id)
	require.Equal(t, 0, len(bus.(*BusImpl).Subscriptions))
}
