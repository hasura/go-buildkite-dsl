package env

import "testing"

func TestBuildkiteAgentMetdata(t *testing.T) {
	key := "BUILDKITE_AGENT_META_DATA_QUEUE"
	value := "aws-spot-ubuntu-small"
	t.Setenv(key, value)

	foundValue := BuildkiteAgentMetadata("QUEUE")
	if foundValue != value {
		t.Errorf("expected: %s, found: %s", value, foundValue)
	}
}
