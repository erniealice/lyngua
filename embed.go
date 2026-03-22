package lyngua

import "embed"

//go:embed translations/en/common/*.json
//go:embed translations/en/general/*.json
//go:embed translations/en/retail/*.json
//go:embed translations/en/service/*.json
//go:embed translations/en/professional/*.json
//go:embed translations/en/common/kb/*.md
//go:embed translations/en/retail/kb/*.md
//go:embed translations/en/service/kb/*.md
var TranslationsFS embed.FS
