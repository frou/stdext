package stdext

import (
	"testing"
	"testing/quick"
)

// QuickCheckStrings tests that the function pred holds true when bombarded by
// randomly generated strings.
func QuickCheckStrings(t *testing.T, pred func(string) bool) {
	QuickCheckStringsN(t, 0 /*Use Default*/, pred)
}

// QuickCheckStringsN tests that the function pred holds true when bombarded by
// n randomly generated strings.
func QuickCheckStringsN(t *testing.T, n int, pred func(string) bool) {
	cfg := quick.Config{
		MaxCount: n,
	}
	if err := quick.Check(pred, &cfg); err != nil {
		t.Error(err)
	}
}
