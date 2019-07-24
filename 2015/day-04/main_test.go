package main

import "testing"

func TestMD5HexHasFiveLeadingZeros(t *testing.T) {
	for n, tc := range []struct {
		b    []byte
		want bool
	}{
		{[]byte("foo"), false},
		{[]byte("abcdef609043"), true},
		{[]byte("pqrstuv1048970"), true},
	} {
		if got, want := md5HexHasFiveLeadingZeros(tc.b), tc.want; got != want {
			t.Errorf("[%d] md5HexHasFiveLeadingZeros(%q) = %t, want %t", n, tc.b, got, want)
		}
	}
}
