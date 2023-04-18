package demo1_test

import (
	"testing"

	"com.github.manishktiwari/go-demo1/demo1"
)

func TestBestUser(t *testing.T) {
	if demo1.BestUSer() != "manishktiwari" {
		t.Fatal("test failed")
	}
}
