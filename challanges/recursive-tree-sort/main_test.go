package retreesort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaintains(t *testing.T) {
	tree := (`{
		"val": 1,
		"children": [
			{"val": 2},
			{
				"val": 3,
				"children": [
					{
						"val": 4,
						"children": [
							{"val": 5},
							{"val": 6},
							{"val": 7}
						]
					}
				]
			}
		]
	}`)
	expected := (`{
		"val": 1,
		"children": [
			{
				"val": 3,
				"children": [
					{
						"val": 4,
						"children": [
							{"val": 7},
							{"val": 5},
							{"val": 6}
						]
					}
				]
			},
			{"val": 2}
		]
	}`)

	sb := []string{}
	tc := copy(sb, tree)
	require.Equal(t, expected, PrioritizeNodes(&tree, 7))
}
