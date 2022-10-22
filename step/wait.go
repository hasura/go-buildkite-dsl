package step

import (
	"encoding/json"
)

// Wait is used to created a Wait step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/wait-step
type Wait struct {
	// Run the next step, even if the previous step has failed.
	ContinueOnFailure *bool `json:"continue_on_failure,omitempty"`

	// A boolean expression that omits the step when false. See [Using conditionals] for supported expressions.
	//
	// Example: build.message != "skip me"
	//
	// [Using conditionals]: https://buildkite.com/docs/pipelines/conditionals
	If *string `json:"if,omitempty"`

	// A list of step keys that this step depends on. This step will only
	// run after the named steps have completed.
	DependsOn []Dependency `json:"depends_on,omitempty"`

	// Whether to continue to run this step if any of the steps named in the depends_on attribute fail.
	AllowDependencyFailure *bool `json:"allow_dependency_failure,omitempty"`
}

func (w Wait) MarshalJSON() ([]byte, error) {
	type psuedo Wait
	type transformedWait struct {
		Wait string `json:"wait"`
		psuedo
	}

	return json.Marshal(transformedWait{
		Wait:   "~",
		psuedo: psuedo(w),
	})
}
