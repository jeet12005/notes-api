package notes

import "testing"

func TestResolveIndex(t *testing.T) {
	if ResolveIndex("") != "" {
		t.Fatal("expected empty")
	}
}
