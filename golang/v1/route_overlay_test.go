package v1

import (
	"testing"

	lyngua "github.com/erniealice/lyngua"
)

// The service tier has no current app consumer, so Lyngua owns its route values.
func TestServiceSubscriptionRouteOverlay(t *testing.T) {
	provider := NewTranslationProviderFromFS(lyngua.TranslationsFS)
	var routes struct {
		ListURL   string `json:"list_url"`
		DetailURL string `json:"detail_url"`
		AddURL    string `json:"add_url"`
		EditURL   string `json:"edit_url"`
		DeleteURL string `json:"delete_url"`
	}
	if err := provider.LoadPath("en", "service", "route.json", "subscription", &routes); err != nil {
		t.Fatal(err)
	}
	for field, tc := range map[string]struct{ got, want string }{
		"list":   {routes.ListURL, "/memberships/list/{status}"},
		"detail": {routes.DetailURL, "/memberships/detail/{id}"},
		"add":    {routes.AddURL, "/action/membership/add"},
		"edit":   {routes.EditURL, "/action/membership/edit/{id}"},
		"delete": {routes.DeleteURL, "/action/membership/delete"},
	} {
		if tc.got != tc.want {
			t.Errorf("service subscription %s route = %q, want %q", field, tc.got, tc.want)
		}
	}
}

func TestServiceClientAndPlanRouteOverlay(t *testing.T) {
	provider := NewTranslationProviderFromFS(lyngua.TranslationsFS)
	for _, tc := range []struct{ path, key, want string }{
		{"client", "list_url", "/customers/list/{status}"},
		{"plan", "list_url", "/packages/list/{status}"},
	} {
		var routes map[string]string
		if err := provider.LoadPath("en", "service", "route.json", tc.path, &routes); err != nil {
			t.Fatal(err)
		}
		if got := routes[tc.key]; got != tc.want {
			t.Errorf("service %s.%s = %q, want %q", tc.path, tc.key, got, tc.want)
		}
	}
}
