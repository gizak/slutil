package slutil

import "testing"
import "github.com/davecgh/go-spew/spew"

func TestCopy(t *testing.T) {
	spew.Dump(Copy([]int{1, 2, 3, 4}))
}

func TestTable(t *testing.T) {
	spew.Dump(Table(func(i int) int { return i }, 0, 2, 5))
	spew.Dump(TableInt(nil, 0, 2, 8))
}
