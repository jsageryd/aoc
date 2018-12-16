package main

import "testing"

func TestScoresAfterNth(t *testing.T) {
	for n, tc := range []struct {
		n      int
		count  int
		scores []int
	}{
		{9, 10, []int{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}},
		{5, 10, []int{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}},
		{18, 10, []int{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}},
		{2018, 10, []int{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}},
	} {
		scores := scoresAfterNth(tc.n, tc.count)

		if len(scores) != len(tc.scores) {
			t.Fatalf("[%d] got %v, want %v", n, scores, tc.scores)
		}

		for n := range tc.scores {
			if scores[n] != tc.scores[n] {
				t.Fatalf("[%d] got %v, want %v", n, scores, tc.scores)
			}
		}
	}
}

func TestMakeRecipes(t *testing.T) {
	wantRecipes := []int{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7, 9, 2}

	recipes := makeRecipes(20)

	if len(recipes) != len(wantRecipes) {
		t.Fatalf("got %v, want %v", recipes, wantRecipes)
	}

	for n := range wantRecipes {
		if recipes[n] != wantRecipes[n] {
			t.Fatalf("got %v, want %v", recipes, wantRecipes)
		}
	}
}
