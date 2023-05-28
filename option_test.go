package optional

import (
	"bytes"
	"testing"
)

func TestOptional(t *testing.T) {
	t.Parallel()
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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var o Optional[int]
			if tc.ok {
				o = New(tc.value)
			} else {
				o = None[int]()
			}
			get, ok := o.Get()
			Equal(t, tc.expectedOk, ok)
			Equal(t, tc.expected, get)
			if tc.expectedOk {
				Equal(t, tc.expected, o.GetOrElse(777))
			} else {
				Equal(t, 777, o.GetOrElse(777))
			}

			if tc.expectedOk {
				NotPanics(t, func() {
					o.Must()
				})
			} else {
				Panics(t, func() {
					o.Must()
				})
				return
			}
			Equal(t, tc.expected, o.Must())
		})

	}
}

func TestOptional_String(t *testing.T) {
	t.Parallel()
	t.Run("string", func(t *testing.T) {
		t.Parallel()
		opStr := New("case1")
		Equal(t, "Some(string)[case1]", opStr.String())
	})
	t.Run("string none", func(t *testing.T) {
		t.Parallel()
		noneSrt := None[string]()
		Equal(t, "None(string)[]", noneSrt.String())
	})
	t.Run("Stringer", func(t *testing.T) {
		t.Parallel()
		opStringer := New(bytes.NewBufferString("case2"))
		Equal(t, "Some(*bytes.Buffer)[case2]", opStringer.String())
	})
	t.Run("Stringer none", func(t *testing.T) {
		t.Parallel()
		noneStringer := None[*bytes.Buffer]()
		Equal(t, "None(*bytes.Buffer)[]", noneStringer.String())
	})
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

func Equal[T comparable](t *testing.T, expected T, has T) {
	t.Helper()
	if expected != has {
		t.Errorf("expected %v, got %v", expected, has)
	}
}

func Panics(t *testing.T, fn func()) {
	defer func() {
		if panicMsg := recover(); panicMsg == nil {
			t.Fatal("expect panic, got nil")
		}
	}()
	fn()
}

func NotPanics(t *testing.T, fn func()) {
	defer func() {
		if panicMsg := recover(); panicMsg != nil {
			t.Fatalf("unexpected panic: %v", panicMsg)
		}
	}()
	fn()
}
