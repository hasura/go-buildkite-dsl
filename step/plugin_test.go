package step_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/hasura/go-buildkite-dsl/step"
)

type dockerPlugin struct {
	Image string `json:"image"`
}

func (dockerPlugin) Name() string {
	return "docker"
}

func (dockerPlugin) Version() string {
	return "5.2.0"
}

func (d dockerPlugin) MarshalJSON() ([]byte, error) {
	type psuedo dockerPlugin
	p := psuedo(d)
	return json.Marshal(p)
}

func TestPlugin(t *testing.T) {
	testCommand := `echo "testing"`
	tt := []struct {
		input   step.Command
		want    string
		wantErr bool
	}{
		{
			input: step.Command{
				Plugins: []step.Plugin{
					dockerPlugin{
						Image: "ubuntu:22.04",
					},
				},
				Command: &testCommand,
			},
			want: `{"command":"echo \"testing\"","plugins":{"docker#5.2.0":{"image":"ubuntu:22.04"}}}`,
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
