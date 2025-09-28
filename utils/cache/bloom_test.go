package cache

import (
	"fmt"
	"testing"
)

func TestBloomFilter_Allow(t *testing.T) {
	s := newFilter(1000, 0.01)
	k1 := uint32(3330196039)
	k2 := uint32(3695537765)
	k3 := uint32(1257462242)
	s.Allow(k1)
	s.Allow(k2)
	e1 := s.MayContain(k1)
	e2 := s.MayContain(k2)
	e3 := s.MayContain(k3)
	fmt.Printf("%v:%v \n", len(s.bitmap), s)
	fmt.Println(e1, e2, e3)
}
