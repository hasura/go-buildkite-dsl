package step

import "encoding/json"

type Plugin interface {
	Name() string
	Version() string

	json.Marshaler
}
