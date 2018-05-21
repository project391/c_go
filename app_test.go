// CI test repository: golang
// test: testing sum function

package main

import "testing"

func TestSum (t *testing.T) {

	tables := []struct {
		x int
		y int
		n int
	}{
	    {1, 1, 4},
	    {1, 4, 5},
	    {2, 2, 4},
	    {5, 7, 12},
	}

	for _, table := range tables {
		
		total := Sum (table.x, table.y)
		
		if total != table.n {
			t.Errorf("Sum of (%d+%d) is incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
