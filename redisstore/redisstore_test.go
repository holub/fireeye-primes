package redisstore

import (
	"testing"
)


type testpair struct {
	from int
	expected int
}

func TestFindFrom(t *testing.T) {
	s := New()
	defer s.Close()
	s.Reset()

	for _, p := range []int{1, 7, 77, 78, 81} {
		s.Add(p)
	}

	var tests = []testpair {
		{2, 7},
		{7, 7},
		{50, 77},
		{79, 81},
		{82, -1},
	}
	for _, pair := range tests {
		actual := s.FindFrom(pair.from)
		if actual != pair.expected {
			t.Error("Expected", pair.expected,  ", got", actual)
		}
	}
}

func TestInvert(t *testing.T) {
	s := New()
	defer s.Close()
	s.Reset()

	s.Invert(17)
	if 1 != s.Get(17) {
		t.Error("Expected 1")
	}

	s.Invert(17)
	if 0 != s.Get(17) {
		t.Error("Expected 0")
	}
}
