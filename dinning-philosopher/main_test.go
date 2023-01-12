package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()

		if len(orderFinished) != 5 {
			t.Errorf("Incorrect slice length expected 5 not %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"Zero delay", time.Second * 0},
		{"Quater second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}
	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: Incorrect slice length expected 5 not %d", e.name, len(orderFinished))
		}
	}
}
