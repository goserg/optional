package optional

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptional(t *testing.T) {
	type TestCase[T any] struct {
		name       string
		value      T
		ok         bool
		expected   T
		expectedOk bool
	}
	tt := []TestCase[int]{
		{
			name:       "ok value",
			value:      10,
			ok:         true,
			expected:   10,
			expectedOk: true,
		},
		{
			name:       "ok zero value",
			value:      0,
			ok:         true,
			expected:   0,
			expectedOk: true,
		},
		{
			name:       "none value",
			value:      0,
			ok:         false,
			expected:   0,
			expectedOk: false,
		},
	}
	for _, tc := range tt {
		var o Optional[int]
		if tc.ok {
			o = New(tc.value)
		} else {
			o = None[int]()
		}
		get, ok := o.Get()
		assert.Equal(t, tc.expectedOk, ok)
		assert.Equal(t, tc.expected, get)
		if tc.expectedOk {
			assert.Equal(t, tc.expected, o.GetOrElse(777))
		} else {
			assert.Equal(t, 777, o.GetOrElse(777))
		}

		if tc.expectedOk {
			require.NotPanics(t, func() {
				o.Must()
			})
		} else {
			require.Panics(t, func() {
				o.Must()
			})
			continue
		}
		assert.Equal(t, tc.expected, o.Must())
	}
}

var x int
var px *int

func BenchmarkInt(b *testing.B) {
	b.Run("optional", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			o := New(i)
			x, _ = o.Get()
		}
	})

	b.Run("pointer", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			o := &i
			px = o
		}
	})
}

type testStr struct {
	value int
}

var s testStr
var ps *testStr

func BenchmarkStruct(b *testing.B) {
	b.Run("struct optional", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			o := New(testStr{value: i})
			s, _ = o.Get()
		}
	})

	b.Run("struct pointer", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			o := &testStr{value: i}
			ps = o
		}
	})
}
