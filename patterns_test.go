package luapatterns

import (
	"fmt"
	"testing"
)

type posTest struct {
	str string
	pat string
	init int
	succ bool
	start int
	end int
}

// These tests have been taken from the Lua test suite, but have been altered
// to reflect the difference in array indexing and substrings (i.e) str[1:1]
// is always an empty string. As a result the start and end returns must be
// adjusted.
var posTests = []posTest{
	posTest{"", "", 0, true, 0, 0},									// special case
	posTest{"alo", "", 0, true, 0, 0},								// special case
	posTest{"a\x00o a\x00o a\x00o", "a", 0, true, 0, 1},
	posTest{"a\x00o a\x00o a\x00o", "a\x00o", 2, true, 4, 7},
	posTest{"a\x00o a\x00o a\x00o", "a\x00o", 8, true, 8, 11},
	posTest{"a\x00oa\x00a\x00a\x00\x00ab", "\x00ab", 1, true, 9, 12},
	posTest{"a\x00oa\x00a\x00a\x00\x00ab", "b", 0, true, 11, 12},
	posTest{"a\x00oa\x00a\x00a\x00\x00ab", "b\x00", 0, false, 0, 0},
	posTest{"", "\x00", 0, false, 0, 0},
	posTest{"alo123alo", "12", 0, true, 3, 5},
	posTest{"alo123alo", "^12", 0, false, 0, 0},
}

func _TestPatternPos(t *testing.T) {
	for _, test := range posTests {
		succ, start, end, _ := FindString(test.str, test.pat, test.init, false)
		if succ != test.succ {
			t.Errorf("find('%s', '%s', %d) returned %t, expected %t", test.str, test.pat, test.init, succ, test.succ)
		}
		if succ && (start != test.start || end != test.end) {
			t.Errorf("find('%s', '%s', %d) returned (%d, %d), expected (%d, %d)", test.str, test.pat, test.init, start, end, test.start, test.end)
		}
	}
}

type subTest struct {
	str string
	pat string
	succ bool
	cap string
}

var subTests = []subTest{
	subTest{"aaab", "a*", true, "aaa"},
	subTest{"aaa", "^.*$", true, "aaa"},
	subTest{"aaa", "b*", true, ""},
	subTest{"aaa", "ab*a", true, "aa"},
	subTest{"aba", "ab*a", true, "aba"},
	subTest{"aaab", "a+", true, "aaa"},
	subTest{"aaa", "^.+$", true, "aaa"},
	subTest{"aaa", "b+", false, ""},
	subTest{"aaa", "ab+a", false, ""},
	subTest{"aba", "ab+a", true, "aba"},
	subTest{"a$a", ".$", true, "a"},
	subTest{"a$a", ".%$", true, "a$"},
	subTest{"a$a", ".$.", true, "a$a"},
	subTest{"a$a", "$$", false, ""},
	subTest{"a$b", "a$", false, ""},
	subTest{"a$a", "$", true ,""},
	subTest{"", "b*", true, ""},
	subTest{"aaa", "bb*", false, ""},
	subTest{"aaab", "a-", true, ""},
	subTest{"aaa", "^.-$", true, "aaa"},
	subTest{"aabaaabaaabaaaba", "b.*b", true, "baaabaaabaaab"},
	subTest{"aabaaabaaabaaaba", "b.-b", true, "baaab"},
	subTest{"alo xo", ".o$", true, "xo"},
	subTest{" \n isto é assim", "%S%S*", true, "isto"},
	subTest{" \n isto é assim", "%S*$", true, "assim"},
	subTest{" \n isto é assim", "[a-z]*$", true, "assim"},
	subTest{"im caracter ? extra", "[^%sa-z]", true, "?"},
	subTest{"", "a?", true, ""},

	// These tests don't work 100% correctly if you use Unicode á instead
	// of the ASCII value \225. In particular the third test fails
	subTest{"\225", "\225?", true, "\225"},
	subTest{"\225bl", "\225?b?l?", true, "\225bl"},
	subTest{"  \225bl", "\225?b?l?", true, ""},

	subTest{"aa", "^aa?a?a", true, "aa"},
	subTest{"]]]\225b", "[^]]", true, "\225"},
	subTest{"0alo alo", "%x*", true, "0a"},
	subTest{"alo alo", "%C+", true, "alo alo"},
}

func TestSubtring(t *testing.T) {
	enableDebug = false

	for _, test := range subTests {

		debug(fmt.Sprintf("=== %s ===", test))
		succ, start, end, caps := FindString(test.str, test.pat, 0, false)
		if succ != test.succ {
			t.Errorf("find('%s', '%s') returned %t, expected %t", test.str, test.pat, succ, test.succ)
		}
		if succ {
			substr := test.str[start:end]
			if substr != test.cap {
				t.Errorf("find('%s', '%s') => substr '%s' does not match expected '%s'", test.str, test.pat, substr, test.cap)
			}
			if len(caps) != 1 || substr != caps[0] {
				t.Errorf("find('%s', '%s') => substr '%s' does not match capture '%s'", test.str, test.pat, substr, caps[0])
			}
		}
	}
}