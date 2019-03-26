package cluster

import (
	"errors"
	"testing"
)

func TestValueFound(t *testing.T) {
	m := NewErrorMap()
	m.Store("cluster1", ErrorValue{
		CmdError:  errors.New("error from cmd"),
		CmdOutput: "There was an error",
	})
	value, found := m.Load("cluster1")
	if found != true {
		t.Error("expected found = true, got: false")
	}
	t.Logf("value.error = %v  value.output = %s", value.CmdError, value.CmdOutput)
}

func TestValueNotFound(t *testing.T) {
	m := NewErrorMap()
	m.Store("cluster1", ErrorValue{
		CmdError:  errors.New("error from cmd"),
		CmdOutput: "There was an error",
	})
	value, found := m.Load("cluster2")
	if found == true {
		t.Error("expected found = false, got: true")
	}
	t.Logf("value = %s", value)
}

func TestDelete(t *testing.T) {
	m := NewErrorMap()
	m.Store("cluster1", ErrorValue{
		CmdError:  errors.New("error from cmd"),
		CmdOutput: "There was an error",
	})
	m.Delete("cluster1")
	value, found := m.Load("cluster1")
	if found == true {
		t.Error("expected found = false, got: true")
	}
	t.Logf("value = %s", value)
}
