package service

import (
	"testing"
)

func TestFindAll(t *testing.T) {
	movie := FindAll()

	if len(movie) < 1 {
		t.Fail()
	}
}
