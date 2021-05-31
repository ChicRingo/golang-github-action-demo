package main

import "testing"

// 猫在叫
func TestCat(t *testing.T) {
	saying := Cat()
	if saying != "miao~~~" {
		t.Errorf("Cat() = %s", saying)
	}
}
