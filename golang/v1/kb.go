package v1

import (
	"fmt"
	"html/template"
	"strings"
)

// KBContent is the parsed result of a single KB Markdown file.
type KBContent struct {
	Topic   string        // e.g. "clients", "sales-detail"
	Title   string        // extracted from first # H1 line
	Summary string        // first non-heading paragraph, truncated to 200 chars
	Body    template.HTML // full rendered HTML, safe to inject into {{.HelpContent}}
}

// LoadKB loads a KB Markdown file for the given locale and businessType using
// the same 3-tier cascade as LoadFile:
//
//	{locale}/common/kb/{topic}.md -> {locale}/general/kb/{topic}.md -> {locale}/{businessType}/kb/{topic}.md
//
// The highest-precedence file that exists wins (files are NOT merged).
// Returns an error only if a file exists but cannot be read or parsed.
func (p *TranslationProvider) LoadKB(locale, businessType, topic string) (*KBContent, error) {
	candidates := []string{
		p.joinPath(locale, "common", "kb", topic+".md"),
		p.joinPath(locale, "general", "kb", topic+".md"),
		p.joinPath(locale, businessType, "kb", topic+".md"),
	}

	var chosen string
	for _, c := range candidates {
		if _, err := p.readFile(c); err == nil {
			chosen = c
		}
	}
	if chosen == "" {
		return nil, fmt.Errorf("lyngua: no KB file found for topic %q (locale=%s, businessType=%s)", topic, locale, businessType)
	}

	return p.loadKBFile(chosen, topic)
}

// LoadKBIfExists is like LoadKB but returns (nil, nil) when no KB file exists
// for the given topic, so callers can skip KB rendering without treating
// missing content as an error.
func (p *TranslationProvider) LoadKBIfExists(locale, businessType, topic string) (*KBContent, error) {
	kb, err := p.LoadKB(locale, businessType, topic)
	if err != nil && strings.Contains(err.Error(), "no KB file found") {
		return nil, nil
	}
	return kb, err
}
