package cache

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	Put("1111", "hhhh", 1000)
	v, ok := Get("1111")
	if ok {
		fmt.Println(v)
	}

}
