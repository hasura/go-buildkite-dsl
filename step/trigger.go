package step

import (
	"encoding/json"
	"errors"

	"github.com/hasura/go-buildkite-dsl/pipeline"
)

// Trigger is used to created a Trigger step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/trigger-step
type Trigger struct {
	// The slug of the pipeline to create a build. You can find it in the URL of your pipeline,
	// and it corresponds to the name of the pipeline, converted to kebab-case.
	//
	// Example: "another-pipeline"
	PipelineSlug pipeline.Slug `json:"trigger"`

	// An optional map of attributes for the triggered build.
	Build *Build `json:"build,omitempty"`

	// The label that will be displayed in the pipeline visualisation in Buildkite. Supports emoji.
	//
	// Example: ":rocket: Deploy"
	Label *string `json:"label,omitempty"`

	// If set to true the step will immediately continue, regardless of the success
	// of the triggered build. If set to false the step will wait for the triggered build to complete
	// and continue only if the triggered build passed. Note that when async is set to true, as long
	// as the triggered build starts, the original pipeline will show that as successful.
	// The original pipeline does not get updated after subsequent steps or after the triggered build completes.
	//
	// Default value: false
	Async *bool `json:"async,omitempty"`

	// The [branch pattern] defining which branches will include this step in their builds.
	//
	// Example: "main stable/*"
	//
	// [branch pattern]: https://buildkite.com/docs/pipelines/branch-configuration#branch-pattern-examples
	Branches *string `json:"branches,omitempty"`

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

	// Whether to skip this step or not.
	//
	// Use `WithReason(string)` or `WithoutReason()` to populate this.
	Skip *Skip `json:"skip,omitempty"`
}

type Build struct {
	// The message for the build. Supports emoji. Default: the label of the trigger step.
	//
	// Example: "Triggered build"
	Message *string `json:"message,omitempty"`

	// The commit hash for the build. Default: "HEAD"
	//
	// Example: "ca82a6d"
	Commit *string `json:"commit,omitempty"`

	// The branch for the build.
	//
	// Default: The triggered pipeline's default branch.
	// Example: "production"
	Branch *string `json:"branch,omitempty"`

	// A map of meta-data for the build.
	//
	// Example: release-version: "1.1"
	Metadata map[string]string `json:"meta_data,omitempty"`

	// A map of environment variables for the build.
	//
	// Example: RAILS_ENV: "test"
	Env map[string]string `json:"env,omitempty"`
}

func (t Trigger) MarshalJSON() ([]byte, error) {
	if len(t.PipelineSlug) == 0 {
		return nil, errors.New("trigger step should have a non-empty pipeline slug")
	}

	type psuedo Trigger
	p := psuedo(t)
	return json.Marshal(p)
}

// Step is defined on Trigger step to allow it to be used as [pipeline.Step]
func (Trigger) Step() {}

// Groupable is defined to allow Trigger step to be used as a [step.GroupStep]
func (Trigger) Groupable() {}
