package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

var rules = []string{
	"bright white bags contain 1 shiny gold bag.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"dotted black bags contain no other bags.",
	"faded blue bags contain no other bags.",
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
}

func TestPart1(t *testing.T) {
	if got, want := part1(rules), 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseRules(t *testing.T) {
	gotMap := parseRules(rules)

	wantMap := map[string][]string{
		"bright white": {"shiny gold"},
		"dark olive":   {"faded blue", "faded blue", "faded blue", "dotted black", "dotted black", "dotted black", "dotted black"},
		"dark orange":  {"bright white", "bright white", "bright white", "muted yellow", "muted yellow", "muted yellow", "muted yellow"},
		"light red":    {"bright white", "muted yellow", "muted yellow"},
		"muted yellow": {"shiny gold", "shiny gold", "faded blue", "faded blue", "faded blue", "faded blue", "faded blue", "faded blue", "faded blue", "faded blue", "faded blue"},
		"shiny gold":   {"dark olive", "vibrant plum", "vibrant plum"},
		"vibrant plum": {"faded blue", "faded blue", "faded blue", "faded blue", "faded blue", "dotted black", "dotted black", "dotted black", "dotted black", "dotted black", "dotted black"},
	}

	gotMapJSON, _ := json.MarshalIndent(gotMap, "", "  ")
	wantMapJSON, _ := json.MarshalIndent(wantMap, "", "  ")

	if !bytes.Equal(gotMapJSON, wantMapJSON) {
		t.Errorf("got %s, want %s", gotMapJSON, wantMapJSON)
	}
}

func TestParseRule(t *testing.T) {
	for n, tc := range []struct {
		rule          string
		wantParentBag string
		wantChildBags []string
	}{
		{
			rule:          "faded blue bags contain no other bags.",
			wantParentBag: "faded blue",
			wantChildBags: []string{},
		},
		{
			rule:          "bright white bags contain 1 shiny gold bag.",
			wantParentBag: "bright white",
			wantChildBags: []string{
				"shiny gold",
			},
		},
		{
			rule:          "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			wantParentBag: "light red",
			wantChildBags: []string{
				"bright white",
				"muted yellow", "muted yellow",
			},
		},
		{
			rule:          "dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			wantParentBag: "dark orange",
			wantChildBags: []string{
				"bright white", "bright white", "bright white",
				"muted yellow", "muted yellow", "muted yellow", "muted yellow",
			},
		},
		{
			rule:          "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			wantParentBag: "vibrant plum",
			wantChildBags: []string{
				"faded blue", "faded blue", "faded blue", "faded blue", "faded blue",
				"dotted black", "dotted black", "dotted black", "dotted black", "dotted black", "dotted black",
			},
		},
		{
			rule:          "plaid indigo bags contain 3 mirrored fuchsia bags, 5 wavy yellow bags, 2 plaid silver bags, 3 vibrant blue bags.",
			wantParentBag: "plaid indigo",
			wantChildBags: []string{
				"mirrored fuchsia", "mirrored fuchsia", "mirrored fuchsia",
				"wavy yellow", "wavy yellow", "wavy yellow", "wavy yellow", "wavy yellow",
				"plaid silver", "plaid silver",
				"vibrant blue", "vibrant blue", "vibrant blue",
			},
		},
	} {
		gotParentBag, gotChildBags := parseRule(tc.rule)

		if gotParentBag != tc.wantParentBag {
			t.Errorf("[%d] parent bag is %q, want %q", n, gotParentBag, tc.wantParentBag)
		}

		if fmt.Sprint(gotChildBags) != fmt.Sprint(tc.wantChildBags) {
			t.Errorf("[%d] child bags are %q, want %q", n, gotChildBags, tc.wantChildBags)
		}
	}
}
