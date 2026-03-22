package v1

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// kbMarkdown is a single goldmark instance shared across all renders.
// goldmark is safe for concurrent use after construction.
var kbMarkdown = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithXHTML(),
		html.WithUnsafe(),
	),
)

// loadKBFile reads relPath from the provider's filesystem, renders Markdown
// to HTML, and extracts title + summary.
func (p *TranslationProvider) loadKBFile(relPath, topic string) (*KBContent, error) {
	raw, err := p.readFile(relPath)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := kbMarkdown.Convert(raw, &buf); err != nil {
		return nil, fmt.Errorf("lyngua: failed to render KB markdown %s: %w", relPath, err)
	}

	return &KBContent{
		Topic:   topic,
		Title:   extractKBTitle(string(raw)),
		Summary: extractKBSummary(string(raw)),
		Body:    template.HTML(buf.String()),
	}, nil
}

// extractKBTitle returns the text of the first # H1 line, or empty string if absent.
func extractKBTitle(src string) string {
	for _, line := range strings.SplitN(src, "\n", 20) {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return ""
}

// extractKBSummary returns the first non-heading, non-blank paragraph,
// truncated to 200 characters.
func extractKBSummary(src string) string {
	pastHeading := false
	for _, line := range strings.Split(src, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			pastHeading = true
			continue
		}
		if !pastHeading || line == "" {
			continue
		}
		if len(line) > 200 {
			return line[:200] + "…"
		}
		return line
	}
	return ""
}
