package main

import (
	"testing"
)

func Test_Unquote(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out string
	}{
		{`""`, ``},
		{`"abc"`, `abc`},
		{`"aaa\"aaa"`, `aaa"aaa`},
		{`"\x27"`, `'`},
	} {
		if got, want := unquote(tc.in), tc.out; got != want {
			t.Errorf("[%d] got `%s`, want `%s`", n, got, want)
		}
	}
}

func Test_Quote(t *testing.T) {
	for n, tc := range []struct {
		in  string
		out string
	}{
		{`""`, `"\"\""`},
		{`"abc"`, `"\"abc\""`},
		{`"aaa\"aaa"`, `"\"aaa\\\"aaa\""`},
		{`"\x27"`, `"\"\\x27\""`},
	} {
		if got, want := quote(tc.in), tc.out; got != want {
			t.Errorf("[%d] got `%s`, want `%s`", n, got, want)
		}
	}
}
