package glob_test

import (
	"testing"

	. "github.com/srikrsna/glob"
)

type MatchTest struct {
	pattern, s string
	match      bool
	err        error
}

var matchTests = []MatchTest{
	{"abc", "abc", true, nil},
	{"*", "abc", true, nil},
	{"*c", "abc", true, nil},
	{"a*", "a", true, nil},
	{"a*", "abc", true, nil},
	{"a*", "ab/c", false, nil},
	{"a*/b", "abc/b", true, nil},
	{"a*/b", "a/c/b", false, nil},
	{"a*b*c*d*e*/f", "axbxcxdxe/f", true, nil},
	{"a*b*c*d*e**f", "axbxcxdxe/f", true, nil},
	{"**", "axbxcxdxe/f", true, nil},
	{"a*b*c*d*e**f", "axbxcxdxe/g", false, nil},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/f", true, nil},
	{"a*b*c*d*e*/f", "axbxcxdxe/xxx/f", false, nil},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/fff", false, nil},
	{"a*b?c*x", "abxbbxdbxebxczzx", true, nil},
	{"a*b?c*x", "abxbbxdbxebxczzy", false, nil},
	{"ab[c]", "abc", true, nil},
	{"ab[b-d]", "abc", true, nil},
	{"ab[e-g]", "abc", false, nil},
	{"ab[!c]", "abc", false, nil},
	{"ab[!b-d]", "abc", false, nil},
	{"ab[!e-g]", "abc", true, nil},
	{"a\\*b", "a*b", true, nil},
	{"a\\*b", "ab", false, nil},
	{"a?b", "a☺b", true, nil},
	{"a[!a]b", "a☺b", true, nil},
	{"a???b", "a☺b", false, nil},
	{"a[!a][!a][!a]b", "a☺b", false, nil},
	{"[a-ζ]*", "α", true, nil},
	{"*[a-ζ]", "A", false, nil},
	{"a?b", "a/b", false, nil},
	{"a*b", "a/b", false, nil},
	{"[\\]a]", "]", true, nil},
	{"[\\-]", "-", true, nil},
	{"[x\\-]", "x", true, nil},
	{"[x\\-]", "-", true, nil},
	{"[x\\-]", "z", false, nil},
	{"[\\-x]", "x", true, nil},
	{"[\\-x]", "-", true, nil},
	{"[\\-x]", "a", false, nil},
	{"[]a]", "]", false, ErrBadPattern},
	{"[-]", "-", false, ErrBadPattern},
	{"[x-]", "x", false, ErrBadPattern},
	{"[x-]", "-", false, ErrBadPattern},
	{"[x-]", "z", false, ErrBadPattern},
	{"[-x]", "x", false, ErrBadPattern},
	{"[-x]", "-", false, ErrBadPattern},
	{"[-x]", "a", false, ErrBadPattern},
	{"\\", "a", false, ErrBadPattern},
	{"[a-b-c]", "a", false, ErrBadPattern},
	{"[", "a", false, ErrBadPattern},
	{"[^", "a", false, ErrBadPattern},
	{"[^bc", "a", false, ErrBadPattern},
	{"a[", "a", false, ErrBadPattern},
	{"a[", "ab", false, ErrBadPattern},
	{"a[", "x", false, ErrBadPattern},
	{"a/b[", "x", false, ErrBadPattern},
	{"*x", "xxx", true, nil},
}

func TestMatch(t *testing.T) {
	for _, tt := range matchTests {
		ok, err := Match(tt.pattern, tt.s)
		if ok != tt.match || err != tt.err {
			t.Errorf("Match(%#q, %#q) = %v, %v want %v, %v", tt.pattern, tt.s, ok, err, tt.match, tt.err)
		}
	}
}

func TestMatchFast(t *testing.T) {
	for _, tt := range matchTests {
		ok := MatchFast(tt.pattern, tt.s)
		if ok != tt.match {
			t.Errorf("MatchFast(%#q, %#q) = %v want %v", tt.pattern, tt.s, ok, tt.match)
		}
	}
}

func BenchmarkMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range matchTests {
			ok, err := Match(tt.pattern, tt.s)
			if ok != tt.match || err != tt.err {
				b.Errorf("Match(%#q, %#q) = %v, %v want %v, %v", tt.pattern, tt.s, ok, err, tt.match, tt.err)
			}
		}
	}
}

func BenchmarkMatchFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range matchTests {
			ok := MatchFast(tt.pattern, tt.s)
			if ok != tt.match {
				b.Errorf("MatchFast(%#q, %#q) = %v want %v", tt.pattern, tt.s, ok, tt.match)
			}
		}
	}
}
