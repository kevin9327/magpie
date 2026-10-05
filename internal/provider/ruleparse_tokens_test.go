package provider

import "testing"

func TestParseTokensRejectsNonFinite(t *testing.T) {
	for _, in := range []string{"nan", "NaN", "inf", "+inf", "1e20", "1e308", "1e20k"} {
		if n, err := ParseTokens(in); err == nil {
			t.Errorf("%q accepted as %d", in, n)
		}
	}
}
