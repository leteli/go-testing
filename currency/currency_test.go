package currency

import (
	"fmt"
	"testing"
)

type mockConverter struct {
	lastAmount float64
	lastFrom   string
	lastTo     string
	calls      int
	result     float64
}

func NewMockConverter() *mockConverter {
	return &mockConverter{result: 42.0}
}

func (mc *mockConverter) Convert(amount float64, from, to string) float64 {
	if mc == nil {
		return 0.0
	}
	mc.lastAmount, mc.lastFrom, mc.lastTo = amount, from, to
	mc.calls++
	if amount <= 0 {
		return 0.0
	}
	return mc.result
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		amount float64
		from   string
		to     string
		calls  int
		want   float64
	}{
		{amount: 600.0, from: "USD", to: "RUB", want: 42.0},
		{amount: 700000.0, from: "USD", to: "RUB", want: 42.0},
		{amount: -1200.0, from: "USD", to: "RUB", want: 0.0},
		{amount: 0.0, from: "USD", to: "RUB", want: 0.0},
	}
	c := NewMockConverter()
	for _, tc := range cases {
		name := fmt.Sprintf("%v_%q_to_%q", tc.amount, tc.from, tc.to)
		t.Run(name, func(t *testing.T) {
			res := PriceIn(tc.amount, tc.from, tc.to, c)
			assertEqual[float64](t, res, tc.want)
			assertEqual[string](t, c.lastFrom, tc.from)
			assertEqual[string](t, c.lastTo, tc.to)
		})
	}
	assertEqual(t, c.calls, len(cases))
}
