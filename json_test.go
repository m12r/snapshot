package snapshot_test

import (
	"testing"

	"github.com/m12r/snapshot"
)

func TestMatchJSON(t *testing.T) {
	testCases := []struct {
		name  string
		value any
	}{
		{
			name:  "nil",
			value: nil,
		},
		{
			name:  "int",
			value: 1,
		},
		{
			name:  "float",
			value: 0.1,
		},
		{
			name:  "true",
			value: true,
		},
		{
			name:  "false",
			value: false,
		},
		{
			name:  "string",
			value: "foo",
		},
		{
			name: "map",
			value: map[string]any{
				"z": 0.1,
				"y": 1,
				"x": true,
				"a": false,
				"b": "foo",
				"c": []int{1, 2, 3, 4},
				"d": []float32{0.1, 0.2, 0.3, 0.4},
				"e": []bool{true, false, true, false},
				"f": []string{"foo", "bar", "baz"},
				"g": map[string]any{
					"foo": "bar",
				},
			},
		},
		{
			name: "struct",
			value: struct {
				Foo string
				Bar int
			}{
				Foo: "foo",
				Bar: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			snapshot.MatchJSON(t, tc.value)
		})
	}
}
