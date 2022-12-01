package pipeline_test

import (
	"encoding/json"
	"fmt"

	"github.com/hasura/go-buildkite-dsl/agent"
	"github.com/hasura/go-buildkite-dsl/pipeline"
	"github.com/hasura/go-buildkite-dsl/step"
)

func str(s string) *string {
	return &s
}

func Example_simple() {
	p := pipeline.New("sample-pipeline")
	p.Agents = agent.Agent{
		"queue": "aws-spot-ubuntu-small",
	}
	jobA := step.Command{
		Label:   str("job a"),
		Key:     str("job-a"),
		Command: str(`echo "hello from job-a"`),
	}
	p.Steps = append(p.Steps, jobA)

	b, err := json.MarshalIndent(p, "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// {
	// 	"agents": {
	// 		"queue": "aws-spot-ubuntu-small"
	// 	},
	// 	"steps": [
	// 		{
	// 			"command": "echo \"hello from job-a\"",
	// 			"key": "job-a",
	// 			"label": "job a"
	// 		}
	// 	]
	// }
}
