package main

import "testing"

func TestCPU(t *testing.T) {
	t.Run("Program from description", func(t *testing.T) {
		program := []string{
			"inc a",
			"jio a, +2",
			"tpl a",
			"inc a",
		}

		c := newCPU(program)

		c.run()

		if got, want := c.reg["a"], 2; got != want {
			t.Errorf(`c.reg["a"] = %d, want %d`, got, want)
		}
	})
}
