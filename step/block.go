package step

import (
	"encoding/json"
	"errors"
)

// Block is used to created a Block step in buildkite.
//
// For more details, refer: https://buildkite.com/docs/pipelines/block-step
type Block struct {
	// Block step name.
	//
	// Example: "Release"
	Block string `json:"block"`

	// The state that the build is set to when the build is blocked by this block step.
	// The default is passed. When the blocked_state of a block step is set to failed,
	// the step that triggered it will be stuck in the running state until it is manually unblocked.
	//
	// Values: passed, failed, running
	BlockedState *BlockedState `json:"blocked_state,omitempty"`

	// A unique string to identify the block step.
	Key *string `json:"key,omitempty"`

	BlockAttributes
}

type BlockAttributes struct {
	// The instructional message displayed in the dialog box when the unblock step is activated.
	//
	// Example: "Release to production?"
	// Example: "Fill out the details for this release"
	Prompt *string `json:"prompt,omitempty"`

	// A list of input fields required to be filled out before unblocking the step.
	//
	// Available input field types: text, select
	Fields []Field `json:"fields,omitempty"`

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
}

type BlockedState string

const (
	Passed  BlockedState = "passed"
	Failed  BlockedState = "failed"
	Running BlockedState = "running"
)

type Field interface {
	MarshalJSON() ([]byte, error)
}

type TextField struct {
	// The meta-data key that stores the field's input (using the buildkite-agent meta-data command)
	// The key may only contain alphanumeric characters, slashes or dashes.
	//
	// Example: "release-name"
	Key string `json:"key"`

	// Text field name
	//
	// Example: "Release Name"
	Text *string `json:"text,omitempty"`

	// The explanatory text that is shown after the label.
	//
	// Example: "What's the code name for this release? :name_badge:"
	Hint *string `json:"hint,omitempty"`

	// A boolean value that defines whether the field is required for form submission.
	Required *bool `json:"required,omitempty"`

	// The value that is pre-filled in the text field.
	//
	// Example: "Flying Dolphin"
	Default *string `json:"default,omitempty"`
}

func (t TextField) MarshalJSON() ([]byte, error) {
	if len(t.Key) == 0 {
		return nil, errors.New("expected a non-empty Key value for TextField")
	}

	type psuedo TextField
	p := psuedo(t)
	return json.Marshal(p)
}

type SelectField struct {
	// The meta-data key that stores the field's input (using the buildkite-agent meta-data command)
	// The key may only contain alphanumeric characters, slashes or dashes.
	//
	// Example: "release-stream"
	Key string `json:"key"`

	// The list of select field options.
	// For 6 or less options they'll be displayed as radio buttons, otherwise they'll be displayed in a dropdown box.
	// If selecting multiple options is permitted the options will be displayed as checkboxes.
	Options []SelectFieldOption `json:"options"`

	// A boolean value that defines whether multiple options may be selected.
	// When multiple options are selected, they are delimited in the meta-data field by a line break (\n)
	//
	// Default: false
	Multiple *bool `json:"bool,omitempty"`

	// The text displayed directly under the select field's label.
	//
	// Example: "Which release stream does this belong in? :fork:"
	Hint *string `json:"hint,omitempty"`

	// A boolean value that defines whether the field is required for form submission.
	//
	// Default: true
	Required *bool `json:"required,omitempty"`

	// The value of the option or options that will be pre-selected.
	// When multiple is enabled, this can be an array of values to select by default.
	//
	// Example: "beta"
	Default *string `json:"default,omitempty"`
}

type SelectFieldOption struct {
	// The text displayed for the option.
	//
	// Example: "Stable"
	Label string `json:"label"`

	// The value to be stored as meta-data (to be later retrieved using the buildkite-agent meta-data command)
	//
	// Example: "stable"
	Value string `json:"value"`
}

func (s SelectField) MarshalJSON() ([]byte, error) {
	if len(s.Key) == 0 {
		return nil, errors.New("expected a non-empty Key value for SelectField")
	}

	for _, option := range s.Options {
		if len(option.Label) == 0 {
			return nil, errors.New("expected a non-empty Label for select field options")
		}

		if len(option.Value) == 0 {
			return nil, errors.New("expected a non-empty Value for select field options")
		}
	}

	type psuedo SelectField
	p := psuedo(s)
	return json.Marshal(p)
}

// Step is defined on Block step to allow it to be used as [pipeline.Step]
func (Block) Step() {}
