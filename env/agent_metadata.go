package env

import (
	"fmt"
	"os"
)

// The value of each agent tag. The tag name is appended to the end of the variable name.
// They can be set using the --tags flag on the buildkite-agent start command, or in the agent configuration file.
// The Queue tag is specifically used for isolating jobs and agents, and appears as the BUILDKITE_AGENT_META_DATA_QUEUE environment variable.
//
// Example:
// "BUILDKITE_AGENT_META_DATA_TAGNAME=tagvalue",
// "BUILDKITE_AGENT_META_DATA_QUEUE=some-queue"
func BuildkiteAgentMetadata(key string) string {
	return os.Getenv(
		fmt.Sprintf("BUILDKITE_AGENT_META_DATA_%s", key),
	)
}
