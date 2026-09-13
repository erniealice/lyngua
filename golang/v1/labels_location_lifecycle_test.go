package v1

import (
	lyngua "github.com/erniealice/lyngua"
	"testing"
)

func TestLocationLifecycleCatalog(t *testing.T) {
	p := NewTranslationProviderFromFS(lyngua.TranslationsFS)
	for _, vertical := range []string{"general", "unknown_vertical", "construction", "education", "laundry-services", "leasing", "manufacturing", "media", "medical-aesthetics", "mutual", "outsourcing", "professional", "retail", "service", "travel-and-tours"} {
		t.Run(vertical, func(t *testing.T) {
			var labels LocationLifecycleLabels
			if err := p.LoadPath("en", vertical, "location_lifecycle.json", "location_lifecycle", &labels); err != nil {
				t.Fatal(err)
			}
			if labels.Label == "" || labels.Help == "" {
				t.Fatal("missing lifecycle heading/help")
			}
			states := []LocationLifecycleStateLabels{labels.States.Planned, labels.States.Operational, labels.States.UnderMaintenance, labels.States.TemporarilyClosed, labels.States.Closed}
			for i, state := range states {
				if state.Label == "" || state.Help == "" {
					t.Fatalf("missing state %d label/help", i)
				}
			}
			if vertical == "education" && labels.Label != "Campus & Room Lifecycle" {
				t.Fatalf("education overlay not embedded: %q", labels.Label)
			}
			if vertical == "construction" && labels.Label != "Site Lifecycle" {
				t.Fatalf("construction overlay not embedded: %q", labels.Label)
			}
			if vertical == "leasing" && labels.Label != "Property & Space Lifecycle" {
				t.Fatalf("leasing overlay not applied: %q", labels.Label)
			}
			if labels.States.Operational.Label != "Operational" {
				t.Fatalf("general state fallback missing: %q", labels.States.Operational.Label)
			}
		})
	}
}
