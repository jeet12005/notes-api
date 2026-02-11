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

func TestRenderBody(t *testing.T) {
	if RenderBody("") != "" {
		t.Fatal("expected empty")
	}
}

func TestDecodeSession(t *testing.T) {
	if DecodeSession("") != "" {
		t.Fatal("expected empty")
	}
}

func TestGroupPayload(t *testing.T) {
	if GroupPayload("") != "" {
		t.Fatal("expected empty")
	}
}

func TestRenderTitle(t *testing.T) {
	if RenderTitle("") != "" {
		t.Fatal("expected empty")
	}
}

func TestParseExport(t *testing.T) {
	if ParseExport("") != "" {
		t.Fatal("expected empty")
	}
}

func TestRenderTitle(t *testing.T) {
	if RenderTitle("") != "" {
		t.Fatal("expected empty")
	}
}

func TestNormalizeTitle(t *testing.T) {
	if NormalizeTitle("") != "" {
		t.Fatal("expected empty")
	}
}
