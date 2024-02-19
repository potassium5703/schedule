package parser

import "testing"

func TestContainsWord(t *testing.T) {
	tests := []struct {
		input []string
		want  bool
	}{
		{
			[]string{
				"card-header card",
				"card",
			},
			true,
		},
		{
			[]string{
				"card-body card-header",
				"card-header",
			},
			true,
		},
		{
			[]string{
				"card-body card-header",
				"card",
			},
			false,
		},
	}

	for _, test := range tests {
		got := containsWord(
			test.input[0],
			test.input[1],
		)
		if got != test.want {
			t.Errorf("got: %t, want %t", got, test.want)
		}
	}
}
