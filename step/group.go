package step

import "encoding/json"

// Group is used to created a Group step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/group-step
type Group struct {
	// Name of the group in the UI.
	Group string `json:"group"`

	// A list of steps in the group; at least 1 step is required.
	Steps []GroupStep `json:"steps"`

	// Whether to continue to run this step if any of the steps named in the depends_on attribute fail.
	AllowDependencyFailure *bool `json:"allow_dependency_failure,omitempty"`

	// A list of step or group keys that this step depends on.
	// This step or group will only run after the named steps have completed.
	// See [managing step dependencies] for more information.
	//
	// [managing step dependencies]: https://buildkite.com/docs/pipelines/dependencies
	// Example: "test-suite"
	DependsOn []Dependency `json:"depends_on,omitempty"`

	// A unique string to identify the step, block, or group.
	//
	// Example: "test-suite"
	Key *string `json:"key,omitempty"`

	// The label that will be displayed in the pipeline visualisation in Buildkite
	// (name of the group in the UI). Supports emoji.
	//
	// Example: ":hammer: Tests" will be rendered as "🔨 Tests"
	Label *string `json:"label,omitempty"`
}

func (g Group) MarshalJSON() ([]byte, error) {
	if len(g.Group) == 0 {
		g.Group = "~"
	}

	type psuedo Group
	p := psuedo(g)

	return json.Marshal(p)
}

type GroupStep interface {
	Groupable()
}
