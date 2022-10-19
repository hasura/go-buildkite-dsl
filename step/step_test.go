package step

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSkip(t *testing.T) {
	type input struct {
		Skip *Skip
	}
	tt := []struct {
		input input
		want  string
	}{
		{
			input: input{
				Skip: WithoutReason(),
			},
			want: `{"Skip":true}`,
		},
		{
			input: input{
				Skip: WithReason("this is a reason to skip"),
			},
			want: `{"Skip":"this is a reason to skip"}`,
		},
	}

	for _, tc := range tt {
		got, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("error while json marshalling: %+v", err)
		}

		if !bytes.Equal(got, []byte(tc.want)) {
			t.Errorf("want: %s, got: %s", tc.want, string(got))
		}
	}
}
