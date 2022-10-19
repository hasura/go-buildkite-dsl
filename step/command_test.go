package step

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSoftFail(t *testing.T) {
	testCommand := `echo "testing"`
	tt := []struct {
		input   Command
		want    string
		wantErr bool
	}{
		{
			input: Command{
				Command:  &testCommand,
				SoftFail: ForAllNonZeroExitCodes(),
			},
			want: `{"command":"echo \"testing\"","soft_fail":true}`,
		},
		{
			input: Command{
				Command: &testCommand,
				SoftFail: ForExitCodes([]uint8{
					1, 2, 3, 4,
				}),
			},
			want: `{"command":"echo \"testing\"","soft_fail":[{"exit_status":1},{"exit_status":2},{"exit_status":3},{"exit_status":4}]}`,
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
