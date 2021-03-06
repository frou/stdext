package stdext

import (
	"regexp"
	"strings"
	"testing"

	"github.com/frou/testext"
)

const qcCount = 100000

func Test_NormalizedFileExt_HasExactlyOneLeadingDot(t *testing.T) {
	testext.QuickCheckStringsN(t, qcCount, func(s string) bool {
		ext := NormalizeFileExt(s)
		return strings.HasPrefix(ext, ".") && !strings.HasPrefix(ext, "..")
	})
}

func Test_NormalizedFileExt_HasNoUppercaseCharacters(t *testing.T) {
	// There is also a regexp class for Unicode uppercase - \p{Lu} - but every
	// uppercase Unicode character does not necessarily have a corresponding
	// lowercase character.
	reUpper := regexp.MustCompile(`[[:upper:]]`)
	testext.QuickCheckStringsN(t, qcCount, func(s string) bool {
		ext := NormalizeFileExt(s)
		return !reUpper.MatchString(ext)
	})
}

func Test_NormalizedFileExt_HasNoSurroundingWhitespace(t *testing.T) {
	testext.QuickCheckStringsN(t, qcCount, func(s string) bool {
		ext := NormalizeFileExt(s)
		return strings.TrimSpace(ext) == ext
	})
}
