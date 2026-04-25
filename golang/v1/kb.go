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

// LoadHelpHTML is a convenience wrapper around LoadKBIfExists that returns
// only the rendered HTML body and a found flag — the two values needed by
// the help-pane plumbing (PageData.HasHelp + PageData.HelpContent).
//
// It exists so callers in framework-level packages (pyeza/view) can match
// this signature via a tiny duck-typed interface without importing lyngua.
// Use `view.LoadHelpInto` from pyeza-golang for the one-liner at the call site.
func (p *TranslationProvider) LoadHelpHTML(locale, businessType, topic string) (template.HTML, bool) {
	kb, _ := p.LoadKBIfExists(locale, businessType, topic)
	if kb == nil {
		return "", false
	}
	return kb.Body, true
}
