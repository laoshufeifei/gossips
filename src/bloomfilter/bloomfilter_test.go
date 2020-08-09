package bloomfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloomFilterBasic(t *testing.T) {
	test := assert.New(t)

	b := New(100, 0.01)
	test.False(b.Contains("abc"))

	b.Put("abc")
	test.True(b.Contains("abc"))

	test.False(b.Contains("abcd"))
}

func TestBloomFilter(t *testing.T) {
	test := assert.New(t)

	allCont := uint64(100)
	b := New(allCont, 0.001)

	putCount := uint64(100)
	for i := uint64(0); i < putCount; i++ {
		s := fmt.Sprintf("%d", i)
		b.Put(s)
	}

	for i := uint64(0); i < putCount; i++ {
		s := fmt.Sprintf("%d", i)
		test.True(b.Contains(s))
	}

	for i := uint64(0); i < putCount; i++ {
		s := fmt.Sprintf("%d", i)
		test.True(b.Contains(s))
	}

	errCount := uint64(0)
	for i := putCount; i < putCount*2; i++ {
		s := fmt.Sprintf("%d", i)
		if b.Contains(s) {
			errCount++
		}
	}

	if errCount > 0 {
		fmt.Println("error count:", errCount)
	}
	test.True(float64(errCount/allCont) <= 0.001)
}
