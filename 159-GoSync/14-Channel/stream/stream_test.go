package stream

import (
	"testing"
)

func TestAsStream(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5)

	for v := range in {
		t.Log(v)
	}
}

func TestTakeN(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := takeN(nil, in, 3)

	for v := range result {
		t.Log(v)
	}
}

func TestTakeFn(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := takeFn(nil, in, func(item any) bool {
		return item.(int) > 5
	})

	for v := range result {
		t.Log(v)
	}
}

func TestTakeWhile(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := takeWhile(nil, in, func(item any) bool {
		return item.(int) < 5
	})

	for v := range result {
		t.Log(v)
	}
}

func TestSkipN(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := skipN(nil, in, 6)
	for v := range result {
		t.Log(v)
	}
}

func TestSkipFn(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := skipFn(nil, in, func(item any) bool {
		return item.(int) > 2
	})

	for v := range result {
		t.Log(v)
	}
}

func TestSkipWhile(t *testing.T) {
	in := asStream(nil, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := skipWhile(nil, in, func(item any) bool {
		return item.(int)%2 == 1
	})

	for v := range result {
		t.Log(v)
	}
}
