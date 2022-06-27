package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBICIdentifier(t *testing.T) {
	bic := BICFIDec2014Identifier("CLNOUS66")
	assert.Equal(t, "CLNO", bic.BankCode())
	assert.Equal(t, "US", bic.CountryCode())
	assert.Equal(t, "66", bic.LocationCode())
	assert.Equal(t, "", bic.BranchCode())

	bic = BICFIDec2014Identifier("CLNOUS66XXX")
	assert.Equal(t, "CLNO", bic.BankCode())
	assert.Equal(t, "US", bic.CountryCode())
	assert.Equal(t, "66", bic.LocationCode())
	assert.Equal(t, "XXX", bic.BranchCode())
}
