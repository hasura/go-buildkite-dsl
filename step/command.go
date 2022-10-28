package step

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Command is used to created a Command step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/command-step
type Command struct {
	// The shell command to run during this step. This can be a single line of commands.
	//
	// Example: "build.sh"
	Command *string `json:"command,omitempty"`

	// Alias for Command. This is a list of commands that must all pass.
	//
	// Example:
	// - "npm install"
	// - "tests.sh"
	Commands []string `json:"commands,omitempty"`

	// A map of [agent tag] keys to values to [target specific agents] for this step.
	//
	// Example: npm: "true"
	//
	// [agent tag]: https://buildkite.com/docs/agent/v3/cli-start#setting-tags
	// [target specific agents]: https://buildkite.com/docs/agent/v3/cli-start#agent-targeting
	Agents interface{} `json:"agents,omitempty"` // TODO: check if there is a more apt type for this

	// Whether to continue to run this step if any of the steps named in the depends_on attribute fail.
	AllowDependencyFailure *bool `json:"allow_dependency_failure,omitempty"`

	// The [glob path] or paths of [artifacts] to upload from this step.
	//
	// [glob path]: https://buildkite.com/docs/agent/v3/cli-artifact#uploading-artifacts
	// [artifacts]: https://buildkite.com/docs/agent/v3/cli-artifact
	ArtifactPaths []string `json:"artifact_paths,omitempty"`

	// The [branch pattern] defining which branches will include this step in their builds.
	//
	// Example: "main stable/*"
	//
	// [branch pattern]: https://buildkite.com/docs/pipelines/branch-configuration#branch-pattern-examples
	Branches *string `json:"branches,omitempty"`

	// Setting this attribute to true cancels the job as soon as the build is marked as failing.
	CancelOnBuildFailing *bool `json:"cancel_on_build_failing,omitempty"`

	// The maximum number of jobs created from this step that are allowed to run at the same time.
	// If you use this attribute, you must also define a label for it with the concurrency_group attribute.
	//
	// Example: 3
	Concurrency *uint `json:"concurrency,omitempty"`

	// A unique name for the concurrency group that you are creating.
	// If you use this attribute, you must also define the concurrency attribute.
	//
	// Example: "my-app/deploy"
	ConcurrencyGroup *string `json:"concurrency_group,omitempty"`

	// A list of step keys that this step depends on. This step will only
	// run after the named steps have completed.
	DependsOn []Dependency `json:"depends_on,omitempty"`

	// A map of environment variables for this step.
	Env map[string]string `json:"env,omitempty"`

	// A boolean expression that omits the step when false. See [Using conditionals] for supported expressions.
	//
	// Example: build.message != "skip me"
	//
	// [Using conditionals]: https://buildkite.com/docs/pipelines/conditionals
	If *string `json:"if,omitempty"`

	// A unique string to identify the step. The value is available in the BUILDKITE_STEP_KEY environment variable.
	// Keys can not have the same pattern as a UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx).
	//
	// Example: "linter"
	Key *string `json:"key,omitempty"`

	// Alias for Key
	Identifier *string `json:"identifier,omitempty"`

	// The label that will be displayed in the pipeline visualisation in Buildkite. Supports [emoji].
	//
	// Example: ":hammer: Tests" will be rendered as "🔨 Tests"
	//
	// [emoji]: https://github.com/buildkite/emojis#emoji-reference
	Label *string `json:"label,omitempty"`

	// Array of values to be used in matrix expansion.
	//
	// This library intentionally does not support advanced matrix configurations like `setup` and `adjustments`.
	// Preference would be to instead use for loops in Go to construct a set of jobs that we aim for the matrix to contain.
	// Another reason is the limitation of such configurations: https://buildkite.com/docs/pipelines/build-matrix#matrix-limits
	Matrix []string `json:"matrix,omitempty"`

	// The number of parallel jobs that will be created based on this step.
	//
	// Example: 3
	Parallelism *uint `json:"parallelism,omitempty"`

	// An array of [plugins] for this step.
	//
	// [plugins]: https://buildkite.com/docs/plugins
	Plugins []Plugin `json:"plugins,omitempty"`

	// Adjust the priority for a specific job, as a positive or negative integer.
	Priority *int `json:"priority,omitempty"`

	// Whether to skip this step or not.
	//
	// Use `WithReason(string)` or `WithoutReason()` to populate this.
	Skip *Skip `json:"skip,omitempty"`

	// Allow specified non-zero exit statuses not to fail the build.
	//
	// Use `ForAllNonZeroExitCodes()` to soft fail for any non-zero exit code.
	// Use `ForExitCodes([]uint8)` to soft fail for specific exit codes.
	SoftFail *SoftFail `json:"soft_fail,omitempty"`

	// The maximum number of minutes a job created from this step is allowed to run.
	// If the job exceeds this time limit, or if it finishes with a non-zero exit status,
	// the job is automatically canceled and the build fails.
	// Jobs that time out with an exit status of 0 are marked as "passed".
	//
	// Note that command steps on the Buildkite Free plan have a maximum job timeout of 240 minutes.
	// You can also set default and maximum timeouts in the Buildkite UI.
	//
	// Example: 60
	TimeoutInMinutes uint `json:"timeout_in_minutes,omitempty"`
}

func (c Command) MarshalJSON() ([]byte, error) {
	if c.Command == nil && len(c.Commands) == 0 {
		return nil, errors.New("either Command or Commands field should be present")
	}

	// If we return json.Marshal(c) directly, then it will
	// lead to stack overflow as it is going to inifitely recurse in the json package.
	// Hence introducing a psuedo type with same data to avoid entering this method
	// during JSON marshalling.
	type psuedo Command

	transformedPlugins := make(map[string]Plugin)
	for _, plugin := range c.Plugins {
		key := fmt.Sprintf("%s#%s", plugin.Name(), plugin.Version())
		transformedPlugins[key] = plugin
	}

	return json.Marshal(struct {
		psuedo
		TransformedPlugins map[string]Plugin `json:"plugins,omitempty"`
	}{
		psuedo(c),
		transformedPlugins,
	})
}

type SoftFail struct {
	ExitCodes []uint8
}

func ForAllNonZeroExitCodes() *SoftFail {
	return &SoftFail{}
}

func ForExitCodes(exitCodes []uint8) *SoftFail {
	return &SoftFail{
		ExitCodes: exitCodes,
	}
}

func (s SoftFail) MarshalJSON() ([]byte, error) {
	if len(s.ExitCodes) == 0 {
		return []byte("true"), nil
	}

	type exitStatus struct {
		ExitStatus uint8 `json:"exit_status"`
	}

	var exitStatuses []exitStatus
	for _, exitCode := range s.ExitCodes {
		exitStatuses = append(exitStatuses, exitStatus{
			ExitStatus: exitCode,
		})
	}

	return json.Marshal(exitStatuses)
}

type CommandRetry struct {
	Automatic *AutomaticRetry `json:"automatic,omitempty"`
	Manual    *ManualRetry    `json:"manual,omitempty"`
}

type AutomaticRetry struct {
	ExitStatus *string `json:"exit_status,omitempty"`
	Limit      *int    `json:"limit,omitempty"`
}

type ManualRetry struct {
	Allowed        *bool   `json:"allowed,omitempty"`
	PermitOnPassed *bool   `json:"permit_on_passed,omitempty"`
	Reason         *string `json:"reason,omitempty"`
}

// Groupable is defined to allow Command step to be used as a [step.GroupStep]
func (Command) Groupable() {
}
