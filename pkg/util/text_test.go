package util_test

import (
	"testing"

	"github.com/casss-dev/go_scraper/pkg/util"
)

func TestStripCurrencySymbols(t *testing.T) {
	// given
	const price = "US $1,239.99"

	// when
	stripped := util.StripCurrencySymbol(price)

	// then
	if stripped != "1239.99" {
		t.Errorf("expected 1239.99, but got '%s'", stripped)
	}
}
