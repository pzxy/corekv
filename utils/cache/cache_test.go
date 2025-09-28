package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCacheBasicCRUD(t *testing.T) {
	cache := NewCache(5)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val := fmt.Sprintf("val%d", i)
		cache.Set(key, val)
		fmt.Printf("->set %s: %s\n", key, cache)
	}

	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key%d", i)
		val := fmt.Sprintf("val%d", i)
		res, ok := cache.Get(key)
		if ok {
			fmt.Printf("->get %s: %s\n", key, cache)
			assert.Equal(t, val, res)
			continue
		}
		assert.Equal(t, res, nil)
	}
	fmt.Printf("at last: %s\n", cache)
}

func TestHash(t *testing.T) {
	v := "123456789"
	x := math.Pow(0.69314718056, 2)
	println(x)
	println(0.69314718056)
	println(Hash([]byte(v)))
}

func TestCmSketch_Clear(t *testing.T) {
	c := int32(1) >> 128
	println(c)
	a := newCmSketch(100)
	a.Increment(1234567890)
	a.Increment(1234567890)
	a.Increment(1234567890)
	b := a.Estimate(1234567890)
	println(b)
}
