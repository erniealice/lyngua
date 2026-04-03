package v1

import (
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// extractKBTitle
// ---------------------------------------------------------------------------

func TestExtractKBTitle_FindsH1(t *testing.T) {
	src := "# My Title\n\nSome body text.\n"
	got := extractKBTitle(src)
	if got != "My Title" {
		t.Errorf("expected 'My Title', got %q", got)
	}
}

func TestExtractKBTitle_SkipsH2(t *testing.T) {
	src := "## Not a title\n\n# Real Title\n"
	got := extractKBTitle(src)
	if got != "Real Title" {
		t.Errorf("expected 'Real Title', got %q", got)
	}
}

func TestExtractKBTitle_EmptyWhenNoH1(t *testing.T) {
	src := "## Only H2\n\nBody text.\n"
	got := extractKBTitle(src)
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestExtractKBTitle_EmptyInput(t *testing.T) {
	got := extractKBTitle("")
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestExtractKBTitle_TrimsWhitespace(t *testing.T) {
	src := "  # Trimmed Title  \n\nBody.\n"
	got := extractKBTitle(src)
	if got != "Trimmed Title" {
		t.Errorf("expected 'Trimmed Title', got %q", got)
	}
}

// ---------------------------------------------------------------------------
// extractKBSummary
// ---------------------------------------------------------------------------

func TestExtractKBSummary_FirstParagraph(t *testing.T) {
	src := "# Title\n\nThis is the summary paragraph.\n\nAnother paragraph.\n"
	got := extractKBSummary(src)
	if got != "This is the summary paragraph." {
		t.Errorf("expected 'This is the summary paragraph.', got %q", got)
	}
}

func TestExtractKBSummary_SkipsBlankLines(t *testing.T) {
	src := "# Title\n\n\n\nActual summary.\n"
	got := extractKBSummary(src)
	if got != "Actual summary." {
		t.Errorf("expected 'Actual summary.', got %q", got)
	}
}

func TestExtractKBSummary_TruncatesAt200(t *testing.T) {
	longLine := strings.Repeat("x", 250)
	src := "# Title\n\n" + longLine + "\n"
	got := extractKBSummary(src)

	// Should be 200 chars + ellipsis
	if len(got) != 201 { // 200 bytes + 3-byte ellipsis "…"
		// The ellipsis "…" is a multi-byte UTF-8 character (3 bytes)
		expectedLen := 200 + len("…")
		if len(got) != expectedLen {
			t.Errorf("expected truncated length %d, got %d", expectedLen, len(got))
		}
	}
	if !strings.HasSuffix(got, "…") {
		t.Error("expected truncated summary to end with ellipsis")
	}
	if !strings.HasPrefix(got, "xxxx") {
		t.Error("expected summary to start with 'xxxx'")
	}
}

func TestExtractKBSummary_NoTruncationUnder200(t *testing.T) {
	src := "# Title\n\nShort summary.\n"
	got := extractKBSummary(src)
	if got != "Short summary." {
		t.Errorf("expected 'Short summary.', got %q", got)
	}
	if strings.HasSuffix(got, "…") {
		t.Error("short summary should not have ellipsis")
	}
}

func TestExtractKBSummary_EmptyWhenNoContent(t *testing.T) {
	src := "# Title\n"
	got := extractKBSummary(src)
	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestExtractKBSummary_SkipsHeadings(t *testing.T) {
	src := "# Title\n## Subtitle\nFirst real paragraph.\n"
	got := extractKBSummary(src)
	if got != "First real paragraph." {
		t.Errorf("expected 'First real paragraph.', got %q", got)
	}
}

// ---------------------------------------------------------------------------
// loadKBFile (integration via LoadKB)
// ---------------------------------------------------------------------------

func TestLoadKBFile_RendersMarkdownToHTML(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/kb/test-topic.md": "# Test Title\n\nA **bold** paragraph.\n",
	})

	kb, err := p.LoadKB("en", "common", "test-topic")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if kb.Topic != "test-topic" {
		t.Errorf("expected topic 'test-topic', got %q", kb.Topic)
	}
	if kb.Title != "Test Title" {
		t.Errorf("expected title 'Test Title', got %q", kb.Title)
	}
	if kb.Summary != "A **bold** paragraph." {
		t.Errorf("expected summary 'A **bold** paragraph.', got %q", kb.Summary)
	}
	body := string(kb.Body)
	if !strings.Contains(body, "<strong>bold</strong>") {
		t.Errorf("expected rendered HTML with <strong>, got: %s", body)
	}
}

func TestLoadKBFile_RendersHTMLHeading(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/kb/heading.md": "# My Heading\n\nContent here.\n",
	})

	kb, err := p.LoadKB("en", "common", "heading")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	body := string(kb.Body)
	if !strings.Contains(body, "<h1") {
		t.Errorf("expected <h1> in rendered body, got: %s", body)
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive extractKBTitle tests
// ---------------------------------------------------------------------------

func TestExtractKBTitle_OnlyH2AndH3_NoH1(t *testing.T) {
	src := "## Section\n### Subsection\n\nBody text.\n"
	got := extractKBTitle(src)
	if got != "" {
		t.Errorf("expected empty string when no H1 present, got %q", got)
	}
}

func TestExtractKBTitle_MultipleH1_ReturnsFirst(t *testing.T) {
	src := "# First Title\n\n# Second Title\n"
	got := extractKBTitle(src)
	if got != "First Title" {
		t.Errorf("expected 'First Title', got %q", got)
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive extractKBSummary tests
// ---------------------------------------------------------------------------

func TestExtractKBSummary_OnlyHeadings_NoParagraphs(t *testing.T) {
	src := "# Title\n## Section A\n## Section B\n### Subsection\n"
	got := extractKBSummary(src)
	if got != "" {
		t.Errorf("expected empty string when only headings present, got %q", got)
	}
}

func TestExtractKBSummary_EmptyString(t *testing.T) {
	got := extractKBSummary("")
	if got != "" {
		t.Errorf("expected empty string for empty input, got %q", got)
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive loadKBFile / RenderMarkdown tests
// ---------------------------------------------------------------------------

func TestLoadKBFile_MaliciousHTMLInjection(t *testing.T) {
	// Goldmark with html.WithUnsafe() will pass through raw HTML.
	// Verify the renderer does not panic and produces output.
	malicious := "# Title\n\n<script>alert('xss')</script>\n\n<img src=x onerror=alert(1)>\n\nNormal paragraph.\n"
	p := testProvider(map[string]string{
		"en/common/kb/xss-test.md": malicious,
	})

	kb, err := p.LoadKB("en", "common", "xss-test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if kb.Title != "Title" {
		t.Errorf("expected title 'Title', got %q", kb.Title)
	}

	body := string(kb.Body)
	// With WithUnsafe(), the script tag passes through. Verify it is present
	// (documenting current behavior — a sanitization layer would strip it).
	if !strings.Contains(body, "<script>") {
		t.Log("script tag was sanitized — good, but unexpected with WithUnsafe()")
	}

	// The function should always return a non-empty body.
	if len(body) == 0 {
		t.Error("expected non-empty rendered body")
	}
}

func TestLoadKBFile_OnlyH2_EmptyTitle(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/kb/no-h1.md": "## Only H2\n\nSome body text.\n",
	})

	kb, err := p.LoadKB("en", "common", "no-h1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if kb.Title != "" {
		t.Errorf("expected empty title when no H1, got %q", kb.Title)
	}
	if kb.Summary != "Some body text." {
		t.Errorf("expected summary 'Some body text.', got %q", kb.Summary)
	}
}
