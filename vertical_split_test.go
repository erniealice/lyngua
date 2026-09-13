package lyngua_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"regexp"
	"sort"
	"strings"
	"testing"

	lyngua "github.com/erniealice/lyngua"
	v1 "github.com/erniealice/lyngua/golang/v1"
)

var clientWord = regexp.MustCompile(`(?i)\bclients?\b`)

func TestLeasingVerticalsAreEmbeddedAndUseKnownKeys(t *testing.T) {
	for _, businessType := range []string{"leasing", "equipment_leasing"} {
		t.Run(businessType, func(t *testing.T) {
			entries, err := fs.ReadDir(lyngua.TranslationsFS, "translations/en/"+businessType)
			if err != nil {
				t.Fatalf("read embedded %s tier: %v", businessType, err)
			}
			if len(entries) < 10 {
				t.Fatalf("%s tier contains only %d files; expected a composed vertical overlay", businessType, len(entries))
			}

			for _, entry := range entries {
				if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
					continue
				}
				assertKnownOverlayKeys(t, businessType, entry.Name())
			}
		})
	}
}

func TestLeasingVerticalVocabularyIsSplit(t *testing.T) {
	provider := v1.NewTranslationProviderFromFS(lyngua.TranslationsFS)
	tests := []struct {
		businessType string
		file         string
		root         string
		path         []string
		want         string
	}{
		{"leasing", "client.json", "client", []string{"page", "heading"}, "Tenants"},
		{"equipment_leasing", "client.json", "client", []string{"page", "heading"}, "Lessees"},
		{"leasing", "subscription_group.json", "subscription_group", []string{"page", "title"}, "Lease Portfolios"},
		{"equipment_leasing", "subscription_group.json", "subscription_group", []string{"page", "title"}, "Fleet Master Agreements"},
		{"leasing", "job.json", "job", []string{"page", "heading"}, "Work Orders"},
		{"equipment_leasing", "job.json", "job", []string{"page", "heading"}, "Service Engagements"},
		{"leasing", "location.json", "", []string{"dashboard", "total_locations"}, "Total Spaces & Units"},
		{"equipment_leasing", "location.json", "", []string{"dashboard", "total_locations"}, "Total Yards & Sites"},
		{"leasing", "conversation.json", "conversation", []string{"list", "title"}, "Maintenance Tickets"},
		{"equipment_leasing", "conversation.json", "conversation", []string{"list", "title"}, "Maintenance Tickets"},
	}
	for _, test := range tests {
		t.Run(test.businessType+"/"+test.file+"/"+strings.Join(test.path, "."), func(t *testing.T) {
			var got map[string]any
			if err := provider.LoadPath("en", test.businessType, test.file, test.root, &got); err != nil {
				t.Fatal(err)
			}
			if value := stringAt(got, test.path...); value != test.want {
				t.Fatalf("got %q, want %q", value, test.want)
			}
		})
	}
}

func TestLeasingVerticalsReplaceGenericClientVocabulary(t *testing.T) {
	provider := v1.NewTranslationProviderFromFS(lyngua.TranslationsFS)
	for _, businessType := range []string{"leasing", "equipment_leasing"} {
		t.Run(businessType, func(t *testing.T) {
			messages, err := provider.LoadMessages("en", businessType)
			if err != nil {
				t.Fatal(err)
			}
			var failures []string
			for key, value := range messages {
				visible := regexp.MustCompile(`\{\{.*?\}\}|\{.*?\}`).ReplaceAllString(value, "")
				if clientWord.MatchString(visible) {
					failures = append(failures, fmt.Sprintf("%s=%q", key, value))
				}
			}
			sort.Strings(failures)
			if len(failures) != 0 {
				t.Fatalf("generic client vocabulary remains in %s:\n%s", businessType, strings.Join(failures, "\n"))
			}
		})
	}
}

func assertKnownOverlayKeys(t *testing.T, businessType, name string) {
	t.Helper()
	overlay := readJSONObject(t, "translations/en/"+businessType+"/"+name)
	base := make(map[string]any)
	for _, tier := range []string{"common", "general"} {
		path := "translations/en/" + tier + "/" + name
		data, err := fs.ReadFile(lyngua.TranslationsFS, path)
		if err != nil {
			if !strings.Contains(err.Error(), "file does not exist") {
				t.Fatalf("read %s: %v", path, err)
			}
			continue
		}
		var decoded map[string]any
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("decode %s: %v", path, err)
		}
		merge(base, decoded)
	}
	if len(base) == 0 {
		t.Fatalf("%s/%s has no common/general source file", businessType, name)
	}
	for _, path := range scalarPaths(overlay, nil) {
		if !pathExists(base, path) {
			t.Errorf("%s/%s adds unknown leaf %s", businessType, name, strings.Join(path, "."))
		}
	}
}

func readJSONObject(t *testing.T, path string) map[string]any {
	t.Helper()
	data, err := fs.ReadFile(lyngua.TranslationsFS, path)
	if err != nil {
		t.Fatal(err)
	}
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatal(err)
	}
	return result
}

func merge(dst, src map[string]any) {
	for key, value := range src {
		if child, ok := value.(map[string]any); ok {
			if existing, ok := dst[key].(map[string]any); ok {
				merge(existing, child)
				continue
			}
		}
		dst[key] = value
	}
}

func scalarPaths(value any, prefix []string) [][]string {
	object, ok := value.(map[string]any)
	if !ok {
		return [][]string{prefix}
	}
	var result [][]string
	for key, child := range object {
		result = append(result, scalarPaths(child, append(append([]string{}, prefix...), key))...)
	}
	return result
}

func pathExists(value any, path []string) bool {
	current := value
	for _, key := range path {
		object, ok := current.(map[string]any)
		if !ok {
			return false
		}
		current, ok = object[key]
		if !ok {
			return false
		}
	}
	return true
}

func stringAt(value map[string]any, path ...string) string {
	var current any = value
	for _, key := range path {
		current = current.(map[string]any)[key]
	}
	result, _ := current.(string)
	return result
}
