package v1

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sync"
)

// TranslationProvider handles loading and caching translation messages.
type TranslationProvider struct {
	translationsPath string
	fsys             fs.FS
	cache            map[string]map[string]string // locale_businessType -> key -> message
	mutex            sync.RWMutex
}

// NewTranslationProvider creates a new TranslationProvider.
func NewTranslationProvider(translationsPath string) *TranslationProvider {
	return &TranslationProvider{
		translationsPath: translationsPath,
		cache:            make(map[string]map[string]string),
	}
}

// TranslationsPath returns the resolved translations directory path.
func (p *TranslationProvider) TranslationsPath() string {
	return p.translationsPath
}

// LoadMessages loads and merges translation messages for a given locale and business type.
// Loading priority (highest wins): businessType -> general -> common.
// Returns a flat map with dot-notation keys (e.g., "client.errors.not_found").
func (p *TranslationProvider) LoadMessages(locale, businessType string) (map[string]string, error) {
	cacheKey := fmt.Sprintf("%s_%s", locale, businessType)

	p.mutex.RLock()
	if messages, ok := p.cache[cacheKey]; ok {
		p.mutex.RUnlock()
		return messages, nil
	}
	p.mutex.RUnlock()

	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Double-check after acquiring lock
	if messages, ok := p.cache[cacheKey]; ok {
		return messages, nil
	}

	mergedMessages := make(map[string]string)

	// 1. Load common translations (base layer)
	commonPath := p.joinPath(locale, "common")
	if err := p.loadDirectory(commonPath, mergedMessages); err != nil {
		fmt.Printf("Warning: failed to load common translations for %s: %v\n", locale, err)
	}

	// 2. Load general translations (middle layer)
	if businessType != "general" {
		generalPath := p.joinPath(locale, "general")
		if err := p.loadDirectory(generalPath, mergedMessages); err != nil {
			fmt.Printf("Warning: failed to load general translations for %s: %v\n", locale, err)
		}
	}

	// 3. Load business-specific translations (top layer, highest priority)
	businessPath := p.joinPath(locale, businessType)
	if err := p.loadDirectory(businessPath, mergedMessages); err != nil {
		return nil, fmt.Errorf("failed to load %s translations for %s: %w", businessType, locale, err)
	}

	p.cache[cacheKey] = mergedMessages
	return mergedMessages, nil
}

// readFile reads a file from the embedded fs.FS or the OS filesystem.
func (p *TranslationProvider) readFile(relPath string) ([]byte, error) {
	if p.fsys != nil {
		return fs.ReadFile(p.fsys, path.Join(p.translationsPath, relPath))
	}
	return os.ReadFile(filepath.Join(p.translationsPath, relPath))
}

// readDir reads a directory from the embedded fs.FS or the OS filesystem.
func (p *TranslationProvider) readDir(relPath string) ([]fs.DirEntry, error) {
	if p.fsys != nil {
		return fs.ReadDir(p.fsys, path.Join(p.translationsPath, relPath))
	}
	return os.ReadDir(filepath.Join(p.translationsPath, relPath))
}

// joinPath joins path segments using forward slashes for fs.FS or OS-specific separators.
func (p *TranslationProvider) joinPath(segments ...string) string {
	if p.fsys != nil {
		return path.Join(segments...)
	}
	return filepath.Join(segments...)
}

// loadDirectory reads all JSON files from a directory and merges them into the provided map.
func (p *TranslationProvider) loadDirectory(dirPath string, targetMap map[string]string) error {
	files, err := p.readDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := p.joinPath(dirPath, file.Name())
		data, err := p.readFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		var messages map[string]any
		if err := json.Unmarshal(data, &messages); err != nil {
			return fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
		}

		flattenAndMerge(messages, "", targetMap)
	}

	return nil
}

// flattenAndMerge recursively flattens a nested map and merges it into the target map.
func flattenAndMerge(source map[string]any, prefix string, target map[string]string) {
	for key, value := range source {
		newKey := key
		if prefix != "" {
			newKey = prefix + "." + key
		}

		switch v := value.(type) {
		case map[string]any:
			flattenAndMerge(v, newKey, target)
		case string:
			target[newKey] = v
		default:
			target[newKey] = fmt.Sprintf("%v", v)
		}
	}
}
