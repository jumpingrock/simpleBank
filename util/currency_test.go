package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsSupportedCurrency(t *testing.T) {
	isSupported := IsSupportedCurrency("abc")
	require.Equal(t, isSupported, false)

	isSupported2 := IsSupportedCurrency("EUR")
	require.NotEmpty(t, isSupported2)
	require.Equal(t, isSupported2, true)
}
