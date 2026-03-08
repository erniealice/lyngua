package lyngua

import "embed"

//go:embed translations/en/common/*.json
//go:embed translations/en/general/*.json
//go:embed translations/en/retail/*.json
//go:embed translations/en/service/*.json
var TranslationsFS embed.FS
