package lyngua

import "embed"

//go:embed translations/en/common/*.json
//go:embed translations/en/general/*.json
//go:embed translations/en/outsourcing/*.json
//go:embed translations/en/retail/*.json
//go:embed translations/en/service/*.json
//go:embed translations/en/professional/*.json
//go:embed translations/en/education/*.json
//go:embed translations/en/leasing/*.json
//go:embed translations/en/equipment_leasing/*.json
//go:embed translations/en/common/kb/*.md
//go:embed translations/en/retail/kb/*.md
//go:embed translations/en/service/kb/*.md
//go:embed translations/en/construction/location_lifecycle.json
//go:embed translations/en/laundry-services/location_lifecycle.json
//go:embed translations/en/manufacturing/location_lifecycle.json
//go:embed translations/en/media/location_lifecycle.json
//go:embed translations/en/medical-aesthetics/location_lifecycle.json
//go:embed translations/en/mutual/location_lifecycle.json
//go:embed translations/en/travel-and-tours/location_lifecycle.json
var TranslationsFS embed.FS
