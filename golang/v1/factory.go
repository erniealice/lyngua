package v1

import (
	"os"
	"path/filepath"
	"runtime"
)

// NewDefaultTranslationProvider creates a TranslationProvider with the default translations path.
// Checks TRANSLATIONS_PATH env var first (for container deployment), then falls back to
// runtime.Caller resolution (for dev/workspace mode).
func NewDefaultTranslationProvider() *TranslationProvider {
	// Container deployment: explicit path via env var
	if envPath := os.Getenv("TRANSLATIONS_PATH"); envPath != "" {
		return NewTranslationProvider(envPath)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return NewTranslationProvider(filepath.Join("..", "..", "translations"))
	}

	// filename is .../packages/lyngua-ryta/golang/v1/factory.go
	// We want .../packages/lyngua-ryta/translations
	lynguaDir := filepath.Dir(filepath.Dir(filepath.Dir(filename)))
	translationsPath := filepath.Join(lynguaDir, "translations")

	return NewTranslationProvider(translationsPath)
}

// NewDefaultTranslationProviderWithWorkspace creates a TranslationProvider that resolves
// the translations path from the go.work root. This is useful when lyngua is consumed
// as a workspace module (via go.work replace directives) rather than as a published package.
func NewDefaultTranslationProviderWithWorkspace() *TranslationProvider {
	dir, err := os.Getwd()
	if err != nil {
		return NewDefaultTranslationProvider()
	}

	// Walk up to find go.work root
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.work")); err == nil {
			translationsPath := filepath.Join(dir, "packages", "lyngua-ryta", "translations")
			return NewTranslationProvider(translationsPath)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// No go.work found, fall back to runtime.Caller resolution
			return NewDefaultTranslationProvider()
		}
		dir = parent
	}
}
