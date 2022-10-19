// Package step consists of the buildkite DSL for defining
// various types of steps.
//
// For more details, refer this [docs page].
//
// [docs page]: https://buildkite.com/docs/pipelines/defining-steps
package step

import "encoding/json"

type Dependency struct {
	Step         string `json:"step"`
	AllowFailure *bool  `json:"allow_failure,omitempty"`
}

type Skip struct {
	Reason *string
}

func WithReason(reason string) *Skip {
	return &Skip{
		Reason: &reason,
	}
}

func WithoutReason() *Skip {
	return &Skip{}
}

func (s Skip) MarshalJSON() ([]byte, error) {
	if s.Reason != nil {
		return json.Marshal(*s.Reason)
	}

	return []byte("true"), nil
}
