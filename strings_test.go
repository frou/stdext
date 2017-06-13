package stdext

import (
	"regexp"
	"strings"
	"testing"
	"testing/quick"
)

func Test_NormalizedFileExt_HasExactlyOneLeadingDot(t *testing.T) {
	quickCheck(t, func(s string) bool {
		ext := NormalizeFileExt(s)
		return strings.HasPrefix(ext, ".") && !strings.HasPrefix(ext, "..")
	})
}

func Test_NormalizedFileExt_HasNoUppercaseCharacters(t *testing.T) {
	// There is also a regexp class for Unicode uppercase - \p{Lu} - but every
	// uppercase Unicode character does not necessarily have a corresponding
	// lowercase character.
	reUpper := regexp.MustCompile(`[[:upper:]]`)
	quickCheck(t, func(s string) bool {
		ext := NormalizeFileExt(s)
		return !reUpper.MatchString(ext)
	})
}

func Test_NormalizedFileExt_HasNoSurroundingWhitespace(t *testing.T) {
	quickCheck(t, func(s string) bool {
		ext := NormalizeFileExt(s)
		return strings.TrimSpace(ext) == ext
	})
}

// ------------------------------------------------------------

func quickCheck(t *testing.T, pred func(string) bool) {
	cfg := quick.Config{
		MaxCount: 100000,
	}
	if err := quick.Check(pred, &cfg); err != nil {
		t.Error(err)
	}
}
