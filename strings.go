package stdext

import "strings"

// NormalizeFileExt normalises a file extension (itself a fragment of a path
// string) to have a single leading dot, be free of whitespace, and be
// lowercased.
//
// Examples:
//   .txt   ->  .txt
//   JPG    ->  .jpg
//   .Go    ->  .go
//   Mp4    ->  .mp4
//   ..xml  ->  .xml
func NormalizeFileExt(ext string) string {
	ext = strings.TrimSpace(ext)
	ext = strings.TrimLeft(ext, ".")
	return "." + strings.ToLower(ext)
}

// Dotted joins the given component strings in order, using . characters as
// separators. Any leading or trailing dots that are part of a component
// itself are trimmed.
func Dotted(components ...string) string {
	for i := range components {
		components[i] = strings.Trim(components[i], ".")
	}
	return strings.Join(components, ".")
}
