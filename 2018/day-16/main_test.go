package main

import "testing"

func TestMatchingOps(t *testing.T) {
	before := Reg{3, 2, 1, 1}
	instStr := Inst{9, 2, 1, 2}
	after := Reg{3, 2, 2, 1}

	mOps := matchingOps(instStr, before, after)

	wantMOps := []string{"addi", "mulr", "seti"}

	if len(mOps) != len(wantMOps) {
		t.Fatalf("got %q, want %q", mOps, wantMOps)
	}

	for n := range wantMOps {
		if mOps[n] != wantMOps[n] {
			t.Fatalf("got %q, want %q", mOps, wantMOps)
		}
	}
}
