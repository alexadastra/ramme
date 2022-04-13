package system

import "testing"

func TestStubHandling(t *testing.T) {
	handling := NewGroupOperator()
	err := handling.Reload()
	if err != ErrNotImplemented {
		t.Error("Expected error", ErrNotImplemented, "got", err)
	}
	err = handling.Maintenance()
	if err != ErrNotImplemented {
		t.Error("Expected error", ErrNotImplemented, "got", err)
	}
	err = handling.Shutdown(nil)
	if err != nil {
		t.Error("Expected error", nil, "got", err)
	}
}
