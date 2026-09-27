package v1

import (
	"testing"

	lyngua "github.com/erniealice/lyngua"
)

func TestLocationLabelDefaultsAcrossTiers(t *testing.T) {
	provider := NewTranslationProviderFromFS(lyngua.TranslationsFS)
	for _, businessType := range []string{"general", "education", "leasing"} {
		t.Run(businessType, func(t *testing.T) {
			var labels struct {
				Page struct {
					HeadingActive   string `json:"heading_active"`
					HeadingInactive string `json:"heading_inactive"`
				} `json:"page"`
				Buttons struct {
					AddLocation string `json:"add_location"`
				} `json:"buttons"`
				Columns struct {
					Name    string `json:"name"`
					Address string `json:"address"`
				} `json:"columns"`
				Empty struct {
					ActiveTitle string `json:"active_title"`
				} `json:"empty"`
				Form struct {
					Name   string `json:"name"`
					Active string `json:"active"`
				} `json:"form"`
			}
			if err := provider.LoadFile("en", businessType, "location.json", &labels); err != nil {
				t.Fatal(err)
			}
			for field, value := range map[string]string{
				"page.heading_active":   labels.Page.HeadingActive,
				"page.heading_inactive": labels.Page.HeadingInactive,
				"buttons.add_location":  labels.Buttons.AddLocation,
				"columns.name":          labels.Columns.Name,
				"columns.address":       labels.Columns.Address,
				"empty.active_title":    labels.Empty.ActiveTitle,
				"form.name":             labels.Form.Name,
			} {
				if value == "" {
					t.Errorf("%s label is empty", field)
				}
			}
			if labels.Form.Active != "Active" {
				t.Errorf("form.active = %q, want Active", labels.Form.Active)
			}
		})
	}
}
