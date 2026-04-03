package v1

import (
	"fmt"
	"testing"
	"testing/fstest"
)

// ---------------------------------------------------------------------------
// flattenAndMerge
// ---------------------------------------------------------------------------

func TestFlattenAndMerge_SimpleKeys(t *testing.T) {
	src := map[string]any{
		"greeting": "Hello",
		"farewell": "Goodbye",
	}
	target := make(map[string]string)
	flattenAndMerge(src, "", target)

	if target["greeting"] != "Hello" {
		t.Errorf("expected greeting='Hello', got %q", target["greeting"])
	}
	if target["farewell"] != "Goodbye" {
		t.Errorf("expected farewell='Goodbye', got %q", target["farewell"])
	}
}

func TestFlattenAndMerge_NestedKeys(t *testing.T) {
	src := map[string]any{
		"client": map[string]any{
			"errors": map[string]any{
				"not_found": "Not found",
				"required":  "Required",
			},
			"name": "Client Name",
		},
	}
	target := make(map[string]string)
	flattenAndMerge(src, "", target)

	tests := map[string]string{
		"client.errors.not_found": "Not found",
		"client.errors.required":  "Required",
		"client.name":             "Client Name",
	}
	for key, want := range tests {
		if got := target[key]; got != want {
			t.Errorf("key %q: expected %q, got %q", key, want, got)
		}
	}
}

func TestFlattenAndMerge_WithPrefix(t *testing.T) {
	src := map[string]any{"key": "value"}
	target := make(map[string]string)
	flattenAndMerge(src, "prefix", target)

	if target["prefix.key"] != "value" {
		t.Errorf("expected prefix.key='value', got %q", target["prefix.key"])
	}
}

func TestFlattenAndMerge_NonStringValue(t *testing.T) {
	src := map[string]any{
		"count": float64(42),
		"flag":  true,
	}
	target := make(map[string]string)
	flattenAndMerge(src, "", target)

	if target["count"] != "42" {
		t.Errorf("expected count='42', got %q", target["count"])
	}
	if target["flag"] != "true" {
		t.Errorf("expected flag='true', got %q", target["flag"])
	}
}

func TestFlattenAndMerge_OverwritesExisting(t *testing.T) {
	target := map[string]string{"a": "old"}
	src := map[string]any{"a": "new"}
	flattenAndMerge(src, "", target)

	if target["a"] != "new" {
		t.Errorf("expected a='new', got %q", target["a"])
	}
}

// ---------------------------------------------------------------------------
// LoadMessages — 3-tier cascade with flat dot-notation keys
// ---------------------------------------------------------------------------

func testProviderInternal(files map[string]string) *TranslationProvider {
	fsMap := make(fstest.MapFS)
	for p, content := range files {
		fsMap["translations/"+p] = &fstest.MapFile{Data: []byte(content)}
	}
	return &TranslationProvider{
		translationsPath: "translations",
		fsys:             fsMap,
		cache:            make(map[string]map[string]string),
	}
}

func TestLoadMessages_ThreeTierCascade(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/labels.json":  `{"greeting":"Hello","farewell":"Goodbye"}`,
		"en/general/labels.json": `{"greeting":"Hi there","extra":"General only"}`,
		"en/retail/labels.json":  `{"farewell":"See ya","retail_only":"Yes"}`,
	})

	msgs, err := p.LoadMessages("en", "retail")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := map[string]string{
		"greeting":    "Hi there", // general overrides common
		"farewell":    "See ya",   // retail overrides common
		"extra":       "General only",
		"retail_only": "Yes",
	}
	for key, want := range tests {
		if got := msgs[key]; got != want {
			t.Errorf("key %q: expected %q, got %q", key, want, got)
		}
	}
}

func TestLoadMessages_NestedJSONFlattened(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/client.json": `{"client":{"errors":{"not_found":"Not found"}}}`,
	})

	msgs, err := p.LoadMessages("en", "common")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "Not found"
	if got := msgs["client.errors.not_found"]; got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestLoadMessages_CacheHit(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/labels.json": `{"key":"value"}`,
	})

	msgs1, err := p.LoadMessages("en", "common")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	msgs2, err := p.LoadMessages("en", "common")
	if err != nil {
		t.Fatalf("unexpected error on cached call: %v", err)
	}

	// Both should return the same cached map
	if msgs1["key"] != msgs2["key"] {
		t.Errorf("expected cached result to match, got %q vs %q", msgs1["key"], msgs2["key"])
	}
}

func TestLoadMessages_MultipleFilesInDirectory(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/labels.json": `{"label":"Label Value"}`,
		"en/common/errors.json": `{"error":"Error Value"}`,
	})

	msgs, err := p.LoadMessages("en", "common")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if msgs["label"] != "Label Value" {
		t.Errorf("expected label='Label Value', got %q", msgs["label"])
	}
	if msgs["error"] != "Error Value" {
		t.Errorf("expected error='Error Value', got %q", msgs["error"])
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive LoadMessages tests
// ---------------------------------------------------------------------------

func TestLoadMessages_MalformedJSON_ReturnsError(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/bad.json": `{"key": BROKEN}`,
	})

	_, err := p.LoadMessages("en", "common")
	if err == nil {
		t.Fatal("expected error for malformed JSON file, got nil")
	}
}

func TestLoadMessages_EmptyLanguageCode(t *testing.T) {
	p := testProviderInternal(map[string]string{
		"en/common/labels.json": `{"key":"value"}`,
	})

	// Empty locale string should fail because no directory exists for "".
	_, err := p.LoadMessages("", "common")
	if err == nil {
		t.Fatal("expected error for empty language code, got nil")
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive flattenAndMerge tests
// ---------------------------------------------------------------------------

func TestFlattenAndMerge_DeeplyNested10Levels(t *testing.T) {
	// Build 10 levels of nesting: level1.level2. ... .level10 = "deep"
	innermost := map[string]any{"level10": "deep"}
	current := innermost
	for i := 9; i >= 1; i-- {
		current = map[string]any{
			fmt.Sprintf("level%d", i): current,
		}
	}

	target := make(map[string]string)
	flattenAndMerge(current, "", target)

	wantKey := "level1.level2.level3.level4.level5.level6.level7.level8.level9.level10"
	if target[wantKey] != "deep" {
		t.Errorf("expected deeply nested key %q = 'deep', got %q", wantKey, target[wantKey])
	}
}

func TestFlattenAndMerge_EmptySource(t *testing.T) {
	target := map[string]string{"existing": "value"}
	flattenAndMerge(map[string]any{}, "", target)

	if target["existing"] != "value" {
		t.Errorf("expected existing key preserved, got %q", target["existing"])
	}
	if len(target) != 1 {
		t.Errorf("expected 1 key in target, got %d", len(target))
	}
}
