package step

// Input is used to created a Block step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/input-step
type Input struct {
	// Input step name.
	//
	// Example: "Release"
	Input string `json:"input"`

	BlockAttributes
}
