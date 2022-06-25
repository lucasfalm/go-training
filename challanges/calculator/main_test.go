package calculator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculator(t *testing.T) {
	require.Equal(t, 2.0, Calc("1 +1"))
	require.Equal(t, -2.0, Calc("-1- 1"))
	require.Equal(t, 1.0, Calc("1 /1"))
	require.Equal(t, 1.0, Calc("1 * 1"))
	require.Equal(t, 1476.0, Calc("12 * 123"))
	require.Equal(t, -1476.0, Calc("12 *(-123)"))
	require.Equal(t, 123.0, Calc("123"))
	require.Equal(t, 61.0, Calc("((80) - (19))"))
	require.Equal(t, -12.0, Calc("12*-1"))
	require.Equal(t, 3.0, Calc("(1 - 2) + -(-(-(-4)))"))

	// require.Equal(t, 21.25, Calc("2 /2+3 * 4.75- -6"))
}
