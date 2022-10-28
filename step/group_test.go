package step_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/hasura/go-buildkite-dsl/step"
)

func TestGroup(t *testing.T) {
	testCommand := `echo "testing"`
	tt := []struct {
		input   step.Group
		want    string
		wantErr bool
	}{
		{
			input: step.Group{
				Group: "just command step",
				Steps: []step.GroupStep{
					step.Command{
						Command: &testCommand,
					},
				},
			},
			want: `{"group":"just command step","steps":[{"command":"echo \"testing\""}]}`,
		},
		{
			input: step.Group{
				Group: "just wait step",
				Steps: []step.GroupStep{
					step.Wait{},
				},
			},
			want: `{"group":"just wait step","steps":[{"wait":"~"}]}`,
		},
		{
			input: step.Group{
				Group: "just trigger step",
				Steps: []step.GroupStep{
					step.Trigger{
						PipelineSlug: "sample-pipeline",
					},
				},
			},
			want: `{"group":"just trigger step","steps":[{"trigger":"sample-pipeline"}]}`,
		},
		{
			input: step.Group{
				Group: "all possible type of group steps",
				Steps: []step.GroupStep{
					step.Command{
						Command: &testCommand,
					},
					step.Wait{},
					step.Trigger{
						PipelineSlug: "sample-pipeline",
					},
				},
			},
			want: `{"group":"all possible type of group steps","steps":[{"command":"echo \"testing\""},{"wait":"~"},{"trigger":"sample-pipeline"}]}`,
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
