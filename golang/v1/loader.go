package v1

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadFile reads a JSON file from the translations directory and unmarshals it
// into the provided target struct. Unlike LoadMessages (which flattens to dot-notation),
// this preserves the nested JSON structure for direct struct mapping.
//
// The filePath is relative to translations/{locale}/, e.g.:
//
//	provider.LoadFile("en", "common/common.json", &labels)
//	provider.LoadFile("en", "retail/client.json", &clientLabels)
func (p *TranslationProvider) LoadFile(locale, filePath string, target any) error {
	absPath := filepath.Join(p.translationsPath, locale, filePath)

	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to read translation file %s: %w", absPath, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal translation file %s: %w", absPath, err)
	}

	return nil
}

// LoadDirectory reads all JSON files from a directory within translations/{locale}/
// and unmarshals the merged result into the provided target struct.
// Files are loaded alphabetically; later files override earlier ones for duplicate keys.
//
// Example:
//
//	provider.LoadDirectory("en", "common", &labels)
func (p *TranslationProvider) LoadDirectory(locale, dirName string, target any) error {
	dirPath := filepath.Join(p.translationsPath, locale, dirName)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	merged := make(map[string]any)

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(dirPath, file.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		var content map[string]any
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
		}

		mergeMap(merged, content)
	}

	jsonBytes, err := json.Marshal(merged)
	if err != nil {
		return fmt.Errorf("failed to marshal merged translations: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return fmt.Errorf("failed to unmarshal merged translations into target: %w", err)
	}

	return nil
}

// mergeMap recursively merges src into dst. src values take precedence.
func mergeMap(dst, src map[string]any) {
	for key, srcVal := range src {
		if dstVal, exists := dst[key]; exists {
			srcMap, srcOk := srcVal.(map[string]any)
			dstMap, dstOk := dstVal.(map[string]any)
			if srcOk && dstOk {
				mergeMap(dstMap, srcMap)
				continue
			}
		}
		dst[key] = srcVal
	}
}
