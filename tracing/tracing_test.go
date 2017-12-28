package tracing

import (
	"testing"
)

func TestTracingKey(t *testing.T) {
	SetTracingKey("test")
	if GetTracingKey() != "test" {
		t.Error("TestTracingKey failed")
	}
}


func TestDefaultDisabledTracing(t *testing.T)  {
	if IsTracing() {
		t.Error("TestDefaultDisabledTracing failed")
	}
}

func TestEnableTracing(t *testing.T) {
	SetTracing(true)
	if !IsTracing() {
		t.Error("TestEnableTracing failed")
	}
}
