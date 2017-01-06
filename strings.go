package stdext

import "strings"

// NormalizeFileExt normalises a file extension path fragment to have a leading
// dot and be lowercase.
func NormalizeFileExt(ext string) string {
	ext = strings.TrimSpace(ext)
	ext = strings.ToLower(ext)
	if dot := byte('.'); ext[0] != dot {
		ext = string(dot) + ext
	}
	return ext
}
