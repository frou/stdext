package stdext

import (
	"testing"
	"testing/quick"
)

func QuickCheckStrings(t *testing.T, pred func(string) bool) {
	QuickCheckStringsN(t, 0 /*Use Default*/, pred)
}

func QuickCheckStringsN(t *testing.T, count int, pred func(string) bool) {
	cfg := quick.Config{
		MaxCount: count,
	}
	if err := quick.Check(pred, &cfg); err != nil {
		t.Error(err)
	}
}
