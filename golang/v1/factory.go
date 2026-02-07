package v1

import (
	"path/filepath"
	"runtime"
)

// NewDefaultTranslationProvider creates a TranslationProvider with the default translations path.
// Automatically resolves the path using runtime.Caller, making it work from any calling package.
func NewDefaultTranslationProvider() *TranslationProvider {
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
