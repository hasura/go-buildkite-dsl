package step_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/hasura/go-buildkite-dsl/step"
)

func TestWait(t *testing.T) {
	continueOnFailure := true
	ifCondition := "build.branch == 'main'"
	tt := []struct {
		input   step.Wait
		want    string
		wantErr bool
	}{
		{
			input: step.Wait{},
			want:  `{"wait":"~"}`,
		},
		{
			input: step.Wait{
				ContinueOnFailure: &continueOnFailure,
				If:                &ifCondition,
			},
			want: `{"wait":"~","continue_on_failure":true,"if":"build.branch == 'main'"}`,
		},
	}

	for _, tc := range tt {
		got, err := json.Marshal(tc.input)
		if err != nil {
			if tc.wantErr {
				continue
			}
			t.Errorf("error while json marshalling: %s", err)
		}

		if !bytes.Equal(got, []byte(tc.want)) {
			t.Errorf("want: %s, got: %s", tc.want, string(got))
		}
	}
}
