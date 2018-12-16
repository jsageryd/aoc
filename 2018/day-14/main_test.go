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

func TestNumberOfRecipesBeforeSequence(t *testing.T) {
	for n, tc := range []struct {
		seq []int
		n   int
	}{
		{[]int{5, 1, 5, 8, 9}, 9},
		{[]int{0, 1, 2, 4, 5}, 5},
		{[]int{9, 2, 5, 1, 0}, 18},
		{[]int{5, 9, 4, 1, 4}, 2018},
	} {
		if got, want := numberOfRecipesBeforeSequence(tc.seq), tc.n; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestMakeRecipes(t *testing.T) {
	wantRecipes := []int{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7, 9, 2}
	var recipes []int

	c := makeRecipes()
	for n := 0; n < 20; n++ {
		recipes = append(recipes, <-c)
	}

	if len(recipes) != len(wantRecipes) {
		t.Fatalf("got %v, want %v", recipes, wantRecipes)
	}

	for n := range wantRecipes {
		if recipes[n] != wantRecipes[n] {
			t.Fatalf("got %v, want %v", recipes, wantRecipes)
		}
	}
}
