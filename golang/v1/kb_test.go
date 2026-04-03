package v1_test

import (
	"strings"
	"testing"
	"testing/fstest"

	v1 "github.com/erniealice/lyngua/golang/v1"
)

func newTestProvider(files map[string]string) *v1.TranslationProvider {
	fsMap := make(fstest.MapFS)
	for path, content := range files {
		fsMap["translations/"+path] = &fstest.MapFile{Data: []byte(content)}
	}
	return v1.NewTranslationProviderFromFS(fsMap)
}

func TestLoadKBIfExists_ReturnsNilForMissingTopic(t *testing.T) {
	p := newTestProvider(map[string]string{})

	kb, err := p.LoadKBIfExists("en", "retail", "nonexistent-topic")
	if err != nil {
		t.Fatalf("expected nil error for missing topic, got: %v", err)
	}
	if kb != nil {
		t.Fatal("expected nil KBContent for missing topic, got non-nil")
	}
}

func TestLoadKBIfExists_ReturnsContentForExistingTopic(t *testing.T) {
	p := newTestProvider(map[string]string{
		"en/common/kb/client.md": "# Clients\n\nManage your clients here.\n",
	})

	kb, err := p.LoadKBIfExists("en", "retail", "client")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if kb == nil {
		t.Fatal("expected KBContent, got nil")
	}
	if kb.Topic != "client" {
		t.Errorf("expected topic 'client', got %q", kb.Topic)
	}
	if kb.Title != "Clients" {
		t.Errorf("expected title 'Clients', got %q", kb.Title)
	}
	if !strings.Contains(string(kb.Body), "Manage your clients") {
		t.Errorf("expected body to contain 'Manage your clients', got: %s", kb.Body)
	}
}

func TestLoadKB_CascadeBusinessTypeWins(t *testing.T) {
	p := newTestProvider(map[string]string{
		"en/common/kb/client.md": "# Clients Common\n\nCommon content.\n",
		"en/retail/kb/client.md": "# Customers\n\nRetail content.\n",
	})

	kb, err := p.LoadKB("en", "retail", "client")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if kb.Title != "Customers" {
		t.Errorf("expected retail override title 'Customers', got %q", kb.Title)
	}
	if !strings.Contains(string(kb.Body), "Retail content") {
		t.Errorf("expected retail body, got: %s", kb.Body)
	}
}

func TestLoadKB_CommonFallbackWhenNoOverride(t *testing.T) {
	p := newTestProvider(map[string]string{
		"en/common/kb/role.md": "# Roles\n\nManage roles.\n",
	})

	kb, err := p.LoadKB("en", "retail", "role")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if kb.Title != "Roles" {
		t.Errorf("expected common title 'Roles', got %q", kb.Title)
	}
}

func TestLoadKB_ErrorForMissingTopic(t *testing.T) {
	p := newTestProvider(map[string]string{})

	_, err := p.LoadKB("en", "retail", "nonexistent")
	if err == nil {
		t.Fatal("expected error for missing topic, got nil")
	}
	if !strings.Contains(err.Error(), "no KB file found") {
		t.Errorf("expected 'no KB file found' error, got: %v", err)
	}
}
