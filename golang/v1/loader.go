package v1

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadFile loads a JSON translation file with business-type cascade and unmarshals
// into the provided target struct. Unlike LoadMessages (which flattens to dot-notation),
// this preserves the nested JSON structure for direct struct mapping.
//
// Loading priority (highest wins): businessType -> general -> common.
// Only files that exist are loaded; missing tiers are silently skipped.
//
// Example:
//
//	provider.LoadFile("en", "retail", "common.json", &labels)   // common -> general -> retail
//	provider.LoadFile("en", "retail", "client.json", &labels)   // common -> general -> retail
func (p *TranslationProvider) LoadFile(locale, businessType, fileName string, target any) error {
	merged := make(map[string]any)

	// Build cascade: common -> general -> businessType
	tiers := []string{"common"}
	if businessType != "general" && businessType != "common" {
		tiers = append(tiers, "general")
	}
	if businessType != "common" {
		tiers = append(tiers, businessType)
	}

	loaded := false
	for _, tier := range tiers {
		absPath := filepath.Join(p.translationsPath, locale, tier, fileName)

		data, err := os.ReadFile(absPath)
		if err != nil {
			continue // tier doesn't have this file, skip
		}

		var content map[string]any
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to unmarshal translation file %s: %w", absPath, err)
		}

		mergeMap(merged, content)
		loaded = true
	}

	if !loaded {
		return fmt.Errorf("translation file %s not found in any tier for %s/%s", fileName, locale, businessType)
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

// LoadPath loads a JSON translation file with business-type cascade, extracts the
// subtree at the given dot-notation path, then unmarshals into the provided target struct.
// Pass empty string for path to use the entire file (equivalent to LoadFile).
//
// Example:
//
//	provider.LoadPath("en", "retail", "client.json", "client", &labels)  // extracts "client" subtree
//	provider.LoadPath("en", "retail", "user.json", "", &labels)          // uses entire file
func (p *TranslationProvider) LoadPath(locale, businessType, fileName, path string, target any) error {
	merged := make(map[string]any)

	// Build cascade: common -> general -> businessType
	tiers := []string{"common"}
	if businessType != "general" && businessType != "common" {
		tiers = append(tiers, "general")
	}
	if businessType != "common" {
		tiers = append(tiers, businessType)
	}

	loaded := false
	for _, tier := range tiers {
		absPath := filepath.Join(p.translationsPath, locale, tier, fileName)

		data, err := os.ReadFile(absPath)
		if err != nil {
			continue // tier doesn't have this file, skip
		}

		var content map[string]any
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to unmarshal translation file %s: %w", absPath, err)
		}

		mergeMap(merged, content)
		loaded = true
	}

	if !loaded {
		return fmt.Errorf("translation file %s not found in any tier for %s/%s", fileName, locale, businessType)
	}

	// Extract subtree at the given path
	var result any = merged
	if path != "" {
		segments := strings.Split(path, ".")
		for _, seg := range segments {
			m, ok := result.(map[string]any)
			if !ok {
				return fmt.Errorf("path segment %q: expected nested map but got %T", seg, result)
			}
			val, exists := m[seg]
			if !exists {
				return fmt.Errorf("path segment %q not found in translation data", seg)
			}
			result = val
		}
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal merged translations: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return fmt.Errorf("failed to unmarshal merged translations into target: %w", err)
	}

	return nil
}

// LoadPathIfExists loads a JSON translation file with business-type cascade, extracts the
// subtree at the given dot-notation path, then unmarshals into the provided target struct.
// Returns nil if the file is absent in all tiers (this is expected — not all business types
// have overrides). Returns error only on parse/decode failure.
//
// Example:
//
//	provider.LoadPathIfExists("en", "retail", "product.json", "product", &labels)
func (p *TranslationProvider) LoadPathIfExists(locale, businessType, fileName, path string, target any) error {
	merged := make(map[string]any)

	// Build cascade: common -> general -> businessType
	tiers := []string{"common"}
	if businessType != "general" && businessType != "common" {
		tiers = append(tiers, "general")
	}
	if businessType != "common" {
		tiers = append(tiers, businessType)
	}

	loaded := false
	for _, tier := range tiers {
		absPath := filepath.Join(p.translationsPath, locale, tier, fileName)

		data, err := os.ReadFile(absPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue // tier doesn't have this file, skip
			}
			return fmt.Errorf("failed to read translation file %s: %w", absPath, err)
		}

		var content map[string]any
		if err := json.Unmarshal(data, &content); err != nil {
			return fmt.Errorf("failed to unmarshal translation file %s: %w", absPath, err)
		}

		mergeMap(merged, content)
		loaded = true
	}

	if !loaded {
		return nil
	}

	// Extract subtree at the given path
	var result any = merged
	if path != "" {
		segments := strings.Split(path, ".")
		for _, seg := range segments {
			m, ok := result.(map[string]any)
			if !ok {
				return fmt.Errorf("path segment %q: expected nested map but got %T", seg, result)
			}
			val, exists := m[seg]
			if !exists {
				return fmt.Errorf("path segment %q not found in translation data", seg)
			}
			result = val
		}
	}

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal merged translations: %w", err)
	}

	if err := json.Unmarshal(jsonBytes, target); err != nil {
		return fmt.Errorf("failed to unmarshal merged translations into target: %w", err)
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
