package apm

import "testing"

func TestNewPwrInfo(t *testing.T) {
	_, err := NewPwrInfo()
	if err != nil {
		t.Error(err)
	}
}
