package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSortV1(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6, 0}

	sorter := newInsertSorterV1()
	sorter.sortImple(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestInsertSortV2(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6, 0}

	sorter := newInsertSorterV2()
	sorter.sortImple(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestInsertSort3(t *testing.T) {
	test := assert.New(t)
	ints := []int{5, 2, 1, 8, 3, 4, 9, 7, 6, 0}

	sorter := newInsertSorterV3()
	sorter.sortImple(ints)
	test.Equal(ints, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}
