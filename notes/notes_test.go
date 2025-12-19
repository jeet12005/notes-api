package notes

import "testing"

func TestResolveIndex(t *testing.T) {
	if ResolveIndex("") != "" {
		t.Fatal("expected empty")
	}
}

func TestRenderBody(t *testing.T) {
	if RenderBody("") != "" {
		t.Fatal("expected empty")
	}
}

func TestSplitNote2(t *testing.T) {
	if SplitNote2("") != "" {
		t.Fatal("expected empty")
	}
}

func TestSplitNote2(t *testing.T) {
	if SplitNote2("") != "" {
		t.Fatal("expected empty")
	}
}

func TestRenderBody(t *testing.T) {
	if RenderBody("") != "" {
		t.Fatal("expected empty")
	}
}

func TestDecodeQuery(t *testing.T) {
	if DecodeQuery("") != "" {
		t.Fatal("expected empty")
	}
}
