package test

import (
	"fmt"
	"testing"
)

func TestDemo01(t *testing.T) {
	a := &map[string]any{
		"a": 1,
	}
	var ok bool
	if i, ok := (*a)["a"]; ok {
		fmt.Println(i, ok)
	}
	fmt.Println(ok)
}
