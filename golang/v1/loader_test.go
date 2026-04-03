package v1

import (
	"testing"
	"testing/fstest"
)

// testProvider creates a TranslationProvider backed by an in-memory fstest.MapFS.
// File paths should be relative to "translations/" (e.g. "en/common/file.json").
func testProvider(files map[string]string) *TranslationProvider {
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

// ---------------------------------------------------------------------------
// mergeMap
// ---------------------------------------------------------------------------

func TestMergeMap_SrcWinsScalar(t *testing.T) {
	dst := map[string]any{"a": "old"}
	src := map[string]any{"a": "new"}
	mergeMap(dst, src)
	if dst["a"] != "new" {
		t.Errorf("expected src to win, got %v", dst["a"])
	}
}

func TestMergeMap_AddsNewKeys(t *testing.T) {
	dst := map[string]any{"a": "1"}
	src := map[string]any{"b": "2"}
	mergeMap(dst, src)
	if dst["b"] != "2" {
		t.Errorf("expected key b=2, got %v", dst["b"])
	}
	if dst["a"] != "1" {
		t.Errorf("expected key a=1 preserved, got %v", dst["a"])
	}
}

func TestMergeMap_RecursiveDeepMerge(t *testing.T) {
	dst := map[string]any{
		"level1": map[string]any{
			"keep": "dst-val",
			"over": "dst-val",
		},
	}
	src := map[string]any{
		"level1": map[string]any{
			"over": "src-val",
			"add":  "new-val",
		},
	}
	mergeMap(dst, src)

	inner, ok := dst["level1"].(map[string]any)
	if !ok {
		t.Fatalf("expected nested map, got %T", dst["level1"])
	}
	if inner["keep"] != "dst-val" {
		t.Errorf("expected keep=dst-val, got %v", inner["keep"])
	}
	if inner["over"] != "src-val" {
		t.Errorf("expected over=src-val (src wins), got %v", inner["over"])
	}
	if inner["add"] != "new-val" {
		t.Errorf("expected add=new-val, got %v", inner["add"])
	}
}

func TestMergeMap_ScalarOverwritesMap(t *testing.T) {
	dst := map[string]any{
		"x": map[string]any{"nested": "val"},
	}
	src := map[string]any{
		"x": "scalar-now",
	}
	mergeMap(dst, src)
	if dst["x"] != "scalar-now" {
		t.Errorf("expected scalar to overwrite map, got %v", dst["x"])
	}
}

// ---------------------------------------------------------------------------
// LoadFile — 3-tier cascade merge
// ---------------------------------------------------------------------------

func TestLoadFile_ThreeTierCascade(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json":  `{"greeting":"Hello","farewell":"Goodbye"}`,
		"en/general/labels.json": `{"greeting":"Hi there"}`,
		"en/retail/labels.json":  `{"farewell":"See ya"}`,
	})

	var result map[string]string
	err := p.LoadFile("en", "retail", "labels.json", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// greeting: common="Hello" overridden by general="Hi there", retail has no greeting
	if result["greeting"] != "Hi there" {
		t.Errorf("expected greeting='Hi there', got %q", result["greeting"])
	}
	// farewell: common="Goodbye" overridden by retail="See ya"
	if result["farewell"] != "See ya" {
		t.Errorf("expected farewell='See ya', got %q", result["farewell"])
	}
}

func TestLoadFile_CommonOnly(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{"key":"common-val"}`,
	})

	var result map[string]string
	err := p.LoadFile("en", "retail", "labels.json", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["key"] != "common-val" {
		t.Errorf("expected common-val, got %q", result["key"])
	}
}

func TestLoadFile_ErrorWhenNoTierExists(t *testing.T) {
	p := testProvider(map[string]string{})

	var result map[string]string
	err := p.LoadFile("en", "retail", "missing.json", &result)
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestLoadFile_GeneralBusinessTypeSkipsGeneralTier(t *testing.T) {
	// When businessType == "general", tiers are [common, general] (no duplicate general).
	p := testProvider(map[string]string{
		"en/common/labels.json":  `{"a":"common"}`,
		"en/general/labels.json": `{"a":"general"}`,
	})

	var result map[string]string
	err := p.LoadFile("en", "general", "labels.json", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["a"] != "general" {
		t.Errorf("expected general to win over common, got %q", result["a"])
	}
}

// ---------------------------------------------------------------------------
// LoadPath — dot-path extraction
// ---------------------------------------------------------------------------

func TestLoadPath_ExtractsDotPath(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/client.json": `{"client":{"name":"Name","errors":{"not_found":"Not found"}}}`,
	})

	var result map[string]any
	err := p.LoadPath("en", "common", "client.json", "client", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["name"] != "Name" {
		t.Errorf("expected name='Name', got %v", result["name"])
	}
}

func TestLoadPath_NestedDotPath(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/client.json": `{"client":{"errors":{"not_found":"Not found"}}}`,
	})

	var result map[string]string
	err := p.LoadPath("en", "common", "client.json", "client.errors", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["not_found"] != "Not found" {
		t.Errorf("expected not_found='Not found', got %q", result["not_found"])
	}
}

func TestLoadPath_EmptyPathReturnsEntireFile(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{"key":"value"}`,
	})

	var result map[string]string
	err := p.LoadPath("en", "common", "labels.json", "", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["key"] != "value" {
		t.Errorf("expected key='value', got %q", result["key"])
	}
}

func TestLoadPath_ErrorForMissingSegment(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{"a":"b"}`,
	})

	var result map[string]string
	err := p.LoadPath("en", "common", "labels.json", "nonexistent", &result)
	if err == nil {
		t.Fatal("expected error for missing path segment, got nil")
	}
}

// ---------------------------------------------------------------------------
// LoadPathIfExists — nil when absent
// ---------------------------------------------------------------------------

func TestLoadPathIfExists_NilWhenAbsent(t *testing.T) {
	p := testProvider(map[string]string{})

	var result map[string]string
	err := p.LoadPathIfExists("en", "retail", "missing.json", "any", &result)
	if err != nil {
		t.Fatalf("expected nil error for absent file, got: %v", err)
	}
	if result != nil {
		t.Errorf("expected nil result for absent file, got %v", result)
	}
}

func TestLoadPathIfExists_LoadsWhenPresent(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{"section":{"key":"val"}}`,
	})

	var result map[string]string
	err := p.LoadPathIfExists("en", "retail", "labels.json", "section", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["key"] != "val" {
		t.Errorf("expected key='val', got %q", result["key"])
	}
}

func TestLoadPathIfExists_CascadeMerge(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{"section":{"a":"common-a","b":"common-b"}}`,
		"en/retail/labels.json": `{"section":{"a":"retail-a"}}`,
	})

	var result map[string]string
	err := p.LoadPathIfExists("en", "retail", "labels.json", "section", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["a"] != "retail-a" {
		t.Errorf("expected a='retail-a', got %q", result["a"])
	}
	if result["b"] != "common-b" {
		t.Errorf("expected b='common-b' (preserved from common), got %q", result["b"])
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive LoadFile tests
// ---------------------------------------------------------------------------

func TestLoadFile_InvalidJSONContent_ReturnsError(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{this is not valid JSON`,
	})

	var result map[string]string
	err := p.LoadFile("en", "retail", "labels.json", &result)
	if err == nil {
		t.Fatal("expected error for invalid JSON content, got nil")
	}
}

func TestLoadFile_EmptyJSON_ReturnsEmptyMap(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/labels.json": `{}`,
	})

	var result map[string]string
	err := p.LoadFile("en", "retail", "labels.json", &result)
	if err != nil {
		t.Fatalf("unexpected error for empty JSON: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty map for {} JSON, got %d entries", len(result))
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive LoadPath tests
// ---------------------------------------------------------------------------

func TestLoadPath_DeeplyNestedPath(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/deep.json": `{"a":{"b":{"c":{"d":{"e":{"f":"deep-value"}}}}}}`,
	})

	var result map[string]string
	err := p.LoadPath("en", "common", "deep.json", "a.b.c.d.e", &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["f"] != "deep-value" {
		t.Errorf("expected f='deep-value', got %q", result["f"])
	}
}

func TestLoadPath_DeeplyNestedPath_MissingIntermediate(t *testing.T) {
	p := testProvider(map[string]string{
		"en/common/deep.json": `{"a":{"b":"leaf"}}`,
	})

	var result map[string]string
	err := p.LoadPath("en", "common", "deep.json", "a.b.c.d.e.f", &result)
	if err == nil {
		t.Fatal("expected error for path traversal through non-map value, got nil")
	}
}

// ---------------------------------------------------------------------------
// Negative / defensive mergeMap tests
// ---------------------------------------------------------------------------

func TestMergeMap_EmptySrc(t *testing.T) {
	dst := map[string]any{"a": "1", "b": "2"}
	src := map[string]any{}
	mergeMap(dst, src)

	if dst["a"] != "1" {
		t.Errorf("expected a='1' preserved, got %v", dst["a"])
	}
	if dst["b"] != "2" {
		t.Errorf("expected b='2' preserved, got %v", dst["b"])
	}
}

func TestMergeMap_EmptyDst(t *testing.T) {
	dst := map[string]any{}
	src := map[string]any{"x": "y", "z": "w"}
	mergeMap(dst, src)

	if dst["x"] != "y" {
		t.Errorf("expected x='y', got %v", dst["x"])
	}
	if dst["z"] != "w" {
		t.Errorf("expected z='w', got %v", dst["z"])
	}
}
