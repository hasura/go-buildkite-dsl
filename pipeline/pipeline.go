package pipeline

import "github.com/hasura/go-buildkite-dsl/agent"

type Step interface {
	Step()
}

type Slug string

type Pipeline struct {
	Slug   Slug              `json:"-"`
	Agents agent.Agent       `json:"agents,omitempty"`
	Env    map[string]string `json:"env,omitempty"`
	Steps  []Step            `json:"steps"`
}

func New(s Slug) *Pipeline {
	return &Pipeline{
		Slug: s,
	}
}

func (p *Pipeline) Register(s Step) {
	p.Steps = append(p.Steps, s)
}
