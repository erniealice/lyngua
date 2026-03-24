# lyngua

Translation package for Ichizen OS applications. Provides JSON-based i18n with business type layering and two loading modes for Go consumers.

## Structure

```
lyngua/
  translations/
    en/
      common/         UI chrome labels (sidebar, header, buttons, table)
      general/        Shared domain translations (client, validation, errors)
      retail/         Retail business type overrides (quote, regulation, etc.)
  golang/v1/          Go consumer package
```

## Translation Layering

Translations follow a priority chain: **common** -> **general** -> **{businessType}**

- `common/` - Shared UI component labels used by pyeza templates (CommonLabels struct)
- `general/` - Shared domain translations across all business types
- `retail/`, `education/`, etc. - Business-specific overrides (same keys override general)

## Go API

### Setup

```go
import lynguaV1 "leapfor.xyz/lyngua/golang/v1"

// Auto-resolves translations/ directory via runtime.Caller
provider := lynguaV1.NewDefaultTranslationProvider()

// Or provide explicit path
provider := lynguaV1.NewTranslationProvider("/path/to/translations")
```

### Structured Loading (for templates)

Preserves nested JSON structure. Use for typed struct mapping (e.g., pyeza CommonLabels).

```go
var labels pyeza.CommonLabels
err := provider.LoadFile("en", "common/common.json", &labels)

// Or load all files in a directory into one struct
err := provider.LoadDirectory("en", "common", &labels)
```

### Flat Loading (for backend lookups)

Flattens nested JSON to dot-notation keys. Use for use case error messages and validation.

```go
messages, err := provider.LoadMessages("en", "retail")
// messages["client.errors.not_found"] = "Client not found"
// messages["validation.email_invalid"] = "Email format is invalid"
```

`LoadMessages` loads `general/` first, then overlays `{businessType}/`. Results are cached and thread-safe.

## Adding Translations

1. Add/edit JSON files in `translations/{locale}/{category}/`
2. For new business types, create a directory: `translations/en/my_business/`
3. Override general keys by using the same filename and key structure

### JSON Format

```json
{
  "header": {
    "welcomeBack": "Welcome back",
    "searchPlaceholder": "Search anything..."
  },
  "buttons": {
    "save": "Save",
    "cancel": "Cancel"
  }
}
```

## Integration with Retail Admin

The retail-admin app loads labels at startup in `container.go` and injects them into every page via the view adapter:

```go
// container.go
translations := lynguaV1.NewDefaultTranslationProvider()
var commonLabels pyeza.CommonLabels
translations.LoadFile("en", "common/common.json", &commonLabels)

// Passed to HTTPAdapter -> ViewAdapter -> injectPageData (reflection)
// Every page's PageData.CommonLabels is set automatically before rendering
```

## Module

```
module: leapfor.xyz/lyngua
go: 1.25.1
```

Referenced in `go.work`:
```
replace leapfor.xyz/lyngua => ./packages/lyngua
```
